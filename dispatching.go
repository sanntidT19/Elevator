package modules

const (
FLOORS = 4
ELEV = 3
)

var hallway [FLOORS*2] bool //array of 2xfloors - first half for order up, second for order down.
var highest int
var lowest int
var yettoserve [FLOORS*2] bool

type Cabin struct {
	name int
	direction bool // true if elevator is going up
	movement bool // true if elevator is moving
	position int // current position
	servicelist [FLOORS*2] bool // list of current orders for elevator
	load int // current load of elevator
	buttons [FLOORS*2]bool //array of 2xfloors - first half for orders up, second for orders down. input from inside of cabin
	
	
}
	

func main() {
	
}
// defining load - load is number of floors that needs to be served (sum of both ways) - we should probably use service time instead
	for j := 0; i < FLOORS; i++ {
        	if c.servicelist[j] {
			c.load++
		}
    	}
}
	

// floors where elevators are going are already served
func (c *Cabin) already {
	for i := 0; i < ELEV; i++ {
		for j := 0; i < FLOORS; i++ {
        		yettoserve[i] = yettoserve[i] - (c.buttons[i] * hallway[i])
    		}
}

func (c *Cabin) level2 {	



}
