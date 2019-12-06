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

	curNode, found := nodes["COM"]
	if !found {
		panic("COM not found")
	}

	total := step(curNode, 0)
	fmt.Println(total)
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
