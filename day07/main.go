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
	input  programInput
	output int
}

func (p program) Read(pos int, mode int) int {
	switch mode {
	case MP:
		return p.memory[p.memory[pos]]
	case MI:
		return p.memory[pos]
	default:
		panic(fmt.Sprintf("unsupported mode %d", mode))
	}
}

func (p program) Write(pos, value int) {
	p.memory[pos] = value
}

func (p program) Len() int {
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

func (p *program) run(ip int) {
	if ip > p.Len() {
		return
	}
	opcode, modes := decodeInstruction(p.Read(ip, MI))
	if opcode == halt {
		return
	}
	nextIp := ip + opcode.size()
	switch opcode {
	case add:
		p.Write(p.Read(ip+3, MI), p.Read(ip+1, modes[0])+p.Read(ip+2, modes[1]))
	case multiply:
		p.Write(p.Read(ip+3, MI), p.Read(ip+1, modes[0])*p.Read(ip+2, modes[1]))
	case input:
		p.Write(p.Read(ip+1, MI), p.input.read())
	case output:
		p.output = p.Read(ip+1, MP)
		//fmt.Println("outputting!", p.output)
	case jumpIfTrue:
		if p.Read(ip+1, modes[0]) != 0 {
			nextIp = p.Read(ip+2, modes[1])
		}
	case jumpIfFalse:
		if p.Read(ip+1, modes[0]) == 0 {
			nextIp = p.Read(ip+2, modes[1])
		}
	case lessThan:
		p1 := p.Read(ip+1, modes[0])
		p2 := p.Read(ip+2, modes[1])
		result := 0
		if p1 < p2 {
			result = 1
		}
		p.Write(p.Read(ip+3, MI), result)
	case equals:
		p1 := p.Read(ip+1, modes[0])
		p2 := p.Read(ip+2, modes[1])
		result := 0
		if p1 == p2 {
			result = 1
		}
		p.Write(p.Read(ip+3, MI), result)
	case halt:
		return
	default:
		panic(fmt.Sprintf("unsupported opcode %d", opcode))
	}
	p.run(nextIp)
}

func programChain(initInput int, phaseSetting []int) int {
	previousPhaseOutput := initInput
	for _, phase := range phaseSetting {
		p := program{
			memory: startingMemory(),
			input: programInput{
				m: [2]int{phase, previousPhaseOutput},
			},
		}
		p.run(0)
		previousPhaseOutput = p.output
	}
	return previousPhaseOutput
}

func main() {
	maxInput := splitInt(0)
	max := programChain(0, maxInput)
	for i := 1; i < 44445; i++ {
		curInp := splitInt(i)
		if !matchCriteria(curInp) {
			continue
		}
		curOutput := programChain(0, curInp)
		if curOutput > max {
			max = curOutput
			maxInput = curInp
		}
	}
	fmt.Println(max, maxInput)
}

func matchCriteria(inp []int) bool {
	for i := 0; i < len(inp); i++ {
		for j := i + 1; j < len(inp); j++ {
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
