package main
//https://groups.google.com/forum/#!topic/golang-china/KG0Bgf0CAQc
import (
	"fmt"
	."net"
	"time"
	"strconv"
)
const (
	MAXWAIT = time.Second
)


func main() {
	nr := backup()
	// spawn new window
	//	mate terminal
	primary(nr)
}

func backup() int {
	//	nr
	buf := make([]byte, 1024)
	address, err := ResolveUDPAddr("udp", "129.241.187.255:20022")
	conn, _ := ListenUDP("udp", address)

	for {
		conn.SetReadDeadline(time.Now().Add(1*time.Second))
		n, _, err := conn.ReadFromUDP(buf) //n contanis numbers of used bytes
		if err == nil {		
			nr = strconv.Atoi(buf[0:n])
		} else {
			break;
		}
	}
	conn.Close()
	return nr;
}

func primary(nr int){
		conn := DialUDP("udp", "localhost", 20022)
		for {
			nr++
			fmt.Println(nr)
			conn.Write([]byte(strconf.Itoa(nr)))
			time.Sleep(300*time.Millisecond)
		}
}
