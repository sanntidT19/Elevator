package main

//https://groups.google.com/forum/#!topic/golang-china/KG0Bgf0CAQc
import (
	"fmt"
	. "net"
	"os/exec"
	"strconv"
	"time"
)

const (
	MAXWAIT = time.Second
)

func main() {
	nr := backup()

	cmd := exec.Command("mate-terminal", "-x", "go", "run", "ex6.go")
	cmd.Start()

	primary(nr)

}

func backup() int {
	//fmt.Println("i have backup")
	var nr int
	buf := make([]byte, 1024)
	address, err := ResolveUDPAddr("udp", ":20019") //leser bare fra porten generellt
	conn, err := ListenUDP("udp", address)
	
	if err != nil {
		fmt.Println("backup conn create error:", err)
	}
	
	fmt.Println("backup")
	for {
		
		conn.SetReadDeadline(time.Now().Add(5 * time.Second))
		n, _, err := conn.ReadFromUDP(buf) //n contanis numbers of used bytes
		if err == nil {
			nr, _ = strconv.Atoi(string(buf[0:n]))
			//fmt.Println("nr: ", nr)
		} else {
			break
		}
	}
	conn.Close()
	return nr
}

func primary(nr int) {
	address, _ := ResolveUDPAddr("udp", "129.241.187.255:20019")
	conn, _ := DialUDP("udp", nil, address)
	
	//fmt.Println(a)
	fmt.Println("primary")
	for {
		
		nr++
		fmt.Println(nr)

		_, err := conn.Write([]byte(strconv.Itoa(nr)))
		if err != nil {
			fmt.Println("fail")
		}
		time.Sleep(2000 * time.Millisecond)
	}
}
