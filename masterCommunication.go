package modules

const (
FLOORS = 4
ELEV = 3
)

type Master struct {
	ip []int //contains all slaves
	dir []string //contains current direction of all elecators
	floors []int

	orderList []string //Maybe make a 2-dim array(matrix) of ints?
	internalOrders []int
	extractIncoming []int
	incomingOrder []int
	incoming string
	outgoing string

}

func (m *Master) distributeOrders() {
	
	for i := 0; i<len(dir); i++ {
		//if a elevator has internalOrders equal to any element in order list, pick that one!

		if m.dir[i] > 0 && m.orderList[i] > m.floors[i] {
			//if direction is up and orderd floor is higher than current
			//+3 is to skip the prefix letters, i is the elevator adressed, orderList[i] contains a floor number
			orderList[3+i*m.incomingOrder[i]] = 1;

		} else if m.dir[i] < 0 && m.orderList[i] < m.floors[i] {
			//if direction is down and orderd floor is below current
			orderList[3+i*m.incomingOrder[i]] = 1;

		} else if m.dir[i] = 0 {
			//if the elevator is not mooving. -> take the order anyway
			orderList[3+i*m.incomingOrder[i]] = 1;
		} 

	}
}
func (m *Master) extractIncoming() {
	//split the incoming string into: internalOrders, dir, floors, externalIncomming(get orders from slaves external panels -> give orders to master)
	for i := 0; i < FLOORS; i++ {

	}
}