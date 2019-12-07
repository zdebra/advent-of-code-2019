package main

import "fmt"

type programInput struct {
	m        [2]int
	accessed bool
}

func (pi *programInput) read() int {
	if !pi.accessed {
		pi.accessed = true
		return pi.m[0]
	}
	return pi.m[1]
}

const (
	MP int = iota
	MI
)

type program struct {
	memory []int
	input  *programInput
	output int
	ip     int
}

func (p *program) Read(pos int, mode int) int {
	switch mode {
	case MP:
		return p.memory[p.memory[pos]]
	case MI:
		return p.memory[pos]
	default:
		panic(fmt.Sprintf("unsupported mode %d", mode))
	}
}

func (p *program) Write(pos, value int) {
	p.memory[pos] = value
}

func (p *program) Len() int {
	return len(p.memory)
}

type instruction int

const (
	add instruction = iota + 1
	multiply
	input
	output
	jumpIfTrue
	jumpIfFalse
	lessThan
	equals
	halt instruction = 99
)

func (i instruction) size() int {
	switch i {
	case add, multiply, lessThan, equals:
		return 4
	case jumpIfTrue, jumpIfFalse:
		return 3
	case input, output:
		return 2
	case halt:
		return 1
	default:
		panic(fmt.Sprintf("unsupported instruction %d", i))
	}
}

func decodeInstruction(code int) (instruction, [3]int) {
	opcode := code % 100
	code = code / 100
	modes := [3]int{}
	for i := range modes {
		modes[i] = code % 10
		code = code / 10
	}
	return instruction(opcode), modes
}

func (p *program) run() bool {
	if p.ip > p.Len() {
		return false
	}
	opcode, modes := decodeInstruction(p.Read(p.ip, MI))
	if opcode == halt {
		return true
	}
	nextIp := p.ip + opcode.size()
	switch opcode {
	case add:
		p.Write(p.Read(p.ip+3, MI), p.Read(p.ip+1, modes[0])+p.Read(p.ip+2, modes[1]))
	case multiply:
		p.Write(p.Read(p.ip+3, MI), p.Read(p.ip+1, modes[0])*p.Read(p.ip+2, modes[1]))
	case input:
		p.Write(p.Read(p.ip+1, MI), p.input.read())
	case output:
		p.output = p.Read(p.ip+1, MP)
		p.ip = nextIp
		return false
	case jumpIfTrue:
		if p.Read(p.ip+1, modes[0]) != 0 {
			nextIp = p.Read(p.ip+2, modes[1])
		}
	case jumpIfFalse:
		if p.Read(p.ip+1, modes[0]) == 0 {
			nextIp = p.Read(p.ip+2, modes[1])
		}
	case lessThan:
		p1 := p.Read(p.ip+1, modes[0])
		p2 := p.Read(p.ip+2, modes[1])
		result := 0
		if p1 < p2 {
			result = 1
		}
		p.Write(p.Read(p.ip+3, MI), result)
	case equals:
		p1 := p.Read(p.ip+1, modes[0])
		p2 := p.Read(p.ip+2, modes[1])
		result := 0
		if p1 == p2 {
			result = 1
		}
		p.Write(p.Read(p.ip+3, MI), result)
	default:
		panic(fmt.Sprintf("unsupported opcode %d", opcode))
	}
	p.ip = nextIp
	return p.run()
}

func programChain(initInput int, phaseSetting []int) int {
	previousPhaseOutput := initInput
	for _, phase := range phaseSetting {
		p := program{
			memory: startingMemory(),
			input: &programInput{
				m: [2]int{phase, previousPhaseOutput},
			},
		}
		p.run()
		previousPhaseOutput = p.output
	}
	return previousPhaseOutput
}

func programLoop(initInput int, phaseSetting []int) int {
	programs := make([]*program, 5)
	// init programs
	for i, phase := range phaseSetting {
		programs[i] = &program{
			memory: startingMemory(),
			input: &programInput{
				m: [2]int{phase},
			},
		}
	}

	// loop until halt
	previousPhaseOutput := initInput
	halted := false
	for !halted {
		for _, prg := range programs {
			prg.input.m[1] = previousPhaseOutput
			halted = prg.run()
			previousPhaseOutput = prg.output
		}
	}
	return programs[4].output
}

func main() {
	maxInput := splitInt(56789)
	max := programLoop(0, maxInput)
	for i := 0; i <= 99999; i++ {
		curInp := splitInt(i)
		if !matchCriteria(curInp) {
			continue
		}
		curOutput := programLoop(0, curInp)
		if curOutput > max {
			max = curOutput
			maxInput = curInp
		}
	}
	fmt.Println(max, maxInput)
}

