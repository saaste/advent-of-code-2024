package puzzle

import (
	"fmt"
	"log"
	"math"
	"strconv"
	"strings"

	"github.com/saaste/advent-of-code-2024/pkg/input"
)

type Day2 struct{}

func (d Day2) Step1(puzzleInput string) string {
	reports := parseReports(puzzleInput)
	validReports := 0

	for _, report := range reports {
		if validateReport(report) {
			validReports++
		}
	}

	return fmt.Sprintf("%d", validReports)
}

func (d Day2) Step2(puzzleInput string) string {
	reports := parseReports(puzzleInput)
	validReports := 0

	for _, report := range reports {
		if validateReport(report) {
			validReports++
			continue
		}

		// Report is invalid. Test variations
		for i := 0; i < len(report); i++ {
			variation := createReportVariation(report, i)
			if validateReport(variation) {
				validReports++
				break
			}
		}

	}

	return fmt.Sprintf("%d", validReports)
}

func parseReports(puzzleInput string) [][]int {
	lines := input.EachLineAsString(puzzleInput)
	reports := make([][]int, 0)

	for _, line := range lines {
		numbers := make([]int, 0)
		strNumbers := strings.Split(line, " ")
		for _, strNumber := range strNumbers {
			number, err := strconv.ParseInt(strings.TrimSpace(strNumber), 10, 32)
			if err != nil {
				log.Fatalf("failed parsing '%s' as integer: %v", strNumber, err)
			}

			numbers = append(numbers, int(number))
		}
		reports = append(reports, numbers)
	}
	return reports
}

func validateReport(report []int) bool {
	direction := report[0] - report[1]
	for i := 0; i < len(report)-1; i++ {
		a := report[i]
		b := report[i+1]

		if direction < 0 && a >= b {
			return false
		}

		if direction > 0 && a <= b {
			return false
		}

		diff := math.Abs(float64(a) - float64(b))
		if diff < 1 || diff > 3 {
			return false
		}
	}
	return true
}

func createReportVariation(report []int, removeIndex int) []int {
	variation := make([]int, 0)
	variation = append(variation, report[:removeIndex]...)
	return append(variation, report[removeIndex+1:]...)
}
