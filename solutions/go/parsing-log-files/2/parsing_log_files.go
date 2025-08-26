package parsinglogfiles

import (
	"fmt"
	"regexp"
)

func IsValidLine(text string) bool {
	re := regexp.MustCompile(`^\[(TRC|DBG|INF|WRN|ERR|FTL)\]`)

	return re.MatchString(text)
}

func SplitLogLine(text string) []string {
	re := regexp.MustCompile(`<[~*=-]*>`)

	return re.Split(text, -1)
}

func CountQuotedPasswords(lines []string) int {
	re := regexp.MustCompile(`(?i)".*password.*"`)
	matchedCount := 0

	for _, l := range lines {
		if re.MatchString(l) {
			matchedCount++
		}
	}

	return matchedCount
}

func RemoveEndOfLineText(text string) string {
	re := regexp.MustCompile(`end-of-line\d*`)

	return re.ReplaceAllString(text, "")
}

func TagWithUserName(lines []string) []string {
	re := regexp.MustCompile(`\bUser\s+(\w+)`)

	linesWithName := make([]string, 0)

	for _, v := range lines {
		name := re.FindStringSubmatch(v)
		if name != nil {
			v = fmt.Sprintf("[USR] %s %s", name[1], v)
		}

		linesWithName = append(linesWithName, v)
	}

	return linesWithName
}
