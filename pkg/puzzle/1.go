package puzzle

import (
	"fmt"
	"log"
	"sort"
	"strconv"
	"strings"

	"github.com/saaste/advent-of-code-2024/pkg/input"
)

type Day1 struct{}

func (d Day1) Step1(puzzleInput string) string {
	list1, list2 := makeSortedLists(puzzleInput)

	distanceSum := 0
	for i, value1 := range list1 {
		value2 := list2[i]
		distanceSum += max(value1, value2) - min(value1, value2)
	}

	return fmt.Sprintf("%d", distanceSum)
}

func (d Day1) Step2(puzzleInput string) string {
	list1, list2 := makeSortedLists(puzzleInput)

	startCheckIndex := 0
	previousNumber := 0
	timesNumberIsFound := 0

	sum := 0

	for _, a := range list1 {
		// If new number is the same as previous number, we can use
		// the previous result without checking
		if a == previousNumber {
			sum += a * timesNumberIsFound
			continue
		}

		// Otherwise reset counter
		timesNumberIsFound = 0

		// Find how many times the number is in list 2
		for i := startCheckIndex; i < len(list2); i++ {
			if list2[i] == a {
				// Number found
				timesNumberIsFound++
			} else if list2[i] > a {
				// Numbers in list 2 are too big, we can stop checking and
				// mark the location where we stopped
				startCheckIndex = i
				break
			}
		}

		// Increase the sum
		sum += a * timesNumberIsFound

		// Set this number as the previous number
		previousNumber = a

	}

	return fmt.Sprintf("%d", sum)
}

func makeSortedLists(puzzleInput string) ([]int, []int) {
	list1 := make([]int, 0)
	list2 := make([]int, 0)

	inputLines := input.EachLineAsString(puzzleInput)
	for _, line := range inputLines {
		lineParts := strings.Split(line, "   ")
		location1, err := strconv.ParseInt(lineParts[0], 10, 32)
		if err != nil {
			log.Fatalf("failed to parse location 1 as integer: %v", err)
		}
		list1 = append(list1, int(location1))

		location2, err := strconv.ParseInt(lineParts[1], 10, 32)
		if err != nil {
			log.Fatalf("failed to parse location 2 as integer: %v", err)
		}
		list2 = append(list2, int(location2))
	}

	sort.Ints(list1)
	sort.Ints(list2)

	return list1, list2
}
