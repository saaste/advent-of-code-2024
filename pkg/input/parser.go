package input

import (
	"log"
	"strconv"
	"strings"
)

func normalize(data string) string {
	return strings.ReplaceAll(data, "\r\n", "\n")
}

func EachLineAsString(data string) []string {
	data = normalize(data)
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

func IntSlice(data string) []int {
	ints := make([]int, 0)
	for _, val := range strings.Split(data, "") {
		intVal, err := strconv.ParseInt(val, 10, 32)
		if err != nil {
			log.Fatalf("unable to parse %s to int", val)
		}
		ints = append(ints, int(intVal))
	}
	return ints
}

func CharacterGrid(data string) [][]string {
	grid := make([][]string, 0)
	rows := EachLineAsString(data)
	for _, row := range rows {
		grid = append(grid, strings.Split(row, ""))
	}
	return grid
}

func SplitByEmptyLine(data string) []string {
	data = normalize(data)
	return strings.Split(data, "\n\n")
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
