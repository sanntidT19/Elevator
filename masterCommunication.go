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
	incoming string
	outgoing string

}

func (m *Master) distributeOrders() {
	
	for i := 0; i<len(dir); i++ {
		if m.dir[i] > 0 && m.orderList[i] > m.floors[i] {
			//if direction is up and orderd floor is higher than current

		} else if m.dir[i] < 0 && m.orderList[i] < m.floors[i] {
			//if direction is down and orderd floor is below current
		} else if m.dir[i] = 0 {
			//if the elevator is not mooving. -> take the order anyway
		} 

	}
}
func (m *Master) extractIncoming() {

}