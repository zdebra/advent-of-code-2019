package main

import "fmt"

const (
	MP int = iota
	MI
)

type program struct {
	memory []int
	input  int
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
		p.Write(p.Read(ip+1, MI), p.input)
	case output:
		p.output = p.Read(ip+1, MP)
		fmt.Println("outputting!", p.output)
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

func main() {
	p := program{
		memory: startingMemory,
		input:  5,
	}

	p.run(0)
	fmt.Println(p.output)
}

var startingMemory = []int{3, 225, 1, 225, 6, 6, 1100, 1, 238, 225, 104, 0, 1001, 152, 55, 224, 1001, 224, -68, 224, 4, 224, 1002, 223, 8, 223, 1001, 224, 4, 224, 1, 224, 223, 223, 1101, 62, 41, 225, 1101, 83, 71, 225, 102, 59, 147, 224, 101, -944, 224, 224, 4, 224, 1002, 223, 8, 223, 101, 3, 224, 224, 1, 224, 223, 223, 2, 40, 139, 224, 1001, 224, -3905, 224, 4, 224, 1002, 223, 8, 223, 101, 7, 224, 224, 1, 223, 224, 223, 1101, 6, 94, 224, 101, -100, 224, 224, 4, 224, 1002, 223, 8, 223, 101, 6, 224, 224, 1, 224, 223, 223, 1102, 75, 30, 225, 1102, 70, 44, 224, 101, -3080, 224, 224, 4, 224, 1002, 223, 8, 223, 1001, 224, 4, 224, 1, 223, 224, 223, 1101, 55, 20, 225, 1102, 55, 16, 225, 1102, 13, 94, 225, 1102, 16, 55, 225, 1102, 13, 13, 225, 1, 109, 143, 224, 101, -88, 224, 224, 4, 224, 1002, 223, 8, 223, 1001, 224, 2, 224, 1, 223, 224, 223, 1002, 136, 57, 224, 101, -1140, 224, 224, 4, 224, 1002, 223, 8, 223, 101, 6, 224, 224, 1, 223, 224, 223, 101, 76, 35, 224, 1001, 224, -138, 224, 4, 224, 1002, 223, 8, 223, 101, 5, 224, 224, 1, 223, 224, 223, 4, 223, 99, 0, 0, 0, 677, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1105, 0, 99999, 1105, 227, 247, 1105, 1, 99999, 1005, 227, 99999, 1005, 0, 256, 1105, 1, 99999, 1106, 227, 99999, 1106, 0, 265, 1105, 1, 99999, 1006, 0, 99999, 1006, 227, 274, 1105, 1, 99999, 1105, 1, 280, 1105, 1, 99999, 1, 225, 225, 225, 1101, 294, 0, 0, 105, 1, 0, 1105, 1, 99999, 1106, 0, 300, 1105, 1, 99999, 1, 225, 225, 225, 1101, 314, 0, 0, 106, 0, 0, 1105, 1, 99999, 1008, 677, 677, 224, 1002, 223, 2, 223, 1006, 224, 329, 1001, 223, 1, 223, 8, 677, 226, 224, 102, 2, 223, 223, 1006, 224, 344, 101, 1, 223, 223, 1107, 226, 226, 224, 1002, 223, 2, 223, 1006, 224, 359, 1001, 223, 1, 223, 1108, 677, 226, 224, 102, 2, 223, 223, 1005, 224, 374, 1001, 223, 1, 223, 1007, 226, 226, 224, 102, 2, 223, 223, 1006, 224, 389, 1001, 223, 1, 223, 108, 677, 677, 224, 1002, 223, 2, 223, 1005, 224, 404, 1001, 223, 1, 223, 1007, 677, 677, 224, 102, 2, 223, 223, 1005, 224, 419, 1001, 223, 1, 223, 8, 226, 677, 224, 102, 2, 223, 223, 1005, 224, 434, 101, 1, 223, 223, 1008, 677, 226, 224, 102, 2, 223, 223, 1006, 224, 449, 1001, 223, 1, 223, 7, 677, 677, 224, 102, 2, 223, 223, 1006, 224, 464, 1001, 223, 1, 223, 8, 226, 226, 224, 1002, 223, 2, 223, 1005, 224, 479, 1001, 223, 1, 223, 7, 226, 677, 224, 102, 2, 223, 223, 1006, 224, 494, 1001, 223, 1, 223, 7, 677, 226, 224, 1002, 223, 2, 223, 1005, 224, 509, 101, 1, 223, 223, 107, 677, 677, 224, 102, 2, 223, 223, 1006, 224, 524, 101, 1, 223, 223, 1007, 677, 226, 224, 102, 2, 223, 223, 1006, 224, 539, 101, 1, 223, 223, 107, 226, 226, 224, 1002, 223, 2, 223, 1006, 224, 554, 101, 1, 223, 223, 1008, 226, 226, 224, 102, 2, 223, 223, 1006, 224, 569, 1001, 223, 1, 223, 1107, 677, 226, 224, 1002, 223, 2, 223, 1005, 224, 584, 101, 1, 223, 223, 1107, 226, 677, 224, 102, 2, 223, 223, 1005, 224, 599, 101, 1, 223, 223, 1108, 226, 677, 224, 102, 2, 223, 223, 1005, 224, 614, 101, 1, 223, 223, 108, 677, 226, 224, 102, 2, 223, 223, 1005, 224, 629, 101, 1, 223, 223, 107, 226, 677, 224, 102, 2, 223, 223, 1006, 224, 644, 1001, 223, 1, 223, 1108, 226, 226, 224, 1002, 223, 2, 223, 1006, 224, 659, 101, 1, 223, 223, 108, 226, 226, 224, 102, 2, 223, 223, 1005, 224, 674, 101, 1, 223, 223, 4, 223, 99, 226}
