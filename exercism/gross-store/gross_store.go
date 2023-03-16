package gross

import "fmt"

// Units stores the Gross Store unit measurements.
func Units() map[string]int {
	newMap := make(map[string]int)
	newMap["quarter_of_a_dozen"] = 3
	newMap["half_of_a_dozen"] = 6
	newMap["small_gross"] = 120
	newMap["dozen"] = 12
	newMap["gross"] = 144
	newMap["great_gross"] = 1728
	return newMap
}

// NewBill creates a new bill.
func NewBill() map[string]int {
	return make(map[string]int)
}

// value, exists := foo["baz"]
// If the key "baz" does not exist,
// value: 0; exists: false
// AddItem adds an item to customer bill.
func AddItem(bill, units map[string]int, item, unit string) bool {
	itemVal, unitExists := units[unit]
	if unitExists == false {
		return false
	}

	_, itemExist := bill[item]

	if itemExist == false {
		bill[item] = itemVal
		return true
	}

	bill[item] += itemVal
	return true
}

// RemoveItem removes an item from customer bill.
// To delete an item from a map, you can use
//
//	delete(foo, "bar")
func RemoveItem(bill, units map[string]int, item, unit string) bool {
	fmt.Println(bill)
	fmt.Println(item, unit)
	itemVal, unitExists := units[unit]
	if unitExists == false {
		fmt.Printf("unit %s does not exist", unit)
		fmt.Println("------ next ------")

		return false
	}

	currItemValue, itemExist := bill[item]
	if itemExist == false {
		fmt.Printf("item %s does not exist", item)
		fmt.Println("------ next ------")
		return false
	}

	diff := currItemValue - itemVal
	fmt.Printf("diff is %d", diff)

	if diff < 0 {
		fmt.Println("-- Negative diff --")
		fmt.Println("------ next ------")
		return false
	}

	if diff == 0 {
		fmt.Println("-- diff is 0, removing item --")
		fmt.Println("------ next ------")
		delete(bill, "item")
		return true
	}

	fmt.Println("-- Positive diff, decreasing item --")
	fmt.Println("------ next ------")
	bill[item] = diff
	return true
}

// GetItem returns the quantity of an item that the customer has in his/her bill.
func GetItem(bill map[string]int, item string) (int, bool) {
	currItemValue, itemExist := bill[item]
	if itemExist == false {
		return 0, false
	}

	return currItemValue, true
}
