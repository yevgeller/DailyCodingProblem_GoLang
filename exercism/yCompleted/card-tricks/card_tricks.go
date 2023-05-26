package cards

// FavoriteCards returns a slice with the cards 2, 6 and 9 in that order.
func FavoriteCards() []int {
	ret := []int{2}
	ret = append(ret, 6, 9)
	return ret
}

// GetItem retrieves an item from a slice at given position.
// If the index is out of range, we want it to return -1.
func GetItem(slice []int, index int) int {
	if len(slice)-1 < index || index < 0 {
		return -1
	}
	return slice[index]
	//   for i, v := range slice {
	//     if v == num {
	//         fmt.Println(num, "found at index", i, "in", nums)
	//         return
	//     }
	// }
}

// SetItem writes an item to a slice at given position overwriting an existing value.
// If the index is out of range the value needs to be appended.
func SetItem(slice []int, index, value int) []int {
	if len(slice)-1 < index || index < 0 {
		return append(slice, value)
	}
	slice[index] = value
	return slice
}

// PrependItems adds an arbitrary number of values at the front of a slice.
func PrependItems(slice []int, values ...int) []int {
	return append(values, slice...)
}

// RemoveItem removes an item from a slice by modifying the existing slice.
func RemoveItem(slice []int, index int) []int {
	if len(slice)-1 < index || index < 0 {
		return slice
	}
	if index == 0 {
		return slice[1:]
	}
	if index == len(slice) {
		return slice[0:]
	}
	return append(slice[0:index], slice[index+1:]...)
}
