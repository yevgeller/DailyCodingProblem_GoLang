package partyrobot

import "fmt"

// Welcome greets a person by name.
func Welcome(name string) string {
	return fmt.Sprintf("Welcome to my party, %s!", name)
}

// HappyBirthday wishes happy birthday to the birthday person and exclaims their age.
func HappyBirthday(name string, age int) string {
	return fmt.Sprintf("Happy birthday %s! You are now %d years old!", name, age)
}

// AssignTable assigns a table to each guest.
func AssignTable(name string, table int, neighbor, direction string, distance float64) string {
	// Welcome to my party, Christiane!
	// You have been assigned to table 027. Your table is on the left, exactly 23.8 meters from here.
	// You will be sitting next to Frank.
	return Welcome(name) + "\n" +
		fmt.Sprintf("You have been assigned to table %03d. ", table) +
		fmt.Sprintf("Your table is %s, exactly %.1f meters from here.", direction, distance) + "\n" +
		fmt.Sprintf("You will be sitting next to %s.", neighbor)
}

//remove comments
