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

func safeSet(panel map[int]map[int]struct{}, pos point) {
	if _, found := panel[pos.X]; !found {
		panel[pos.X] = map[int]struct{}{}
	}
	panel[pos.X][pos.Y] = struct{}{}
}

func mark(panel map[int]map[int]struct{}, lastPos point, cmd string) point {
	delta, err := strconv.Atoi(cmd[1:])
	if err != nil {
		panic(err)
	}
	x, y := lastPos.X, lastPos.Y
	switch cmd[0] {
	case cmdRight:
		for i := 0; i <= delta; i++ {
			x, y = lastPos.X+i, lastPos.Y
			safeSet(panel, point{x, y})
		}
	case cmdLeft:
		for i := 0; i <= delta; i++ {
			x, y = lastPos.X-i, lastPos.Y
			safeSet(panel, point{x, y})
		}
	case cmdUp:
		for i := 0; i <= delta; i++ {
			x, y = lastPos.X, lastPos.Y+i
			safeSet(panel, point{x, y})
		}
	case cmdDown:
		for i := 0; i <= delta; i++ {
			x, y = lastPos.X, lastPos.Y-i
			safeSet(panel, point{x, y})
		}
	default:
		panic(fmt.Sprintf("unsupported cmd %s", string(cmd[0])))
	}
	return point{x, y}
}

func intersect(markedPanel map[int]map[int]struct{}, lastPos point, cmd string) ([]point, point) {
	intersections := []point{}
	u := func(x, y int) {
		if x == 0 && y == 0 {
			return
		}
		if _, foundX := markedPanel[x]; foundX {
			if _, foundY := markedPanel[x][y]; foundY {
				intersections = append(intersections, point{x, y})
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
			u(x, y)
		}
	case cmdLeft:
		for i := 0; i <= delta; i++ {
			x, y = lastPos.X-i, lastPos.Y
			u(x, y)
		}
	case cmdUp:
		for i := 0; i <= delta; i++ {
			x, y = lastPos.X, lastPos.Y+i
			u(x, y)
		}
	case cmdDown:
		for i := 0; i <= delta; i++ {
			x, y = lastPos.X, lastPos.Y-i
			u(x, y)
		}
	default:
		panic(fmt.Sprintf("unsupported cmd %s", string(cmd[0])))
	}
	return intersections, point{x, y}
}

func distance(p1, p2 point) float64 {
	return math.Abs(float64(p1.X)-float64(p2.X)) + math.Abs(float64(p1.Y)-float64(p2.Y))
}

func printPanel(panel map[int]map[int]struct{}) {
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

	wireA := map[int]map[int]struct{}{}
	lastPos := point{0, 0}
	safeSet(wireA, lastPos)
	for _, cmd := range wireACommands {
		lastPos = mark(wireA, lastPos, cmd)
	}

	//printPanel(wireA)

	lastPos = point{0, 0}
	intersections := []point{}
	for _, cmd := range wireBCommands {
		var intersectionsPart []point
		intersectionsPart, lastPos = intersect(wireA, lastPos, cmd)
		intersections = append(intersections, intersectionsPart...)
	}

	fmt.Println(intersections)

	start := point{0, 0}
	closest := intersections[0]
	closestDistance := distance(start, closest)
	for i := 1; i < len(intersections); i++ {
		d := distance(start, intersections[i])
		if d < closestDistance {
			closest = intersections[i]
			closestDistance = d
		}
	}

	fmt.Printf("closest intersection is (%d,%d) with len from start %f\n", closest.X, closest.Y, closestDistance)

}
