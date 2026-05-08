package main

import (
	"container/heap"
	"fmt"
	"math"
)

type Edge struct {
	to     string
	weight int
}

type Item struct {
	node     string
	distance int
	index    int
}

type PriorityQueue []*Item

func (pq PriorityQueue) Len() int {
	return len(pq)
}

func (pq PriorityQueue) Less(i, j int) bool {
	return pq[i].distance < pq[j].distance
}

func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
	pq[i].index = i
	pq[j].index = j
}

func (pq *PriorityQueue) Push(x interface{}) {
	n := len(*pq)
	item := x.(*Item)
	item.index = n
	*pq = append(*pq, item)
}

func (pq *PriorityQueue) Pop() interface{} {
	old := *pq
	n := len(old)
	item := old[n-1]
	item.index = -1
	*pq = old[0 : n-1]
	return item
}

func dijkstra(graph map[string][]Edge, start string) map[string]int {
	costs := make(map[string]int)
	for node := range graph {
		costs[node] = math.MaxInt32
	}

	costs[start] = 0
	pq := &PriorityQueue{}

	heap.Init(pq)
	heap.Push(pq, &Item{node: start, distance: 0})
	for pq.Len() > 0 {
		current := heap.Pop(pq).(*Item)
		if current.distance > costs[current.node] {
			continue
		}

		for _, edge := range graph[current.node] {
			newCost := current.distance + edge.weight
			if newCost < costs[edge.to] {
				costs[edge.to] = newCost
				heap.Push(pq, &Item{node: edge.to, distance: newCost})
			}
		}
	}

	return costs
}

func main() {
	graph := map[string][]Edge{
		"Start": {{to: "A", weight: 6}, {to: "B", weight: 2}},
		"A":     {{to: "Start", weight: 6}, {to: "Fin", weight: 1}, {to: "B", weight: 3}},
		"B":     {{to: "Start", weight: 2}, {to: "Fin", weight: 5}, {to: "A", weight: 3}},
		"Fin":   {{to: "A", weight: 1}, {to: "B", weight: 5}},
	}

	distances := dijkstra(graph, "Start")
	fmt.Println("Shortest distances from Start:")
	for node, dist := range distances {
		fmt.Printf("Node %s: %d\n", node, dist)
	}
}
