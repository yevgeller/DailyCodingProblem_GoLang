package bob

import (
	"strings"
)

func Hey(remark string) string {
	if len(strings.TrimSpace(remark)) == 0 {
		return "Fine. Be that way!"
	}

	isQuestion := strings.HasSuffix(strings.TrimSpace(remark), "?")
	isAllUpper := strings.ToUpper(remark) == remark && strings.ToLower(remark) != remark

	if isQuestion && !isAllUpper {
		return "Sure."
	}

	if !isQuestion && isAllUpper {
		return "Whoa, chill out!"
	}

	if isQuestion && isAllUpper {
		return "Calm down, I know what I'm doing!"
	}

	return "Whatever."
}
