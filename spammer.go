package main

import (
	"cmp"
	"fmt"
	"slices"
	"strconv"
	"sync"
)

func runAndClose(c cmd, in, out chan interface{}, waiter *sync.WaitGroup) {
	c(in, out)

	defer func() {
		close(out)
		waiter.Done()
	}()
}

func RunPipeline(cmds ...cmd) {
	if len(cmds) == 0 {
		return
	}
	in := make(chan interface{})
	out := make(chan interface{})

	wg := &sync.WaitGroup{}
	go runAndClose(cmds[0], in, out, wg)
	wg.Add(1)

	for _, c := range cmds[1:] {
		in = make(chan interface{})

		go runAndClose(c, out, in, wg)
		wg.Add(1)
		out = in
	}

	wg.Wait()
}

// in - string
// out - Use
func SelectUsers(in, out chan interface{}) {
	var usersID []uint64
	var mu sync.Mutex
	wg := &sync.WaitGroup{}

	for input := range in {
		input := input
		wg.Add(1)

		go func() {
			user := GetUser(input.(string))
			mu.Lock()
			if !slices.Contains(usersID, user.ID) {
				usersID = append(usersID, user.ID)
				out <- user
			}
			mu.Unlock()
			wg.Done()
		}()
	}

	wg.Wait()
}

func callGetMesasges(users []User, out chan interface{}, wg *sync.WaitGroup) {
	//mu.Lock()
	if messages, err := GetMessages(users...); err != nil {
		fmt.Println("Error: ", err)
		return
	} else {
		for _, msg := range messages {
			out <- msg
		}
	}
	//mu.Unlock()
	wg.Done()
}

// in - User
// out - MsgID
func SelectMessages(in, out chan interface{}) {
	wg := &sync.WaitGroup{}
	var users []User
	//var mu &sync.Mutex

	for input := range in {
		users = append(users, input.(User))
		if len(users) < 2 {
			continue
		}

		wg.Add(1)
		go callGetMesasges(users, out, wg)
		users = []User{}
	}

	if len(users) == 1 {
		wg.Add(1)
		go callGetMesasges(users, out, wg)
	}

	wg.Wait()
}

// in - MsgID
// out - MsgData
func CheckSpam(in, out chan interface{}) {
	wg := &sync.WaitGroup{}

	for i := 0; i < 5; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()

			for input := range in {
				if hasSpam, err := HasSpam(input.(MsgID)); err != nil {
					fmt.Println("Error: ", err)
					return
				} else {
					out <- MsgData{input.(MsgID), hasSpam}
				}
			}
		}()
	}

	wg.Wait()
}

// in - MsgData
// out - string
func CombineResults(in, out chan interface{}) {
	var msgsData []MsgData
	for input := range in {
		msgsData = append(msgsData, input.(MsgData))
	}

	slices.SortFunc(msgsData, func(a, b MsgData) int {
		if a.HasSpam && !b.HasSpam {
			return -1
		} else if !a.HasSpam && b.HasSpam {
			return 1
		}

		return cmp.Compare(a.ID, b.ID)
	})

	for _, data := range msgsData {
		s := strconv.FormatBool(data.HasSpam) + " " + strconv.FormatUint(uint64(data.ID), 10)
		out <- s
	}
}
