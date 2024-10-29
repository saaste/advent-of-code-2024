package main

import (
	"flag"
	"log"

	"github.com/saaste/advent-of-code-2024/pkg/puzzle"
)

func main() {
	day := flag.Int("d", 0, "day of month (1-25)")
	step := flag.Int("s", 0, "puzzle step (1,2)")
	validate := flag.Bool("v", false, "puzzle step (1,2)")

	flag.Parse()

	if *day < 1 || *day > 25 {
		log.Fatalln("Invalid day. Must be between 1 and 25.")
	}

	if *validate {
		puzzle.ValidatePuzzle(*day, *step)
	} else {
		puzzle.RunPuzzle(*day, *step)
	}

}
