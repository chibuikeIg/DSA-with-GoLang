package main

import "fmt"

type Graph struct {
	numberOfNodes int
	adjacentList  map[string][]string
}

func NewGraph() *Graph {

	m := make(map[string][]string)

	return &Graph{adjacentList: m}
}

func (g *Graph) addNode(node string) {

	g.adjacentList[node] = nil

	g.numberOfNodes++

}

func (g *Graph) addEgde(node1, node2 string) {

	g.adjacentList[node1] = append(g.adjacentList[node1], node2)
	g.adjacentList[node2] = append(g.adjacentList[node2], node1)

}

func main() {

	graph := NewGraph()

	graph.addNode("0")
	graph.addNode("1")
	graph.addNode("2")
	graph.addNode("3")
	graph.addNode("4")
	graph.addNode("5")
	graph.addNode("6")

	graph.addEgde("3", "1")
	graph.addEgde("3", "4")
	graph.addEgde("4", "2")
	graph.addEgde("4", "5")
	graph.addEgde("1", "2")
	graph.addEgde("1", "0")
	graph.addEgde("0", "2")
	graph.addEgde("6", "5")

	fmt.Println(graph.adjacentList)

}
