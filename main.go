package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

type node struct {
	finishingNumber, sccID int
	visited                bool
	edges, rEdges          []int
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
	n.sccID = -1
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

func (g *graph) resetVisited() {
	for _, n := range g.nodes {
		n.visited = false
	}
}

func (g *graph) showGraph() {
	for k, v := range g.nodes {
		fmt.Printf("Node %d (SCC ID#: %d):\nEdges: %v\nBackwards Edges: %v\n\n", k, v.sccID, v.edges, v.rEdges)
	}
}

func (g *graph) createFinishingOrder() []*node {
	g.resetVisited()
	t := make([]*node, 0, len(g.nodes))
	for _, v := range g.nodes {
		if v.visited == false {
			dfsAssignFinishingNumber(v, g, &t)
		}
	}
	return t
}

func dfsAssignFinishingNumber(n *node, g *graph, t *[]*node) {
	n.visited = true
	for _, neighborLabel := range n.rEdges {
		if g.nodes[neighborLabel].visited == false {
			dfsAssignFinishingNumber(g.nodes[neighborLabel], g, t)
		}
	}
	(*t) = append(*t, n)
}

func dfsMarkScc(n *node, g *graph, s int) {
	n.visited = true
	n.sccID = s
	for _, neighborLabel := range n.edges {
		if g.nodes[neighborLabel].visited == false {
			dfsMarkScc(g.nodes[neighborLabel], g, s)
		}
	}
}

func (g *graph) generateScc() {
	p := 0
	fo := g.createFinishingOrder()
	g.resetVisited()
	for i := len(fo) - 1; i >= 0; i-- {
		n := fo[i]
		if n.visited == false {
			s := p
			p++
			dfsMarkScc(n, g, s)
		}
	}
}

func (g *graph) getTopFiveSccs() []int {
	a := make(map[int]int)
	for _, n := range g.nodes {
		a[n.sccID]++
	}
	r := make([]int, 5)
	for _, sccPopulation := range a {
		if sccPopulation > r[0] {
			r[0] = sccPopulation
		}
		// sort descending
		sort.Ints(r)
	}
	return r
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
	g.generateScc()
	fmt.Println(g.getTopFiveSccs())
}
