package main

import (
	"fmt"
	"time"
)

type msg_log struct {
	msg      string
	index    int
	date_str string
}

func print_msg(ch chan msg_log) {
	for {
		data := <-ch
		fmt.Printf("[*]%s:%d: %s \n", data.date_str, data.index, data.msg)
	}
}

func main() {
	msg_chan := make(chan msg_log)

	//fmt.Println(GoMuLog.TestAdd(1, 2))

	go print_msg(msg_chan)

	time.Sleep(1 * time.Second)

	msg_chan <- msg_log{
		msg:      "info 1",
		index:    0,
		date_str: "14:26",
	}
	time.Sleep(1 * time.Second)

	msg_chan <- msg_log{
		msg:      "info 2",
		index:    0,
		date_str: "14:27",
	}
	time.Sleep(1 * time.Second)

	msg_chan <- msg_log{
		msg:      "info 3",
		index:    0,
		date_str: "14:28",
	}
	time.Sleep(1 * time.Second)

	msg_chan <- msg_log{
		msg:      "info 4",
		index:    0,
		date_str: "14:29",
	}
	time.Sleep(5 * time.Second)
}