func matchCriteria(inp []int) bool {
	for i := 0; i < len(inp); i++ {
		if inp[i] < 5 {
			return false
		}
		for j := i + 1; j < len(inp); j++ {
			if inp[j] < 5 {
				return false
			}
			if inp[i] == inp[j] {
				return false
			}
		}
	}
	return true
}

func splitInt(inp int) []int {
	res := make([]int, 5)
	for i := 4; i >= 0; i-- {
		res[i] = inp % 10
		inp /= 10
	}
	return res
}

func startingMemory() []int {
	return []int{3, 8, 1001, 8, 10, 8, 105, 1, 0, 0, 21, 38, 55, 64, 81, 106, 187, 268, 349, 430, 99999, 3, 9, 101, 2, 9, 9, 1002, 9, 2, 9, 101, 5, 9, 9, 4, 9, 99, 3, 9, 102, 2, 9, 9, 101, 3, 9, 9, 1002, 9, 4, 9, 4, 9, 99, 3, 9, 102, 2, 9, 9, 4, 9, 99, 3, 9, 1002, 9, 5, 9, 1001, 9, 4, 9, 102, 4, 9, 9, 4, 9, 99, 3, 9, 102, 2, 9, 9, 1001, 9, 5, 9, 102, 3, 9, 9, 1001, 9, 4, 9, 102, 5, 9, 9, 4, 9, 99, 3, 9, 1002, 9, 2, 9, 4, 9, 3, 9, 101, 2, 9, 9, 4, 9, 3, 9, 1002, 9, 2, 9, 4, 9, 3, 9, 1001, 9, 2, 9, 4, 9, 3, 9, 1001, 9, 2, 9, 4, 9, 3, 9, 101, 1, 9, 9, 4, 9, 3, 9, 1001, 9, 1, 9, 4, 9, 3, 9, 1001, 9, 2, 9, 4, 9, 3, 9, 101, 1, 9, 9, 4, 9, 3, 9, 1001, 9, 1, 9, 4, 9, 99, 3, 9, 1002, 9, 2, 9, 4, 9, 3, 9, 101, 2, 9, 9, 4, 9, 3, 9, 1001, 9, 1, 9, 4, 9, 3, 9, 101, 1, 9, 9, 4, 9, 3, 9, 101, 2, 9, 9, 4, 9, 3, 9, 101, 2, 9, 9, 4, 9, 3, 9, 1001, 9, 1, 9, 4, 9, 3, 9, 101, 1, 9, 9, 4, 9, 3, 9, 102, 2, 9, 9, 4, 9, 3, 9, 101, 2, 9, 9, 4, 9, 99, 3, 9, 1002, 9, 2, 9, 4, 9, 3, 9, 101, 2, 9, 9, 4, 9, 3, 9, 102, 2, 9, 9, 4, 9, 3, 9, 101, 2, 9, 9, 4, 9, 3, 9, 1001, 9, 2, 9, 4, 9, 3, 9, 1002, 9, 2, 9, 4, 9, 3, 9, 1002, 9, 2, 9, 4, 9, 3, 9, 101, 2, 9, 9, 4, 9, 3, 9, 1001, 9, 2, 9, 4, 9, 3, 9, 101, 1, 9, 9, 4, 9, 99, 3, 9, 102, 2, 9, 9, 4, 9, 3, 9, 1001, 9, 2, 9, 4, 9, 3, 9, 1002, 9, 2, 9, 4, 9, 3, 9, 102, 2, 9, 9, 4, 9, 3, 9, 102, 2, 9, 9, 4, 9, 3, 9, 101, 2, 9, 9, 4, 9, 3, 9, 101, 1, 9, 9, 4, 9, 3, 9, 101, 1, 9, 9, 4, 9, 3, 9, 1001, 9, 1, 9, 4, 9, 3, 9, 102, 2, 9, 9, 4, 9, 99, 3, 9, 101, 1, 9, 9, 4, 9, 3, 9, 1002, 9, 2, 9, 4, 9, 3, 9, 102, 2, 9, 9, 4, 9, 3, 9, 1002, 9, 2, 9, 4, 9, 3, 9, 101, 1, 9, 9, 4, 9, 3, 9, 102, 2, 9, 9, 4, 9, 3, 9, 1002, 9, 2, 9, 4, 9, 3, 9, 1002, 9, 2, 9, 4, 9, 3, 9, 101, 1, 9, 9, 4, 9, 3, 9, 102, 2, 9, 9, 4, 9, 99}
}
