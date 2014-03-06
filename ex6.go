package main

import (
	"fmt"
	"net"
	"time"
)
const (
	MAXWAIT = time.Secound
)

type contatiner struct {
	nr int
	time time.Time
}
var ch chan

func main() {
	ch = make(chan struct)
	go countingNumbers(0, true)
}

func countingNumbers(nr int, isPrimary bool) {
	if isPrimary{ 
		nr++
		cont := contatiner{}
		cont.nr := nr
		cont.time := time.Now()
		ch <- cont
		
		for {
			//this should be a critical section
			//send timestamp(?) and i
			fmt.Println(i)
			ch <- i
			i++
		}
	}
	else { //backup section
		select {
		case cont <- ch:
			if cont.time-time.Now() > MAXWAIT {
				//make this primary
				go countingNumbers(cont.nr, true)
			}
		default :
			fmt.Println("error")
	}
	// recive "i am alive"
	//if long since "iah" become primary, make new backup	
	}
}
