package clock

import "fmt"

// Define the Clock type here.
type Clock struct {
	hour int
	min  int
}

func New(h, m int) Clock {
	// if h < 0 {
	// 	fmt.Println("h < 0, ", h, " h%24", h%24)
	// }
	extraH := m / 60
	newH := ((h+extraH)%24 + 24) % 24
	newM := (m%60 + 60) % 60
	fmt.Println("NEW: Incoming h, m: ", h, m, ", extraH: ", extraH, ", newH:", newH, ", newM: ", newM)
	return Clock{hour: newH, min: newM}
	//panic("Please implement the New function")
}

func (c Clock) Add(m int) Clock {
	//panic("Please implement the Add function")
	c.min += m
	extraH := c.min / 60
	c.hour = (c.hour + extraH) % 24
	c.min = c.min % 60
	return c
}

func (c Clock) Subtract(m int) Clock {
	//panic("Please implement the Subtract function")
	c.min -= m
	if c.min < 0 {
		extraH := c.min / 60
		c.hour -= extraH - 1
		c.min %= 60
		c.hour %= 24
	}
	return c
}

func (c Clock) String() string {
	return fmt.Sprintf("%02d:%02d", c.hour%24, c.min%60)
	//panic("Please implement the String function")
}
