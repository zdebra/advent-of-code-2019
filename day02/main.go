package main

import "fmt"

var input = []int{1, 0, 0, 3, 1, 1, 2, 3, 1, 3, 4, 3, 1, 5, 0, 3, 2, 10, 1, 19, 1, 6, 19, 23, 1, 23, 13, 27, 2, 6, 27, 31, 1, 5, 31, 35, 2, 10, 35, 39, 1, 6, 39, 43, 1, 13, 43, 47, 2, 47, 6, 51, 1, 51, 5, 55, 1, 55, 6, 59, 2, 59, 10, 63, 1, 63, 6, 67, 2, 67, 10, 71, 1, 71, 9, 75, 2, 75, 10, 79, 1, 79, 5, 83, 2, 10, 83, 87, 1, 87, 6, 91, 2, 9, 91, 95, 1, 95, 5, 99, 1, 5, 99, 103, 1, 103, 10, 107, 1, 9, 107, 111, 1, 6, 111, 115, 1, 115, 5, 119, 1, 10, 119, 123, 2, 6, 123, 127, 2, 127, 6, 131, 1, 131, 2, 135, 1, 10, 135, 0, 99, 2, 0, 14, 0}

const (
	add      = 1
	multiply = 2
	halt     = 99

	nounIndex     = 1
	verbIndex     = 2
	desiredOutput = 19690720
)

func processIntcode(inp []int, i int) {
	if i > len(inp) {
		return
	}
	switch inp[i] {
	case add:
		inp[inp[i+3]] = inp[inp[i+1]] + inp[inp[i+2]]
	case multiply:
		inp[inp[i+3]] = inp[inp[i+1]] * inp[inp[i+2]]
	case halt:
		return
	default:
		panic(fmt.Sprintf("unsupported opcode %d", inp[i]))
	}
	processIntcode(inp, i+4)
}

func start(inp []int, noun, verb int) int {
	inp[nounIndex] = noun
	inp[verbIndex] = verb
	processIntcode(inp, 0)
	return inp[0]
}

func main() {
	for noun := 0; noun < 100; noun++ {
		for verb := 0; verb < 100; verb++ {
			inpCopy := make([]int, len(input))
			copy(inpCopy, input)
			pos0 := start(inpCopy, noun, verb)
			if pos0 == desiredOutput {
				fmt.Printf("noun = %d, verb = %d\n", noun, verb)
				fmt.Println(100*noun + verb)
				return
			}
		}
	}

	fmt.Printf("couldn't find noun and verb to match %d\n", desiredOutput)
}
