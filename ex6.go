package main
//https://groups.google.com/forum/#!topic/golang-china/KG0Bgf0CAQc
import (
	"fmt"
	."net"
	"time"
	"strconv"
	"os/exec"
)
const (
	MAXWAIT = time.Second
)


func main() {
	nr := backup()
	cmd, _ := exec.Command("mate-terminal").Output()
	// spawn new window
	//	mate terminal
	primary(nr)
	cmd.exec.Command("go run ex6.go").Output()
}

func backup() int {
	var nr int
	buf := make([]byte, 1024)
	address, err := ResolveUDPAddr("udp", "129.241.187.255:20020")
	conn, _ := ListenUDP("udp", address)

	for {
		conn.SetReadDeadline(time.Now().Add(1*time.Second))
		n, _, _:= conn.ReadFromUDP(buf) //n contanis numbers of used bytes
		if err == nil {		
			nr,_ = strconv.Atoi(string(buf[0:n]))
		} else {
			break;
		}
	}
	conn.Close()
	return nr;
}

func primary(nr int){
		address, _ := ResolveUDPAddr("udp", "129.241.187.255:20020")
		conn, _ := DialUDP("udp", address, nil)
		for {
			nr++
			fmt.Println(nr)
			conn.Write([]byte(strconv.Itoa(nr)))
			time.Sleep(300*time.Millisecond)
		}
}
