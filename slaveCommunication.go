//Slave communication module
package modules //(?)
import (
	"net"
	"strings"
	"strconv"//for Itoa; int to string.
	"encoding/json"
	)

type Slave struct {
	nr int //or name is ip?
	conn net.Conn
	internalOrderList []int
	//externalOrders -> send to master for distributing
	quit chan bool
	dir string
	currentFloor int
	incoming string //holds all relevatn elevator information and lists
	outgoing string //internal list, direction, floor
}
func (s *Slave) incommingHandler() {
	if strings.HasPrefix(incoming, "odr") {
		//update OrderList
	} else if strings.HasPrefix(incoming, "acn") {
		//master is still there.
	}
}

//Sending messages to master, including internal list, direction and currentFLoor
func (s *Slave) outgoingHandler() { 
	var orderString string

	for i := 0; i < len(s.internalOrderList); i++ {
		orderString += strconv.Itoa(s.internalOrderList[i])
	}
	s.outgoing += orderString + s.dir + strconv.Itoa(s.currentFloor)
}


func (s *Slave) Read(buffer []byte) bool {
	structRead, err := s.conn.Read(buffer)
	if err != nil {
		
		Log("error in Read " + err)
		return false
	}
	s.incoming := string(structRead)
	return true
}


func (s *Slave) Close() {
	s.quit <- true
	s.conn.Close()
	s.ReMoveMe()
}


func Log(v ...interface{}) {
	fmt.Println(v...)
}

//NOT unit tested
func IOHandler(incomingOrder <-chan string, slaveList *list.List) {
    for {
        //waiting for input
        input := <- incomingOrder
        for e := slaveList.Front(); e != nil; e = e.Next() {
            slave := e.Value.(slave)
            slave.incomingOrder <-input
        }
    }
}
//NOT Unit tested
func SlaveInputReader(slave *Slave) {
	buffer := make([]byte 1024)

	for slave.Read(buffer) {
		order := string(buffer)
		slave.totalOrderList = order 
		send := slave.nr+"> "+string(buffer)
		slave.outgoing <- send
		for i := 0; i < 1024; i++ {
			buffer[i] = 0x00

		}
	}

	slave.outgoing <- slave.nr + "still here"
}
//NOT unit tested
func (s *Slave) slaveSender() {
	bytes, err :=  json.Marshal(s.outgoing)
	
	if err =! nil {
		Log("Error in slave sender" + err)
	}
}