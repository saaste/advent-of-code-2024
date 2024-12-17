package puzzle

import (
	"fmt"
	"log"
	"math"
	"slices"
	"strings"

	"github.com/saaste/advent-of-code-2024/pkg/input"
)

type Day17 struct{}

func (d Day17) Step1(puzzleInput string) string {
	computer := parseComputer(puzzleInput)
	computer.Run()
	return computer.ConcatenateOutput()
}

func (d Day17) Step2(puzzleInput string) string {
	computer := parseComputer(puzzleInput)
	return fmt.Sprintf("%d", computer.FindSelfReplicatingA(0, 0))
}

func parseComputer(puzzleInput string) *Computer {
	lines := input.EachLineAsString(puzzleInput)
	regA := input.StringAsInt64(strings.Split(lines[0], ": ")[1])
	regB := input.StringAsInt64(strings.Split(lines[1], ": ")[1])
	regC := input.StringAsInt64(strings.Split(lines[2], ": ")[1])

	commands := make([]int64, 0)
	commandsString := strings.Split(lines[4], ": ")[1]
	for _, number := range strings.Split(commandsString, ",") {
		commands = append(commands, input.StringAsInt64(number))
	}

	return &Computer{
		RegA:     regA,
		RegB:     regB,
		RegC:     regC,
		Commands: commands,
		Output:   make([]int64, 0),
		Pointer:  0,
		Jump:     true,
	}
}

type Computer struct {
	RegA int64
	RegB int64
	RegC int64

	Pointer  int
	Commands []int64
	Output   []int64
	Jump     bool

	CommandsCompleted int64
}

func (c *Computer) Reset() {
	c.RegB = 0
	c.RegC = 0
	c.Pointer = 0
	c.Output = make([]int64, 0)
	c.Jump = true
	c.CommandsCompleted = 0
}

func (c *Computer) Run() {
	for {
		runAgain := c.RunStep()
		if !runAgain {
			break
		}
	}
}

func (c *Computer) RunStep() bool {
	// Pointer is outside the commands. Program is done!
	if c.Pointer > len(c.Commands)-2 {
		return false
	}

	currentCommand := c.Commands[c.Pointer]
	operand := c.Commands[c.Pointer+1]

	switch currentCommand {
	case 0:
		c.Adv(operand)
	case 1:
		c.Bxl(operand)
	case 2:
		c.Bst(operand)
	case 3:
		c.Jnz(operand)
	case 4:
		c.Bxc(operand)
	case 5:
		c.Out(operand)
	case 6:
		c.Bdv(operand)
	case 7:
		c.Cdv(operand)
	}

	c.CommandsCompleted++

	if c.Jump {
		c.Pointer += 2
	}
	c.Jump = true
	return true
}

func (c *Computer) getComboOperand(operand int64) int64 {
	if operand >= 0 && operand <= 3 {
		return operand
	}
	if operand == 4 {
		return c.RegA
	}
	if operand == 5 {
		return c.RegB
	}
	if operand == 6 {
		return c.RegC
	}

	log.Fatalf("Received invalid operand 7!")
	return 0
}

func (c *Computer) Adv(operand int64) {
	// OP: 0
	numerator := c.RegA
	comboOperand := c.getComboOperand(operand)
	denomimator := int64(math.Pow(2, float64(comboOperand)))
	result := numerator / denomimator
	c.RegA = result
}

func (c *Computer) Bxl(operand int64) {
	// OP: 1
	result := c.RegB ^ operand
	c.RegB = result
}

func (c *Computer) Bst(operand int64) {
	// OP: 2
	result := c.getComboOperand(operand) % 8
	c.RegB = result
}

func (c *Computer) Jnz(operand int64) {
	// OP: 3
	if c.RegA == 0 {
		return
	}

	c.Pointer = int(operand)
	c.Jump = false
}

func (c *Computer) Bxc(operand int64) {
	// OP: 4
	result := c.RegB ^ c.RegC
	c.RegB = result
}

func (c *Computer) Out(operand int64) {
	// OP: 5
	result := c.getComboOperand(operand) % 8
	c.Output = append(c.Output, result)
}

func (c *Computer) Bdv(operand int64) {
	// OP: 6
	numerator := c.RegA
	comboOperand := c.getComboOperand(operand)
	denomimator := int64(math.Pow(2, float64(comboOperand)))

	result := numerator / denomimator
	c.RegB = result
}

func (c *Computer) Cdv(operand int64) {
	// OP: 7
	numerator := c.RegA
	comboOperand := c.getComboOperand(operand)
	denomimator := int64(math.Pow(2, float64(comboOperand)))
	result := numerator / denomimator
	c.RegC = result
}

func (c *Computer) ConcatenateOutput() string {
	output := ""

	for _, value := range c.Output {
		output += fmt.Sprintf("%d,", value)
	}

	return strings.TrimSuffix(output, ",")
}

func (c *Computer) FindSelfReplicatingA(a, depth int64) int64 {
	target := slices.Clone(c.Commands)
	slices.Reverse(target)

	if depth == int64(len(c.Commands)) {
		return a
	}

	for i := int64(0); i < 8; i++ {
		c.Reset()
		c.RegA = a*8 + i
		c.Run()

		if len(c.Output) > 0 && c.Output[0] == target[depth] {
			result := c.FindSelfReplicatingA((a*8 + i), depth+1)
			if result > 0 {
				return result
			}
		}
	}
	return 0
}
