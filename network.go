import("net"; "fmt"; "bufio";"time")
//Master
/*
// Send a message to the server:  "Connect to: " ~ your IP ~ ":" ~ port ~ "\0"

// do not need IP, because we will set it to listening state
addr = new InternetAddress(localPort)
acceptSock = new Socket(tcp)

// You may not be able to use the same port twice when you restart the program, unless you set this option
acceptSock.setOption(reuseAddr, true)
acceptSock.bind(addr)
// backlog = Max number of pending connections waiting to connect()
newSock = acceptSock.listen(backlog)

// Spawn new thread to handle recv()/send() on newSock
*/

func server(ip string){ 

     connect, err := net.Dial("tcp", ip) 
     if err != nil {
     	    fmt.Printf("error server func: %s\n", err.Error())
     }

	status, err := bufio.NewReader(connect).ReadString('\n')
	fmt.Print(status)
}

//Client
/*
addr = new InternetAddress(serverIP, serverPort) 
sock = new Socket(tcp) // TCP, aka SOCK_STREAM
sock.connect(addr)
// use sock.recv() and sock.send()
*/
func client(ip string) {


}
/*************************************/


