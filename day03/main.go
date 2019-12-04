package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

const (
	cmdRight uint8 = 'R'
	cmdLeft  uint8 = 'L'
	cmdUp    uint8 = 'U'
	cmdDown  uint8 = 'D'
)

type point struct {
	X, Y int
}

type panel map[int]map[int]int

func safeSet(panel panel, pos point, steps int) {
	if _, found := panel[pos.X]; !found {
		panel[pos.X] = map[int]int{}
	}
	panel[pos.X][pos.Y] = steps
}

func mark(panel panel, lastPos point, cmd string, steps int) (point, int) {
	delta, err := strconv.Atoi(cmd[1:])
	if err != nil {
		panic(err)
	}
	x, y := lastPos.X, lastPos.Y
	switch cmd[0] {
	case cmdRight:
		for i := 0; i <= delta; i++ {
			x, y = lastPos.X+i, lastPos.Y
			safeSet(panel, point{x, y}, steps)
			steps++
		}
	case cmdLeft:
		for i := 0; i <= delta; i++ {
			x, y = lastPos.X-i, lastPos.Y
			safeSet(panel, point{x, y}, steps)
			steps++
		}
	case cmdUp:
		for i := 0; i <= delta; i++ {
			x, y = lastPos.X, lastPos.Y+i
			safeSet(panel, point{x, y}, steps)
			steps++
		}
	case cmdDown:
		for i := 0; i <= delta; i++ {
			x, y = lastPos.X, lastPos.Y-i
			safeSet(panel, point{x, y}, steps)
			steps++
		}
	default:
		panic(fmt.Sprintf("unsupported cmd %s", string(cmd[0])))
	}
	steps--
	return point{x, y}, steps
}

func intersect(markedPanel panel, lastPos point, cmd string, steps int) ([]int, point, int) {
	intersectSteps := []int{}
	u := func(x, y, steps int) {
		if x == 0 && y == 0 {
			return
		}
		if _, foundX := markedPanel[x]; foundX {
			if _, foundY := markedPanel[x][y]; foundY {
				wireASteps, wireBSteps := markedPanel[x][y], steps
				intersectSteps = append(intersectSteps, wireASteps+wireBSteps)
			}
		}
	}

	delta, err := strconv.Atoi(cmd[1:])
	if err != nil {
		panic(err)
	}
	x, y := lastPos.X, lastPos.Y
	switch cmd[0] {
	case cmdRight:
		for i := 0; i <= delta; i++ {
			x, y = lastPos.X+i, lastPos.Y
			u(x, y, steps)
			steps++
		}
	case cmdLeft:
		for i := 0; i <= delta; i++ {
			x, y = lastPos.X-i, lastPos.Y
			u(x, y, steps)
			steps++
		}
	case cmdUp:
		for i := 0; i <= delta; i++ {
			x, y = lastPos.X, lastPos.Y+i
			u(x, y, steps)
			steps++
		}
	case cmdDown:
		for i := 0; i <= delta; i++ {
			x, y = lastPos.X, lastPos.Y-i
			u(x, y, steps)
			steps++
		}
	default:
		panic(fmt.Sprintf("unsupported cmd %s", string(cmd[0])))
	}
	steps--
	return intersectSteps, point{x, y}, steps
}

func distance(p1, p2 point) float64 {
	return math.Abs(float64(p1.X)-float64(p2.X)) + math.Abs(float64(p1.Y)-float64(p2.Y))
}

func printPanel(panel panel) {
	s := 30
	for i := -s; i < s; i++ {
		for j := -s; j < s; j++ {
			if i == 0 && j == 0 {
				fmt.Print("o")
				continue
			}
			var foundX, foundY bool
			if _, foundX = panel[i]; foundX {
				if _, foundY = panel[i][j]; foundY {
					fmt.Print("*")
				}
			}
			if !foundX || !foundY {
				fmt.Print(".")
			}
		}
		fmt.Print("\n")
	}
	fmt.Print("\n")
}

func main() {

	f, err := os.Open("day03/input")
	if err != nil {
		panic(err)
	}

	sc := bufio.NewScanner(f)
	sc.Scan()
	wireACommands := strings.Split(sc.Text(), ",")
	sc.Scan()
	wireBCommands := strings.Split(sc.Text(), ",")

	wireA := panel{}
	lastPos := point{0, 0}
	safeSet(wireA, lastPos, 0)
	steps := 0
	for _, cmd := range wireACommands {
		lastPos, steps = mark(wireA, lastPos, cmd, steps)
	}

	lastPos = point{0, 0}
	intersections := []int{}
	steps = 0
	for _, cmd := range wireBCommands {
		var intersectionsPart []int
		intersectionsPart, lastPos, steps = intersect(wireA, lastPos, cmd, steps)
		intersections = append(intersections, intersectionsPart...)
	}

	leastSteps := intersections[0]
	for i := 1; i < len(intersections); i++ {
		if intersections[i] < leastSteps {
			leastSteps = intersections[i]
		}
	}

	fmt.Println(leastSteps)

}
