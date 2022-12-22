package main

import (
	"os"
	"strings"
)

func main() {
}

func getFileContents(path string) string {
	data, err := os.ReadFile(path)
	if err != nil {
		panic(err)
	}
	return string(data)
}

func getFewestStep(data string) int {
	graph, startNode, endNode := buildGraph(data)
	if startNode == nil {
		panic("start node not set")
	}
	if endNode == nil {
		panic("end node not set")
	}
	linkNodes(graph)
	queue := make([]*node, 0)
	startNode.isVisited = true
	queue = append(queue, startNode)
	breadthFirstSearch(queue)
	return endNode.step
}

func getFewestStepMultipleStart(data string) int {
	graph, _, endNode := buildGraph(data)
	if endNode == nil {
		panic("end node not set")
	}
	linkNodes(graph)
	queue := make([]*node, 0)
	for _, row := range graph {
		for _, col := range row {
			if col.value == int('a') {
				col.isVisited = true
				queue = append(queue, col)
			}
		}
	}
	breadthFirstSearch(queue)
	// fmt.Printf("%+v\n\n", *&graph)
	return endNode.step
}

func breadthFirstSearch(queue []*node) {
	if len(queue) == 0 {
		return
	}
	current := queue[0]
	if current.isEnd {
		return
	}
	// current.isVisited = true
	queue = queue[1:]
	nextStep := current.step + 1
	if current.bNode != nil && !current.bNode.isVisited {
		current.bNode.isVisited = true
		current.bNode.step = nextStep
		queue = append(queue, current.bNode)
	}
	if current.rNode != nil && !current.rNode.isVisited {
		current.rNode.isVisited = true
		current.rNode.step = nextStep
		queue = append(queue, current.rNode)
	}
	if current.tNode != nil && !current.tNode.isVisited {
		current.tNode.isVisited = true
		current.tNode.step = nextStep
		queue = append(queue, current.tNode)
	}
	if current.lNode != nil && !current.lNode.isVisited {
		current.lNode.isVisited = true
		current.lNode.step = nextStep
		queue = append(queue, current.lNode)
	}
	breadthFirstSearch(queue)
}

func linkNodes(nodes [][]*node) {
	lastRow := len(nodes) - 1
	lastCol := len(nodes[0]) - 1
	for r, row := range nodes {
		for c, col := range row {
			if r+1 <= lastRow && col.value+1 >= nodes[r+1][c].value {
				col.bNode = nodes[r+1][c]
			}
			if c+1 <= lastCol && col.value+1 >= nodes[r][c+1].value {
				col.rNode = nodes[r][c+1]
			}
			if r-1 >= 0 && col.value+1 >= nodes[r-1][c].value {
				col.tNode = nodes[r-1][c]
			}
			if c-1 >= 0 && col.value+1 >= nodes[r][c-1].value {
				col.lNode = nodes[r][c-1]
			}
		}
	}
}

func buildGraph(data string) ([][]*node, *node, *node) {
	graph := make([][]*node, 0)
	var startNode *node
	var endNode *node
	rows := strings.Split(data, "\n")
	for _, row := range rows {
		graphR := make([]*node, 0)
		for _, col := range row {
			init := int(col)
			graphC := node{step: 0}
			if init == int('S') {
				graphC.value = int('a')
				graphC.isStart = true
				startNode = &graphC
			} else if init == int('E') {
				graphC.value = int('z')
				graphC.isEnd = true
				endNode = &graphC
			} else {
				graphC.value = init
			}
			graphR = append(graphR, &graphC)
		}
		graph = append(graph, graphR)
	}
	return graph, startNode, endNode
}

type node struct {
	value     int
	step      int
	isVisited bool
	isStart   bool
	isEnd     bool
	bNode     *node
	rNode     *node
	tNode     *node
	lNode     *node
}
