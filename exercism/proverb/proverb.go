package proverb

import "fmt"

func Proverb(rhyme []string) []string {
	var ret []string
	ln := len(rhyme)
	if ln == 0 {
		return ret
	}
	a := 0

	for a < (ln - 1) {
		ret = append(ret, fmt.Sprintf("For want of a %s the %s was lost.", rhyme[a], rhyme[a+1]))
		a++
	}

	ret = append(ret, fmt.Sprintf("And all for the want of a %s.", rhyme[0]))
	return ret
}
