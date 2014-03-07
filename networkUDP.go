package modules
import (
	"net"
)

//SEND[psudocode]

//broadcastIP = #.#.#.255. First three bytes are from the local IP
//addr = new InternetAddress(broadcastIP, port)
//sendSock = new Socket(udp) // UDP, aka SOCK_DGRAM
//sendSock.setOption(broadcast, true)
//sendSock.sendTo(message, addr)

func senderConnction() {
	brodcastIP 	:= 
	port 		:=

	var buf [1024]byte 
	addr, err = net.resolveUDPaddr("udp", port) //connects
	sock, err = net.ListenUDP("udp", addr)

	for {
		rlen, remote, err := sock.ReadFormUDP(buf)
	}

	if err != nil {
		//handle error
	}
	
}

//RECIVE[psudocode]
func UDPClient() {
	masterAddr := getMasterAddr()
	buffer := make(byte[1024])
	serverAddr, err := net.resolveUDPaddr("udp", masterAddr)
}
 // makes a bytearray

//InternetAddress fromWho
//recvSock = new Socket(udp)
//recvSock.bind(addr)         // same addr as sender
//loop {
//    buffer.clear
    // fromWho will be modified by ref here. Or it's a return value. Depends.
 //   recvSock.receiveFrom(buffer, ref fromWho)
//    if(fromWho.IP != localIP){      // check we are not receiving from ourselves
        // do stuff with buffer
  //  }
//}

