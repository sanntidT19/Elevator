import(	"net"; 
		"fmt")
//The main idea: used as an interface between computers to communicate

func configureNetwork() {
	//start by udp, recive callback from elevators(slaves)
	resolveUdpAdrr("udp","localhost")

	master := chooseMasterSlaveConfig()
	var ipList []int

}
func chooseMaster(ipList [] int) { 
	//choose new master
	//Pick at random(?)
	connectTcp()
}

func connectTcp(ipList []int)  {
	//assuming master is on top of list

	//connects to all slaves
	for i = 0; i < len(ipList); i++ {
		conn, err := net.Dial("tcp", ipList[i]) 
	
	    if err != nil {
    	    fmt.Printf("error Tcp connection: %s\n", err.Error())
    	}
	}
}

func send(sending string, conn net) {
	conn.bufio.
}
//does this need a connection to read from?
func recive(elevStatus string, conn) {

	status, err := bufio.NewReader(conn).ReadString('\n')
	fmt.Print(status)
}
func detectChangedNetwork() {

}

func closeConnection() {

}