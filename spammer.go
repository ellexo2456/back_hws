package main

import (
	"log"
	"slices"
)

func RunPipeline(cmds ...cmd) {
	if len(cmds) == 0 {
		return
	}

	var in, out chan interface{}
	cmds[0](in, out)

	for _, cmd := range cmds[1:] {
		in = make(chan interface{})
		cmd(out, in)
		out = in
	}
}

// in - string
// out - User
func SelectUsers(in, out chan interface{}) {
	for input := range in {
		user := GetUser(input.(string))
		out <- user
	}

	close(out)
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

	close(out)
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

	close(out)
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
	close(out)
}
