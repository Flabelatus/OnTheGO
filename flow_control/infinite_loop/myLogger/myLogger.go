package mylogger

import "log"

func ListenForLog(ch chan string) {
	// infinite loop
	for {
		msg := <-ch
		log.Println(msg)
	}
}
