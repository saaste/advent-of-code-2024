package input

import (
	"fmt"
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

func SpaceSeparatedInts(data string) []int64 {
	values := strings.Split(data, " ")
	return stringsToInt64s(values)
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

func StringAsInt64(s string) int64 {
	value, err := strconv.ParseInt(s, 10, 64)
	if err != nil {
		log.Fatalf("unable to parse %s as int64", s)
	}
	return value
}

func CharacterGrid(data string) [][]string {
	grid := make([][]string, 0)
	rows := EachLineAsString(data)
	for _, row := range rows {
		grid = append(grid, strings.Split(row, ""))
	}
	return grid
}

func CharacterMap(data string) map[string]string {
	result := make(map[string]string)
	rows := EachLineAsString(data)
	for y, row := range rows {
		for x, val := range row {
			key := fmt.Sprintf("%d,%d", x, y)
			result[key] = string(val)
		}
	}
	return result
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

func stringsToInt64s(values []string) []int64 {
	result := make([]int64, 0)
	for _, line := range values {
		value, err := strconv.ParseInt(line, 10, 64)
		if err != nil {
			log.Fatalf("failed to parse %s to int: %v", line, err)
		}
		result = append(result, value)
	}
	return result
}
