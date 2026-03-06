package main

import "fmt"

// directed graph

type vertex struct {
	key      int
	adjacent []*vertex
}

func NewVertex(key int) *vertex {
	return &vertex{
		key: key,
	}
}

// represents an adjacency list graph
type Graph struct {
	vertices []*vertex
}

func (g *Graph) AddEdge(from, to int) {
	fromVertex := g.getVertex(from)
	toVertex := g.getVertex(to)

	if fromVertex == nil || toVertex == nil {
		err := fmt.Errorf("Invalid edge (%v -> %v)", from, to)
		fmt.Println(err.Error())
	} else if contains(fromVertex.adjacent, to) {
		err := fmt.Errorf("Edge (%v -> %v) Already Exists! ", from, to)
		fmt.Println(err.Error())
	} else {
		fromVertex.adjacent = append(fromVertex.adjacent, toVertex)
	}

}

func (g *Graph) getVertex(key int) *vertex {
	for i, v := range g.vertices {
		if key == v.key {
			return g.vertices[i]
		}
	}
	return nil
}

func (g *Graph) AddVertex(key int) {
	if contains(g.vertices, key) {
		err := fmt.Errorf("Vertex %v not added because it is an exisiting key", key)
		fmt.Println(err.Error())
	} else {
		g.vertices = append(g.vertices, NewVertex(key))
	}
}

func contains(vertices []*vertex, key int) bool {
	for _, v := range vertices {
		if key == v.key {
			return true
		}
	}
	return false
}

func (g *Graph) Print() {
	for _, v := range g.vertices {
		fmt.Printf("\nVertex %v:", v.key)
		for _, v := range v.adjacent {
			fmt.Printf("%v", v.key)
		}
	}
	fmt.Println()
}

func main() {
	graph := &Graph{}

	for i := range 5 {
		graph.AddVertex(i)
	}

	graph.AddVertex(1)

	graph.AddEdge(0, 4)
	graph.AddEdge(0, 4)
	graph.AddEdge(11, 20)

	graph.Print()
}
