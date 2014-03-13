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
	/*
		//_ = exec.Command("go run testofgithub.go")
		cmd := exec.Command("go run testofgithub.go")
		err := cmd.Start()
		if err != nil {
			fmt.Println("err")
		}
		err = cmd.Wait()
	*/
	cmd := exec.Command("mate-terminal", "-x", "go", "run", "testofgithub.go")
	cmd.Start()

	primary(nr)

	// spawn new window
	//	mate terminal

	//cmd.Output("go run testofgithub.go")

}

func backup() int {
	fmt.Println("i have backup")
	var nr int
	buf := make([]byte, 1024)
	address, err := ResolveUDPAddr("udp", ":20019")
	conn, err := ListenUDP("udp", address)
	fmt.Println("backup conn create error:", err)

	for {
		fmt.Println("backup")
		conn.SetReadDeadline(time.Now().Add(1 * time.Second))
		n, _, err := conn.ReadFromUDP(buf) //n contanis numbers of used bytes
		if err == nil {
			fmt.Println("h")
			fmt.Println("buf: ", string(buf[0:n]))
			nr, _ = strconv.Atoi(string(buf[0:n]))
			fmt.Println("nr: ", nr)
		} else {
			fmt.Println("t")
			break
		}
	}
	conn.Close()
	return nr
}

func primary(nr int) {
	address, _ := ResolveUDPAddr("udp", "129.241.187.255:20019")
	conn, _ := DialUDP("udp", nil, address)
	var a []byte
	fmt.Println(a)
	for {
		fmt.Println("primary")
		nr++
		fmt.Println(nr)

		a = []byte(strconv.Itoa(nr))
		_, err := conn.Write(a)
		if err != nil {
			fmt.Println("fail")
		}
		time.Sleep(300 * time.Millisecond)
	}
}
