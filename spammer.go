package main

import (
	"log"
	"slices"
)

func runAndClose(c cmd, in, out chan interface{}) {
	c(in, out)
	defer close(out)
}

func RunPipeline(cmds ...cmd) {
	if len(cmds) == 0 {
		return
	}
	in := make(chan interface{})
	out := make(chan interface{})
	//go func(in, out chan interface{}, c cmd) {
	//	c(in, out)
	//	defer close(out)
	//}(in, out, cmds[0])
	go runAndClose(cmds[0], in, out)

	for _, c := range cmds[1:] {
		in = make(chan interface{})
		//fmt.Println(i)
		//go func(in, out chan interface{}, c cmd) {
		//	c(in, out)
		//	defer close(out)
		//}(out, in, c)
		go runAndClose(c, out, in)
		out = in
	}
}

// in - string
// out - Use
func SelectUsers(in, out chan interface{}) {
	var usersID []uint64
	for input := range in {
		user := GetUser(input.(string))
		if !slices.Contains(usersID, user.ID) {
			usersID = append(usersID, user.ID)
			out <- user
		}
	}

	//close(out)
}

// in - User
// out - MsgID
func SelectMessages(in, out chan interface{}) {
	for input := range in {
		if messages, err := GetMessages(input.(User)); err != nil {
			log.Fatal(err)
		} else {
			for _, msg := range messages {
				out <- msg
			}
		}
	}
}

// in - MsgID
// out - MsgData
func CheckSpam(in, out chan interface{}) {
	for input := range in {
		if hasSpam, err := HasSpam(input.(MsgID)); err != nil {
			log.Fatal(err)
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
