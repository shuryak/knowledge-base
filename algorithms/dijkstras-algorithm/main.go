package main

import (
	"fmt"
	"math"
)

type node string

type weight map[node]int

type graph map[node]weight

func main() {
	graph := make(graph)
	graph["start"] = map[node]int{}
	graph["start"]["a"] = 6
	graph["start"]["b"] = 2

	graph["a"] = weight{}
	graph["a"]["finish"] = 1

	graph["b"] = weight{}
	graph["b"]["a"] = 3
	graph["b"]["finish"] = 5

	graph["finish"] = weight{}

	costs, parents := findShortestPath(graph, "start", "finish")
	fmt.Println(costs, parents)
}

func findShortestPath(graph graph, startNode, finishNode node) (weight, map[node]node) {
	costs := make(weight)
	costs[finishNode] = math.MaxInt32

	parents := make(map[node]node)
	parents[finishNode] = ""

	processed := make(map[node]bool)

	for node, cost := range graph[startNode] {
		costs[node] = cost
		parents[node] = startNode
	}

	lowestCostNode := findLowestCostNode(costs, processed)
	for lowestCostNode != "" {
		for node, cost := range graph[lowestCostNode] {
			newCost := costs[lowestCostNode] + cost
			if newCost < costs[node] {
				costs[node] = newCost
				parents[node] = lowestCostNode
			}
		}

		processed[lowestCostNode] = true
		lowestCostNode = findLowestCostNode(costs, processed)
	}

	return costs, parents
}

func findLowestCostNode(costs weight, processed map[node]bool) node {
	lowestCost := math.MaxInt32
	var lowestCostNode node = ""
	for node, cost := range costs {
		if _, isProcessed := processed[node]; cost < lowestCost && !isProcessed {
			lowestCost = cost
			lowestCostNode = node
		}
	}

	return lowestCostNode
}
