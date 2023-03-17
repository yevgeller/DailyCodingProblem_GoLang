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
func RemoveItem(bill, units map[string]int, item, unit string) bool {
	fmt.Println(bill)
	fmt.Println(item, unit)
	itemVal, unitExists := units[unit]
	if unitExists == false {
		fmt.Printf("unit %s does not exist", unit)
		fmt.Printf("\n\n")

		return false
	}

	currItemValue, itemExist := bill[item]
	if itemExist == false {
		fmt.Printf("item %s does not exist", item)
		fmt.Printf("\n\n")
		return false
	}

	diff := currItemValue - itemVal
	fmt.Printf("diff is %d", diff)

	if diff < 0 {
		fmt.Println("-- Negative diff --")
		fmt.Printf("\n\n")

		return false
	}

	if diff == 0 {
		fmt.Println("-- diff is 0, removing item --")
		fmt.Printf("\n\n")
		delete(bill, item)
		//fmt.Println(bill)
		return true
	}

	fmt.Println("-- Positive diff, decreasing item --")
	fmt.Printf("\n\n")
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
