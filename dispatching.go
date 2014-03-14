package modules

const (
FLOORS = 4
ELEV = 3
TIMEOFSTOP = 3
TIMETRAVEL = 1
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
	timeload int // current load of elevator
	buttons [FLOORS*2]bool //array of 2xfloors - first half for orders up, second for orders down. input from inside of cabin
	LowestFloorUp int
	LowestFloorDown int
	HighestFloorUp int
	HighestFloorDown int
	
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
		for j := 0; j < FLOORS; i++ {
        		yettoserve[i] = yettoserve[i] - (c.buttons[i] * hallway[i])
    		}
}

// this should probably be evaluated by slave sent along with the other stuff
func (c *Cabin) find_HighestFloorUp {	
	for j := 0; j < FLOORS; j++ {
        	if buttons[j] {
			c.HighestFloorUp = j		
		}
    	}	

func (c *Cabin) find_LowestFloorDown {	
	for j := FLOORS*2; j > FLOORS; j-- {
        	if buttons[j] {
			c.LowestFloorDown = j		
		}
    	}	

func (c *Cabin) calculate_timeload {	
	var stops int
	for j := 0; j < FLOORS*2; j++ {
        	if buttons[j] {
			stops++
		}
	}	
	if c.direction {
			c.timeload = (2 * c.HighestFloorUp - c.position - c.LowestFloorDown) * TIMETRAVEL + stops * TIMEOFSTOP)
	}	
	else {
			c.timeload = (c.position + c.HighestFloorUp - 2 * c.LowestFloorDown) * TIMETRAVEL + stops * TIMEOFSTOP)
	}

}









