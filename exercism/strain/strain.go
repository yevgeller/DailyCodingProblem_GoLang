package strain

type Ints []int
type Lists [][]int
type Strings []string

func (i Ints) Keep(filter func(int) bool) Ints {
	ret := []int(nil)

	for _, i := range i {
		if filter(i) {
			ret = append(ret, i)
		}
	}
	return ret
}

func (i Ints) Discard(filter func(int) bool) Ints {
	ret := []int(nil)

	for _, i := range i {
		if !filter(i) {
			ret = append(ret, i)
		}
	}
	return ret
}

func (l Lists) Keep(filter func([]int) bool) Lists {
	//panic("Please implement the Keep function")
	ret := [][]int(nil)

	for _, i := range l {
		if filter(i) {
			ret = append(ret, i)
		}
	}
	return ret
}

func (s Strings) Keep(filter func(string) bool) Strings {
	//panic("Please implement the Keep function")
	ret := []string(nil)

	for _, i := range s {
		if filter(i) {
			ret = append(ret, i)
		}
	}
	return ret
}

// func filter(int) bool {
// 	return
// }
