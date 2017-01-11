package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type node struct {
	finishingNumber, startingNumber int
	visited                         bool
	edges, rEdges                   []int
}

type graph struct {
	nodes map[int]*node
}

func newGraph() *graph {
	var g graph
	g.nodes = make(map[int]*node)
	return &g
}

func newNode() *node {
	var n node
	n.startingNumber = -1
	n.finishingNumber = -1
	return &n
}

func (g *graph) addEdge(t, h int) {
	if _, ok := g.nodes[t]; !ok {
		panic("No node for edge tail")
	}
	if _, ok := g.nodes[h]; !ok {
		panic("No node for edge head")
	}
	g.nodes[t].edges = append(g.nodes[t].edges, h)
	g.nodes[h].rEdges = append(g.nodes[h].rEdges, t)
}

func (g *graph) addNode(label int) bool {
	if _, ok := g.nodes[label]; !ok {
		n := newNode()
		g.nodes[label] = n
		return true
	}
	return false
}

func (g *graph) showGraph() {
	for k, v := range g.nodes {
		fmt.Printf("Node %d:\nEdges: %v\nBackwards Edges: %v\n\n", k, v.edges, v.rEdges)
	}
}

func main() {
	g := newGraph()

	flag.Parse()
	if len(flag.Args()) < 1 {
		panic("Enter the name of the file with the graph edges list")
	}
	f, err := os.Open(flag.Args()[0])
	if err != nil {
		panic(err)
	}
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		line := scanner.Text()
		iS := strings.Fields(line)
		if len(iS) != 2 {
			panic(fmt.Sprintf("Bad line in graph file: %v", iS))
		}
		t, err1 := strconv.Atoi(iS[0])
		if err1 != nil {
			panic(err1)
		}
		h, err2 := strconv.Atoi(iS[1])
		if err2 != nil {
			panic(err2)
		}
		g.addNode(t)
		g.addNode(h)
		g.addEdge(t, h)
	}
	g.showGraph()
}
