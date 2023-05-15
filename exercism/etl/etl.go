package etl

import "strings"

func Transform(in map[int][]string) map[string]int {
	var ret = make(map[string]int)

	for k, v := range in {
		for _, b := range v {
			ret[strings.ToLower(b)] = k
		}
	}
	return ret
}
