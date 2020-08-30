package chapter4

import (
	"container/list"
)

type State int

const (
	Unvisited State = iota
	Visited
	Visiting
)

type Graph struct {
	nodes []*Node
}

type Node struct {
	name     string
	children []*Node
	state    State
}

func ExistRoot(g *Graph, start, end *Node) bool {
	if start == end {
		return true
	}
	// init queue
	q := list.New()

	for _, u := range g.nodes {
		u.state = Unvisited
	}
	q.PushBack(start)
	for q.Len() != 0 {
		e := q.Remove(q.Front())
		u := e.(*Node)
		if u == nil {
			continue
		}
		for _, v := range u.children {
			if v.state != Unvisited {
				continue
			}
			if v == end {
				return true
			} else {
				v.state = Visiting
				q.PushBack(v)
			}
		}
		u.state = Visited
	}
	return false
}
