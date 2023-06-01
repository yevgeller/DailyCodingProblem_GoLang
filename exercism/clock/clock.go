package clock

import "fmt"

type Clock struct {
	hour int
	min  int
}

func New(h, m int) Clock {
	return initializeClock(h, m) //Clock{hour: newH, min: newM}
	//panic("Please implement the New function")
}

func initializeClock(h, m int) Clock {
	extraH := m / 60
	newM := m % 60
	if newM < 0 {
		extraH -= 1
		newM += 60
	}

	newH := (h + extraH) % 24
	if newH < 0 {
		newH += 24
	}
	//fmt.Println("INIT: Incoming h, m: ", h, m, ", extraH: ", extraH, ", newH:", newH, ", newM: ", newM)
	return Clock{hour: newH, min: newM}
}

func (c Clock) Add(m int) Clock {

	return initializeClock(c.hour, c.min+m)

	//panic("Please implement the Add function")
	// extraH = m / 60
	// c.min += m
	// extraH := c.min / 60
	// c.hour = (c.hour + extraH) % 24
	// c.min = c.min % 60
	// return c
}

func (c Clock) Subtract(m int) Clock {
	return initializeClock(c.hour, c.min-m)
	//panic("Please implement the Subtract function")
	// c.min -= m
	// if c.min < 0 {
	// 	extraH := c.min / 60
	// 	c.hour -= extraH - 1
	// 	c.min %= 60
	// 	c.hour %= 24
	// }
	// return c
}

func (c Clock) String() string {
	return fmt.Sprintf("%02d:%02d", c.hour%24, c.min%60)
	//panic("Please implement the String function")
}
