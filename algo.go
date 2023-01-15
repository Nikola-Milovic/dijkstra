package main

import (
	"fmt"
	"math"
)

// ItemGraph struct predstavlja graf koji se koristi u algoritmu za pronalaženje najkraćeg puta.
type ItemGraph struct {
	Nodes []*Node
	Edges map[Node][]*Edge
}

// Funkcija AddNode dodaje novi čvor u graf.
func (g *ItemGraph) AddNode(n *Node) {
	g.Nodes = append(g.Nodes, n)
}

// Funkcija AddEdge dodaje novu granu između dva čvora u graf sa odgovarajućom težinom.
func (g *ItemGraph) AddEdge(n1, n2 *Node, weight int) {
	if g.Edges == nil {
		g.Edges = make(map[Node][]*Edge)
	}
	ed1 := Edge{
		Node:   n2,
		Weight: weight,
	}

	ed2 := Edge{
		Node:   n1,
		Weight: weight,
	}
	g.Edges[*n1] = append(g.Edges[*n1], &ed1)
	g.Edges[*n2] = append(g.Edges[*n2], &ed2)
}

// Dijkstra algoritam
// g je instanca ItemGraph koji se koristi u algoritmu.
func getShortestPath(startNode *Node, endNode *Node, g *ItemGraph) ([]string, int) {
	// visited mapa se koristi da se proveri da li je neki čvor već bio posećen.
	visited := make(map[string]bool)
	// dist mapa cuva udaljenost svakog cvora od početnog cvora.
	dist := make(map[string]int)
	// prev mapa cuva prethodni cvor za svaki cvor.
	prev := make(map[string]string)
	// q je prioritetni red koji se koristi u algoritmu.
	q := NodeQueue{}
	pq := q.NewQueue()
	// start predstavlja početni čvor sa inicijalnom udaljenošću 0.
	start := Vertex{
		Node:     startNode,
		Distance: 0,
	}
	for _, nval := range g.Nodes {
		dist[nval.Value] = math.MaxInt64
	}
	dist[startNode.Value] = start.Distance
	pq.Enqueue(start)
	for !pq.IsEmpty() {
		v := pq.Dequeue()
		if visited[v.Node.Value] {
			continue
		}
		visited[v.Node.Value] = true

		// neighbours predstavlja listu susednih cvorova za trenutni cvor.
		neighbours := g.Edges[*v.Node]

		for _, val := range neighbours {
			if !visited[val.Node.Value] {
				if dist[v.Node.Value]+val.Weight < dist[val.Node.Value] {
					// store predstavlja trenutni cvor sa izracunatom udaljenošću.
					store := Vertex{
						Node:     val.Node,
						Distance: dist[v.Node.Value] + val.Weight,
					}
					dist[val.Node.Value] = dist[v.Node.Value] + val.Weight
					prev[val.Node.Value] = v.Node.Value
					pq.Enqueue(store)
				}
			}
		}
	}
	fmt.Println(dist)
	fmt.Println(prev)
	pathval := prev[endNode.Value]

	// finalArr predstavlja krajnji niz koji sadrzi najkraci put.
	var finalArr []string
	finalArr = append(finalArr, endNode.Value)
	for pathval != startNode.Value {
		finalArr = append(finalArr, pathval)
		pathval = prev[pathval]
	}
	finalArr = append(finalArr, pathval)
	fmt.Println(finalArr)
	for i, j := 0, len(finalArr)-1; i < j; i, j = i+1, j-1 {
		finalArr[i], finalArr[j] = finalArr[j], finalArr[i]
	}
	return finalArr, dist[endNode.Value]

}

// Funkcija CreateGraph kreira graf na osnovu ulaznih podataka.
func CreateGraph(data InputGraph) *ItemGraph {
	var g ItemGraph
	nodes := make(map[string]*Node)
	for _, v := range data.Graph {
		if _, found := nodes[v.Source]; !found {
			nA := Node{v.Source}
			nodes[v.Source] = &nA
			g.AddNode(&nA)
		}
		if _, found := nodes[v.Destination]; !found {
			nA := Node{v.Destination}
			nodes[v.Destination] = &nA
			g.AddNode(&nA)
		}
		g.AddEdge(nodes[v.Source], nodes[v.Destination], v.Weight)
	}
	return &g
}

// Pomocna funkcija koja zapocinje algoritam i vraca rezultat
func GetShortestPath(from, to string, g *ItemGraph) *APIResponse {
	nA := &Node{from}
	nB := &Node{to}

	path, distance := getShortestPath(nA, nB, g)
	return &APIResponse{
		Path:     path,
		Distance: distance,
	}
}
