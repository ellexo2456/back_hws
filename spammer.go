package main

import (
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
	var workerCount int64
	var mu sync.Mutex
	wg := &sync.WaitGroup{}

	//mu.Lock()
	for input := range in {
		var flag bool
		mu.Lock()
		if workerCount < 5 {
			flag = true
		} else {
			flag = false
		}
		mu.Unlock()

		if flag {
			mu.Lock()
			workerCount++
			mu.Unlock()

			//atomic.AddInt64(&workerCount, 1)
			wg.Add(1)
			go func(input interface{}) {
				defer func() {
					mu.Lock()
					workerCount--
					mu.Unlock()
					//atomic.AddInt64(&workerCount, -1)

					wg.Done()
				}()

				if hasSpam, err := HasSpam(input.(MsgID)); err != nil {
					fmt.Println("Error: ", err)
					return
				} else {
					out <- MsgData{input.(MsgID), hasSpam}
				}

			}(input)

		}
	}

	wg.Wait()
	//mu.Unlock()
	//
	//mu.Lock()
	//for workerCount != 0 {
	//}
	//mu.Unlock()
}

// in - MsgData
// out - string
func CombineResults(in, out chan interface{}) {
	var msgsData []MsgData
	for input := range in {
		msgsData = append(msgsData, input.(MsgData))
	}

	slices.SortStableFunc(msgsData, func(a, b MsgData) int {
		if a.HasSpam {
			return -1
		}
		if b.HasSpam {
			return 1
		}

		return 0
	})

	for _, data := range msgsData {
		s := strconv.FormatBool(data.HasSpam) + " " + strconv.FormatUint(uint64(data.ID), 10)
		out <- s
	}
}
