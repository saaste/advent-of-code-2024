package puzzle

import (
	"fmt"

	"github.com/saaste/advent-of-code-2024/pkg/input"
)

type Day9 struct{}

type Disk struct {
	DiskImage        []int
	EmptySpaceBlocks int
}

func (d *Disk) AddBlock(block []int) {
	for i := 0; i < len(block); i++ {
		if block[0] == -1 {
			d.EmptySpaceBlocks++
			d.DiskImage = append(d.DiskImage, -1)
		} else {
			d.DiskImage = append(d.DiskImage, block[0])
		}
	}
}

func (d *Disk) DefragByBlock() {
	// Move number from right to empty spaces
	i := len(d.DiskImage) - 1
	for d.EmptySpaceBlocks > 0 {
		d.MoveBlockToEmpty(i)
		i--
	}
}

func (d *Disk) DefragByFile() {
	// Move blocks from right to suitable empty spaces
	i := len(d.DiskImage) - 1
	for i > 0 {
		value := d.DiskImage[i]

		// Empty space. Skip
		if value == -1 {
			i--
			continue
		}

		fileStart := i
		fileLength := 1

		// Find start of the file
		for d.DiskImage[fileStart-1] == value {
			fileStart--
			fileLength++
			if fileStart <= 0 {
				break
			}
		}

		// Set index to file start
		i = fileStart

		// Find suitable space for the file
		j := 0
		for j < fileStart {
			emptyValue := d.DiskImage[j]

			// If this is not an empty space, continue
			if emptyValue != -1 {
				j++
				continue
			}

			// Check the length of the empty
			emptyStart := j
			emptyLength := 1

			for d.DiskImage[emptyStart+emptyLength] == -1 {
				emptyLength++
			}

			// Empty space is not long enough. Continue
			if emptyLength < fileLength {
				j++
				continue
			}

			// We have an empty space that is long enough for the file.
			for loop := 0; loop < fileLength; loop++ {
				// Replace empty spaces with file
				d.SwitchValues(fileStart+loop, emptyStart+loop)
			}
			// Exit the empty search loop
			break
		}

		i--
	}
}

func (d *Disk) MoveBlockToEmpty(i int) {
	value := d.DiskImage[i]
	if value == -1 {
		d.EmptySpaceBlocks--
		return
	}

	// Find first empty space
	// Could be optimized by storing the index for the next search
	emptyIndex := 0
	for emptyIndex < i {
		val := d.DiskImage[emptyIndex]
		if val == -1 {
			break
		}
		emptyIndex++
	}

	d.DiskImage[emptyIndex] = value
	d.DiskImage[i] = -1
	d.EmptySpaceBlocks--
}

func (d *Disk) SwitchValues(from, to int) {
	fromValue := d.DiskImage[from]
	toValue := d.DiskImage[to]
	d.DiskImage[to] = fromValue
	d.DiskImage[from] = toValue
}

func (d *Disk) CalculateCheckSumFromImage() int {
	checksum := 0
	for i, value := range d.DiskImage {
		if value == -1 {
			continue
		}

		checksum += i * value
	}
	return checksum
}

func (d Day9) Step1(puzzleInput string) string {
	unorderedDisk := parseDisk(puzzleInput)
	unorderedDisk.DefragByBlock()
	return fmt.Sprintf("%d", unorderedDisk.CalculateCheckSumFromImage())
}

func (d Day9) Step2(puzzleInput string) string {
	unorderedDisk := parseDisk(puzzleInput)
	unorderedDisk.DefragByFile()
	checksum := unorderedDisk.CalculateCheckSumFromImage()
	return fmt.Sprintf("%d", checksum)
}

func parseDisk(puzzleInput string) Disk {
	numbers := input.IntSlice(puzzleInput)
	disk := Disk{}

	// Split input into files and empty blocks
	var fileID int = 0
	for i, value := range numbers {
		if value == 0 {
			continue
		}
		// Even positions are files
		if i%2 == 0 {
			content := make([]int, value)
			for j := 0; j < value; j++ {
				content[j] = fileID
			}
			disk.AddBlock(content)
			fileID++
		} else {
			content := make([]int, value)
			for j := 0; j < value; j++ {
				content[j] = -1 // -1 = empty space
			}
			disk.AddBlock(content)
		}
	}

	return disk
}
