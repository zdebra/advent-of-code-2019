package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type node struct {
	parent   *node
	children []*node
	id       string
	visited  bool
}

func (n *node) notVisitedNeighbors() []*node {
	nv := make([]*node, 0, len(n.children)+1)
	if n.parent != nil && !n.parent.visited {
		nv = append(nv, n.parent)
	}
	for _, ch := range n.children {
		if !ch.visited {
			nv = append(nv, ch)
		}
	}
	return nv
}

func newNode(id string) *node {
	return &node{
		id:       id,
		children: []*node{},
	}
}

func main() {
	nodes := map[string]*node{}

	f, err := os.Open("day06/input")
	guard(err)

	sc := bufio.NewScanner(f)
	for sc.Scan() {
		splits := strings.Split(sc.Text(), ")")
		if len(splits) != 2 {
			panic("invalid input")
		}
		a, b := splits[0], splits[1]
		if _, found := nodes[a]; !found {
			nodes[a] = newNode(a)
		}
		if _, found := nodes[b]; !found {
			nodes[b] = newNode(b)
		}
		nodes[b].parent = nodes[a]
		nodes[a].children = append(nodes[a].children, nodes[b])
	}

	you, san := nodes["YOU"], nodes["SAN"]
	start := you.parent
	finish := san.parent

	findNode(finish.id, start.notVisitedNeighbors(), 0)
}

func findNode(finishID string, in []*node, cnt int) {
	cnt++
	for _, n := range in {
		if n.id == finishID {
			fmt.Printf("found finish in %d steps\n", cnt)
		}
		n.visited = true
		findNode(finishID, n.notVisitedNeighbors(), cnt)
	}
}

func step(curNode *node, cnt int) int {
	total := cnt
	for _, ch := range curNode.children {
		total += step(ch, cnt+1)
	}
	return total
}

func guard(err error) {
	if err != nil {
		panic(err)
	}
}
