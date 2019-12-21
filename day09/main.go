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
	MP int = iota // ModePosition
	MI            // ModeIntermediate
	MR            // ModeRelative
)

type program struct {
	memory       []int
	input        *programInput
	output       int
	ip           int
	relativeBase int
}

func (p *program) Read(posDelta int, mode int) int {
	pos := p.ip + posDelta

	var retValue int
	switch mode {
	case MR:
		pos = p.relativeBase + posDelta
		fallthrough
	case MP:
		retValue = p.memory[p.memory[pos]]
	case MI:
		retValue = p.memory[pos]
	default:
		panic(fmt.Sprintf("unsupported mode %d", mode))
	}
	fmt.Println("read:", "pos", pos, "mode", mode, "retValue", retValue, "posDelta", posDelta)
	return retValue
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
	adjustRelativeBase
	halt instruction = 99
)

func (i instruction) size() int {
	switch i {
	case add, multiply, lessThan, equals:
		return 4
	case jumpIfTrue, jumpIfFalse:
		return 3
	case input, output, adjustRelativeBase:
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
	opcode, modes := decodeInstruction(p.Read(0, MI))
	if opcode == halt {
		return true
	}
	nextIp := p.ip + opcode.size()
	fmt.Println("opcode", opcode)
	switch opcode {
	case add:
		p.Write(p.Read(3, MI), p.Read(1, modes[0])+p.Read(2, modes[1]))
	case multiply:
		p.Write(p.Read(3, MI), p.Read(1, modes[0])*p.Read(2, modes[1]))
	case input:
		p.Write(p.Read(1, MI), p.input.read())
	case output:
		p.output = p.Read(1, modes[0])
		p.ip = nextIp
		fmt.Println("outputing", p.output)
	case jumpIfTrue:
		if p.Read(p.ip+1, modes[0]) != 0 {
			nextIp = p.Read(2, modes[1])
		}
	case jumpIfFalse:
		if p.Read(p.ip+1, modes[0]) == 0 {
			nextIp = p.Read(2, modes[1])
		}
	case lessThan:
		p1 := p.Read(1, modes[0])
		p2 := p.Read(2, modes[1])
		result := 0
		if p1 < p2 {
			result = 1
		}
		p.Write(p.Read(p.ip+3, MI), result)
	case equals:
		p1 := p.Read(1, modes[0])
		p2 := p.Read(2, modes[1])
		result := 0
		if p1 == p2 {
			result = 1
		}
		p.Write(p.Read(3, MI), result)
	case adjustRelativeBase:
		p.relativeBase = p.relativeBase + p.Read(1, MI)
		fmt.Println("relative base", p.relativeBase)
	default:
		panic(fmt.Sprintf("unsupported opcode %d", opcode))
	}
	p.ip = nextIp
	return p.run()
}

func main() {

	p := program{
		memory:       startingMemory(1000),
		ip:           0,
		relativeBase: 0,
	}

	halt := p.run()
	fmt.Println(halt, p.output)

}

func startingMemory(programSpace int) []int {
	m := []int{109, 1, 204, -1, 1001, 100, 1, 100, 1008, 100, 16, 101, 1006, 101, 0, 99}
	zeros := make([]int, programSpace)
	m = append(m, zeros...)
	return m
}
