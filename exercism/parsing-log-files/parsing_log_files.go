package parsinglogfiles

import (
	"regexp"
)

func IsValidLine(text string) bool {
	return regexp.MustCompile(`^\[(TRC|DBG|INF|WRN|ERR||FTL)\]`).MatchString(text)
}

func SplitLogLine(text string) []string {
	return regexp.MustCompile(`<[*~=-]*>`).Split(text, -1)
}

func CountQuotedPasswords(lines []string) int {
	re := regexp.MustCompile(`(?i)".*password.*"`)
	count := 0
	for _, line := range lines {
		if re.MatchString(line) {
			count++
		}
	}
	return count
}

func RemoveEndOfLineText(text string) string {
	//panic("Please implement the RemoveEndOfLineText function")
	return regexp.MustCompile(`end-of-line\d+`).ReplaceAllString(text, "")
}

func TagWithUserName(lines []string) []string {
	//panic("Please implement the TagWithUserName function")
	var tagged []string
	re := regexp.MustCompile(`User\s+([A-Za-z0-9]*)`)
	for _, line := range lines {
		matches := re.FindStringSubmatch(line)
		if matches != nil {
			tagged = append(tagged, "[USR] "+matches[1]+" "+line)
		} else {
			tagged = append(tagged, line)
		}
	}
	return tagged
}
