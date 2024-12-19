package puzzle

import (
	"fmt"
	"strings"

	"github.com/saaste/advent-of-code-2024/pkg/input"
)

type Day19 struct{}

func (d Day19) Step1(puzzleInput string) string {
	towelStatus := parseTowelStatus(puzzleInput)
	possibleCount := 0
	for _, request := range towelStatus.Requests {
		result := findTowelCombination(request, towelStatus.Towels, make([]string, 0))
		if len(result) > 0 {
			possibleCount++
		}
	}
	return fmt.Sprintf("%d", possibleCount)
}

func (d Day19) Step2(puzzleInput string) string {
	towelStatus := parseTowelStatus(puzzleInput)
	possibleCount := 0
	for _, request := range towelStatus.Requests {
		possibleCount += findAllTowelCombinations(request, towelStatus.TowelsMap, make(map[string]int), 0)
	}
	return fmt.Sprintf("%d", possibleCount)
}

func findTowelCombination(request string, towels []string, currentTowels []string) []string {
	currentStripes := strings.Join(currentTowels, "")
	// We found the combination
	if request == "" {
		return currentTowels
	}

	// Stripes we still need to find
	newRequest, _ := strings.CutPrefix(request, currentStripes)

	// Try each towel and try to find a match that could build the request
	for _, towel := range towels {
		// If towel is too long, it does not work
		if len(towel) > len(newRequest) {
			continue
		}

		// If stripes don't start with the towel, it does not work
		if !strings.HasPrefix(newRequest, towel) {
			continue
		}

		// This could be a possible combination, let's give it a go.
		newSuffix, _ := strings.CutPrefix(newRequest, towel)
		newCurrentTowels := append(currentTowels, towel)

		result := findTowelCombination(newSuffix, towels, newCurrentTowels)
		if result != nil {
			return result
		}
	}

	return nil
}

func findAllTowelCombinations(request string, towels map[string]bool, cache map[string]int, depth int) int {
	currentDepthCount := 0

	prefix := request[len(request)-depth-1 : len(request)-depth]
	suffix := request[len(request)-depth:]

	for {
		if _, prefixFound := towels[prefix]; prefixFound {
			if suffixCount, suffixFound := cache[suffix]; suffixFound {
				currentDepthCount += suffixCount
			}

			if len(suffix) == 0 {
				currentDepthCount += 1
			}
		}

		if len(suffix) > 0 {
			prefix += suffix[0:1]
			suffix = suffix[1:]
		} else {
			break
		}
	}

	cache[prefix] = currentDepthCount

	if prefix == request {
		return currentDepthCount
	} else {
		return findAllTowelCombinations(request, towels, cache, depth+1)
	}
}

type TowelStatus struct {
	Towels    []string
	Requests  []string
	TowelsMap map[string]bool
}

func parseTowelStatus(puzzleInput string) *TowelStatus {
	lines := input.EachLineAsString(puzzleInput)
	towels := strings.Split(lines[0], ", ")
	requests := lines[2:]
	towelsMap := make(map[string]bool)

	for _, towel := range towels {
		towelsMap[towel] = true
	}

	return &TowelStatus{
		Towels:    towels,
		Requests:  requests,
		TowelsMap: towelsMap,
	}

}
