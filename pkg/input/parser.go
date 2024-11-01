package input

import (
	"log"
	"strconv"
	"strings"
)

func EachLineAsString(data string) []string {
	return strings.Split(data, "\n")
}

func EachLineAsInt(data string) []int32 {
	lines := EachLineAsString(data)
	return stringsToInts(lines)
}

func CommaSeparatedInts(data string) []int32 {
	values := strings.Split(data, ",")
	return stringsToInts(values)
}

type StringGroup struct {
	GroupIdentifier string
	Lines           []string
}

func GroupedStrings(data string, groupIdentifier string) []StringGroup {
	lines := EachLineAsString(data)
	groups := make([]StringGroup, 0)
	currentGroup := StringGroup{
		Lines: make([]string, 0),
	}
	for _, line := range lines {
		if strings.TrimSpace(line) == "" {
			continue
		}

		if strings.Contains(line, groupIdentifier) {
			if currentGroup.GroupIdentifier == "" {
				currentGroup.GroupIdentifier = line
			} else {
				groups = append(groups, currentGroup)
				currentGroup = StringGroup{
					GroupIdentifier: line,
					Lines:           make([]string, 0),
				}
			}
			continue
		}

		currentGroup.Lines = append(currentGroup.Lines, line)
	}
	groups = append(groups, currentGroup)
	return groups
}

func stringsToInts(values []string) []int32 {
	result := make([]int32, 0)
	for _, line := range values {
		value, err := strconv.ParseInt(line, 10, 32)
		if err != nil {
			log.Fatalf("failed to parse %s to int: %v", line, err)
		}
		result = append(result, int32(value))
	}
	return result
}
