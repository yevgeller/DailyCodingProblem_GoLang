package etl

import "strings"

func Transform(in map[int][]string) map[string]int {
	//panic("Please implement the Transform function")
	var ret = make(map[string]int)

	for k, v := range in {
		// k - key, v -value
		for _, b := range v {
			ret[strings.ToLower(b)] = k
		}
	}
	return ret
	//return make(map[string]int)
}
