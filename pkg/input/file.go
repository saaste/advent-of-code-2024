package input

import (
	"fmt"
	"log"
	"os"
)

func ReadFile(day int) string {
	filename := fmt.Sprintf("inputs/%d.txt", day)
	b, err := os.ReadFile(filename)
	if err != nil {
		log.Fatalf("failed to read file %s: %v", filename, err)
	}
	return string(b)
}
