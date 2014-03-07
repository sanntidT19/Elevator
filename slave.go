package slave

import (
	"fmt"
	"master"
	"string"
	"strconv"
	"time"
	"workingChannels"
)
const (
	UP=2
	DOWN=1
	STATIONARY=0
)

type Slave struct {
	nr int //witch number in orderString
	dir int 
	currentFloor int
	internalList []int
	LastTimeStamp Time.time

	incommning string
	outgoing string
}

func (s *Slave) extractIncomming( <- workingChannels.orderChannel) {
	if strings.HasPrefix(s.incomming, "odr") {
		s.incomming = strings.TrimPrefix(s.incomming, "odr")

	} else if strings.HasPrefix(s.incomming, "acn") {
		//if master is still there, continue.
		//else chooseMaster
	}
}

func (s *Slave) packOutgoing() {
	var orderString string

	for i := 0; i < len(s.internalOrderList); i++ {
		orderString += strconv.Itoa(s.internalOrderList[i])
	}
	s.outgoing = strconv.Itoa(s.nr) + orderString + strconv.Itoa(s.dir) + strconv.Itoa(s.currentFloor)
	workingChannels.orderChannel <- s.outgoing
}

func executeOrder() {
	//run elevator to next floor in queue


}
//needs to be runned in a for loop
func (s *Slave) checkMaster() {
	if Time.Before(Time.Millisecond*10) {
		temp := 	
	}
	
}