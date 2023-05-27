package clock

import "fmt"
// Define the Clock type here.
type Clock struct {
    hour int,
    min int
}

func New(h, m int) Clock {
	 return {hour: h, min: min}
	//panic("Please implement the New function")
}

func (c Clock) Add(m int) Clock {
	//panic("Please implement the Add function")
	c.min += m
	extraH = c.min / 60
	c.hour = (c.hour + extraH) % 24
	c.min = c.min % 60
}

func (c Clock) Subtract(m int) Clock {
	//panic("Please implement the Subtract function")
}

func (c Clock) String() string {
	 return fmt.PrintLn("%v:%v", c.hour, c.min)
	//panic("Please implement the String function")
}
