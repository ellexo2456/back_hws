package main

import (
	"fmt"
	"slices"
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

// in - User
// out - MsgID
func SelectMessages(in, out chan interface{}) {
	wg := &sync.WaitGroup{}

	for input := range in {
		input := input
		wg.Add(1)
		go func() {
			if messages, err := GetMessages(input.(User)); err != nil {
				fmt.Println("Error: ", err)
				return
			} else {
				for _, msg := range messages {
					out <- msg
				}
			}
			wg.Done()
		}()
	}

	wg.Wait()
}

// in - MsgID
// out - MsgData
func CheckSpam(in, out chan interface{}) {
	for input := range in {
		if hasSpam, err := HasSpam(input.(MsgID)); err != nil {
			fmt.Println("Error: ", err)
			return
		} else {
			out <- MsgData{input.(MsgID), hasSpam}
		}
	}
}

// in - MsgData
// out - string
func CombineResults(in, out chan interface{}) {
	var msgsData []MsgData
	for input := range in {
		msgsData = append(msgsData, input.(MsgData))
	}

	slices.SortFunc(msgsData, func(a, b MsgData) int {
		if a.HasSpam {
			return 1
		}
		if b.HasSpam {
			return -1
		}

		return 0
	})
}
