package puzzle

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/saaste/advent-of-code-2024/pkg/input"
)

type Day11 struct{}

func (d Day11) Step1(puzzleInput string) string {
	stones := input.SpaceSeparatedInts(puzzleInput)
	stoneCount := blink(stones, 25)
	return fmt.Sprintf("%d", stoneCount)
}

func (d Day11) Step2(puzzleInput string) string {
	stones := input.SpaceSeparatedInts(puzzleInput)
	stoneCount := blink(stones, 75)
	return fmt.Sprintf("%d", stoneCount)
}

func blink(stones []int64, blinks int) int64 {
	precalculatedCounts := int64(0)

	for i := 0; i < blinks; i++ {
		updatedStones := make([]int64, 0)
		for _, stone := range stones {
			stoneDigits := fmt.Sprintf("%d", stone)
			stoneDigitsLength := len(stoneDigits)
			if stone == 0 && i < 35 {
				// Use normal logic for the first 35 blinks
				updatedStones = append(updatedStones, 1)
			} else if stone >= 0 && stone <= 9 && i >= 35 {
				// Remove single digit numbers from the list and
				// sum precalculated values
				blinksLeft := blinks - i
				if stone == 0 {
					precalculatedCounts += precalculatedZeros[blinksLeft]
				} else if stone == 1 {
					precalculatedCounts += precalculatedZeros[blinksLeft+1]
				} else if stone == 2 {
					precalculatedCounts += precalculatedTwos[blinksLeft]
				} else if stone == 3 {
					precalculatedCounts += precalculatedThrees[blinksLeft]
				} else if stone == 4 {
					precalculatedCounts += precalculatedFours[blinksLeft]
				} else if stone == 5 {
					precalculatedCounts += precalculatedFives[blinksLeft]
				} else if stone == 6 {
					precalculatedCounts += precalculatedSixes[blinksLeft]
				} else if stone == 7 {
					precalculatedCounts += precalculatedSevents[blinksLeft]
				} else if stone == 8 {
					precalculatedCounts += precalculatedEights[blinksLeft]
				} else {
					precalculatedCounts += precalculatedNines[blinksLeft]
				}
			} else if stoneDigitsLength%2 == 0 {
				// Split even digits stones into two stones
				left := stoneDigits[:int32(stoneDigitsLength)/2]
				right := stoneDigits[stoneDigitsLength/2:]

				leftNumber, err := strconv.ParseInt(left, 10, 64)
				if err != nil {
					log.Fatalf("unable to parse %s to int: %v", left, err)
				}

				right = strings.TrimPrefix(right, "0")
				if right == "" {
					right = "0"
				}
				rightNumber, err := strconv.ParseInt(right, 10, 64)
				if err != nil {
					log.Fatalf("unable to parse %s to int: %v", right, err)
				}

				if leftNumber < 0 || rightNumber < 0 {
					log.Fatalln("negative")
				}

				updatedStones = append(updatedStones, leftNumber, rightNumber)
			} else {
				// Otherwise replace stone with old number * 2024
				if stone*2024 < 0 {
					log.Fatalln("negative")
				}
				updatedStones = append(updatedStones, stone*2024)
			}
		}
		stones = updatedStones
	}
	return int64(len(stones)) + precalculatedCounts
}

var precalculatedZeros = map[int]int64{
	1:  1,
	2:  1,
	3:  2,
	4:  4,
	5:  4,
	6:  7,
	7:  14,
	8:  16,
	9:  20,
	10: 39,
	11: 62,
	12: 81,
	13: 110,
	14: 200,
	15: 328,
	16: 418,
	17: 667,
	18: 1059,
	19: 1546,
	20: 2377,
	21: 3572,
	22: 5602,
	23: 8268,
	24: 12343,
	25: 19778,
	26: 29165,
	27: 43726,
	28: 67724,
	29: 102131,
	30: 156451,
	31: 234511,
	32: 357632,
	33: 549949,
	34: 819967,
	35: 1258125,
	36: 1916299,
	37: 2886408,
	38: 4414216,
	39: 6669768,
	40: 10174278,
	41: 15458147,
	42: 23333796,
	43: 35712308,
	44: 54046805,
	45: 81997335,
	46: 125001266,
}

var precalculatedTwos = map[int]int64{
	1:  1,
	2:  2,
	3:  4,
	4:  4,
	5:  6,
	6:  12,
	7:  16,
	8:  19,
	9:  30,
	10: 57,
	11: 92,
	12: 111,
	13: 181,
	14: 295,
	15: 414,
	16: 661,
	17: 977,
	18: 1501,
	19: 2270,
	20: 3381,
	21: 5463,
	22: 7921,
	23: 11819,
	24: 18712,
	25: 27842,
	26: 42646,
	27: 64275,
	28: 97328,
	29: 150678,
	30: 223730,
	31: 343711,
	32: 525238,
	33: 784952,
	34: 1208065,
	35: 1824910,
	36: 2774273,
	37: 4230422,
	38: 6365293,
	39: 9763578,
	40: 14777945,
	41: 22365694,
	42: 34205743,
	43: 51643260,
}

