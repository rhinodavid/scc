package main

import "fmt"

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
	g.nodes[h].rEdges = append(g.nodes[t].rEdges, t)
}

func (g *graph) addNode(label int) {
	if _, ok := g.nodes[label]; !ok {
		n := newNode()
		g.nodes[label] = n
	}
}

func main() {
	g := newGraph()
	g.addNode(1)
	g.addNode(2)
	g.addEdge(1, 2)
	fmt.Println(g.nodes[3])
}
