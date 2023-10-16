package main

import (
	"cmp"
	"fmt"
	"slices"
	"strconv"
	"sync"
)

func runAndClose(c cmd, in, out chan interface{}, waiter *sync.WaitGroup) {
	defer func() {
		close(out)
		waiter.Done()
	}()

	c(in, out)
}

func RunPipeline(cmds ...cmd) {
	if len(cmds) == 0 {
		return
	}
	in, out := make(chan interface{}), make(chan interface{})

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
	var usersID = make(map[uint64]string)
	var mu sync.Mutex
	wg := &sync.WaitGroup{}

	for input := range in {
		input := input
		wg.Add(1)

		go func() {
			user := GetUser(input.(string))

			mu.Lock()
			defer func() {
				mu.Unlock()
				wg.Done()
			}()

			_, exist := usersID[user.ID]
			if !exist {
				usersID[user.ID] = user.Email
				out <- user
			}
		}()
	}

	wg.Wait()
}

func callGetMesasges(users []User, out chan interface{}, wg *sync.WaitGroup) {
	defer wg.Done()

	messages, err := GetMessages(users...)
	if err != nil {
		fmt.Println("Error: ", err)
		return
	}

	for _, msg := range messages {
		out <- msg
	}
}

// in - User
// out - MsgID
func SelectMessages(in, out chan interface{}) {
	wg := &sync.WaitGroup{}
	var users []User
	const UsersBatchLen = 2
	const UsersWithoutBatchLen = 1

	for input := range in {
		users = append(users, input.(User))
		if len(users) < UsersBatchLen {
			continue
		}

		wg.Add(1)
		go callGetMesasges(users, out, wg)
		users = []User{}
	}

	if len(users) == UsersWithoutBatchLen {
		wg.Add(1)
		go callGetMesasges(users, out, wg)
	}

	wg.Wait()
}

// in - MsgID
// out - MsgData
func CheckSpam(in, out chan interface{}) {
	wg := &sync.WaitGroup{}
	const workersCount = 5

	for i := 0; i < workersCount; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()

			for input := range in {
				hasSpam, err := HasSpam(input.(MsgID))
				if err != nil {
					fmt.Println("Error: ", err)
					continue
				}

				out <- MsgData{input.(MsgID), hasSpam}
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
		out <- strconv.FormatBool(data.HasSpam) + " " + strconv.FormatUint(uint64(data.ID), 10)
	}
}
