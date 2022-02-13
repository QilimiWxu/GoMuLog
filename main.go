package GoMuLog

import (
	"fmt"
	"time"
)

type msg_log struct{
	msg string
	index int
	date_str string
}

g_msg_chan := make(chan msg_log)

func print(){
	for{
		data = <- g_msg_chan
		if data != nil{
			fmt.println("[*]%s:%d: %s \n", data.)
		}
		else{
			break;
		}
	}
}

func main() {
	fmt.Println("welcome GoMuLog")
	//fmt.Println(GoMuLog.TestAdd(1, 2))

	go print()

	time.Sleep(1 * time.Second())
}