var precalculatedThrees = map[int]int64{
	1:  1,
	2:  2,
	3:  4,
	4:  4,
	5:  5,
	6:  10,
	7:  16,
	8:  26,
	9:  35,
	10: 52,
	11: 79,
	12: 114,
	13: 202,
	14: 294,
	15: 401,
	16: 642,
	17: 987,
	18: 1556,
	19: 2281,
	20: 3347,
	21: 5360,
	22: 7914,
	23: 12116,
	24: 18714,
	25: 27569,
	26: 42628,
	27: 64379,
	28: 98160,
	29: 150493,
	30: 223231,
	31: 344595,
	32: 524150,
	33: 788590,
	34: 1210782,
	35: 1821382,
	36: 2779243,
	37: 4230598,
	38: 6382031,
	39: 9778305,
	40: 14761601,
}

var precalculatedFours = map[int]int64{
	1:  1,
	2:  2,
	3:  4,
	4:  4,
	5:  4,
	6:  8,
	7:  16,
	8:  27,
	9:  30,
	10: 47,
	11: 82,
	12: 115,
	13: 195,
	14: 269,
	15: 390,
	16: 637,
	17: 951,
	18: 1541,
	19: 2182,
	20: 3204,
	21: 5280,
	22: 7721,
	23: 11820,
	24: 17957,
	25: 26669,
	26: 41994,
	27: 62235,
	28: 95252,
	29: 146462,
	30: 216056,
	31: 336192,
	32: 508191,
	33: 766555,
	34: 1178119,
	35: 1761823,
	36: 2709433,
	37: 4110895,
	38: 6188994,
	39: 9515384,
	40: 14316637,
}

var precalculatedFives = map[int]int64{
	1:  1,
	2:  1,
	3:  2,
	4:  4,
	5:  8,
	6:  8,
	7:  11,
	8:  22,
	9:  32,
	10: 45,
	11: 67,
	12: 109,
	13: 163,
	14: 223,
	15: 383,
	16: 597,
	17: 808,
	18: 1260,
	19: 1976,
	20: 3053,
	21: 4529,
	22: 6675,
	23: 10627,
	24: 15847,
	25: 23822,
	26: 37090,
	27: 55161,
	28: 84208,
	29: 128121,
	30: 194545,
	31: 298191,
	32: 444839,
	33: 681805,
	34: 1042629,
	35: 1565585,
	36: 2396146,
	37: 3626619,
	38: 5509999,
	39: 8396834,
	40: 12678459,
}
var precalculatedSixes = map[int]int64{
	1:  1,
	2:  1,
	3:  2,
	4:  4,
	5:  8,
	6:  8,
	7:  11,
	8:  22,
	9:  32,
	10: 54,
	11: 68,
	12: 103,
	13: 183,
	14: 250,
	15: 401,
	16: 600,
	17: 871,
	18: 1431,
	19: 2033,
	20: 3193,
	21: 4917,
	22: 7052,
	23: 11371,
	24: 16815,
	25: 25469,
	26: 39648,
	27: 57976,
	28: 90871,
	29: 136703,
	30: 205157,
	31: 319620,
	32: 473117,
	33: 727905,
	34: 1110359,
	35: 1661899,
	36: 2567855,
	37: 3849988,
	38: 5866379,
	39: 8978479,
	40: 13464170,
}

var precalculatedSevents = map[int]int64{
	1:  1,
	2:  1,
	3:  2,
	4:  4,
	5:  8,
	6:  8,
	7:  11,
	8:  22,
	9:  32,
	10: 52,
	11: 72,
	12: 106,
	13: 168,
	14: 242,
	15: 413,
	16: 602,
	17: 832,
	18: 1369,
	19: 2065,
	20: 3165,
	21: 4762,
	22: 6994,
	23: 11170,
	24: 16509,
	25: 25071,
	26: 39034,
	27: 57254,
	28: 88672,
	29: 134638,
	30: 203252,
	31: 312940,
	32: 465395,
	33: 716437,
	34: 1092207,
	35: 1637097,
	36: 2519878,
	37: 3794783,
	38: 5771904,
	39: 8814021,
	40: 13273744,
}

var precalculatedEights = map[int]int64{
	1:  1,
	2:  1,
	3:  2,
	4:  4,
	5:  7,
	6:  7,
	7:  11,
	8:  22,
	9:  31,
	10: 48,
	11: 69,
	12: 103,
	13: 161,
	14: 239,
	15: 393,
	16: 578,
	17: 812,
	18: 1322,
	19: 2011,
	20: 3034,
	21: 4580,
	22: 6798,
	23: 10738,
	24: 16018,
	25: 24212,
	26: 37525,
	27: 55534,
	28: 85483,
	29: 130183,
	30: 196389,
	31: 301170,
	32: 450896,
	33: 691214,
	34: 1054217,
	35: 1583522,
	36: 2428413,
	37: 3669747,
	38: 5573490,
	39: 8505207,
	40: 12835708,
}

var precalculatedNines = map[int]int64{
	1:  1,
	2:  1,
	3:  2,
	4:  4,
	5:  8,
	6:  8,
	7:  11,
	8:  22,
	9:  32,
	10: 54,
	11: 70,
	12: 103,
	13: 183,
	14: 262,
	15: 419,
	16: 586,
	17: 854,
	18: 1468,
	19: 2131,
	20: 3216,
	21: 4888,
	22: 7217,
	23: 11617,
	24: 17059,
	25: 25793,
	26: 40124,
	27: 58820,
	28: 92114,
	29: 139174,
	30: 208558,
	31: 322818,
	32: 480178,
	33: 740365,
	34: 1126352,
	35: 1685448,
	36: 2602817,
	37: 3910494,
	38: 5953715,
	39: 9102530,
	40: 13675794,
}
