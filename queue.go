package main

type Node struct {
	Value string
}

type Edge struct {
	Node   *Node
	Weight int
}

type Vertex struct {
	Node     *Node
	Distance int
}

type PriorityQueue []*Vertex

type InputGraph struct {
	Graph []InputData `json:"graf"`
	From  string      `json:"od"`
	To    string      `json:"do"`
}

type InputData struct {
	Source      string `json:"izvor"`
	Destination string `json:"destinacija"`
	Weight      int    `json:"tezina"`
}

type NodeQueue struct {
	Items []Vertex
}

type APIResponse struct {
	Path     []string `json:"putanja"`
	Distance int      `json:"distanca"`
}

func (s *NodeQueue) Enqueue(t Vertex) {
	if len(s.Items) == 0 {
		s.Items = append(s.Items, t)
		return
	}
	var insertFlag bool
	for k, v := range s.Items {
		if t.Distance < v.Distance {
			s.Items = append(s.Items[:k+1], s.Items[k:]...)
			s.Items[k] = t
			insertFlag = true
		}
		if insertFlag {
			break
		}
	}
	if !insertFlag {
		s.Items = append(s.Items, t)
	}
}

// Funkcija Dequeue uklanja prvi čvor iz prioritetnog reda.
func (s *NodeQueue) Dequeue() *Vertex {
	item := s.Items[0]
	s.Items = s.Items[1:len(s.Items)]
	return &item
}

// Funkcija NewQueue kreira novi prioritetni red.
func (s *NodeQueue) NewQueue() *NodeQueue {
	s.Items = []Vertex{}
	return s
}

func (s *NodeQueue) IsEmpty() bool {
	return len(s.Items) == 0
}

func (s *NodeQueue) Size() int {
	return len(s.Items)
}
