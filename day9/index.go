package main

import (
	"os"
	"strconv"
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

func countPositionsTailVisited(contents string) int {
	commands := buildCommands(contents)
	head := node{history: make(map[string]int, 0)}
	updateHistory(&head)
	tail := node{history: make(map[string]int, 0)}
	updateHistory(&tail)
	for _, command := range commands {
		process(&head, command, &tail)
	}
	return len(tail.history)
}

func countPositionsTail9Visited(contents string) int {
	commands := buildCommands(contents)
	var tails []*node
	for i := 0; i < 10; i++ {
		tail := node{history: make(map[string]int, 0), name: strconv.Itoa(i)}
		updateHistory(&tail)
		tails = append(tails, &tail)
	}

	for _, command := range commands {
		processMultiple(command, tails)
	}
	return len(tails[9].history)
}

func process(head *node, move command, tail *node) {
	for i := move.steps; i > 0; i-- {
		moveHead(head, move.direction)
		updateTail(tail, head)
	}
}

func processMultiple(move command, tails []*node) {
	for i := move.steps; i > 0; i-- {
		moveHead(tails[0], move.direction)
		// doNotChangeFlag := false
		for j := 1; j < len(tails); j++ {
			originalX := tails[j].x
			originalY := tails[j].y
			updateTail(tails[j], tails[j-1])
			// if doNotChangeFlag && (originalX != tails[j].x && originalY != tails[j].y) {
			// 	fmt.Printf("should not change %v\n", j-1)
			// }
			// if doNotChangeFlag {
			// 	continue
			// }
			if originalX == tails[j].x && originalY == tails[j].y {
				// doNotChangeFlag = true
				break
			}
		}
		// fmt.Printf("tail: %v,%v\n", tails[9].x, tails[9].y)
	}
}

func updateTail(tail *node, head *node) {
	if head.x == tail.x && head.y == tail.y {
		return
	}

	diagonal := false
	if tail.x != head.x && tail.y != head.y {
		diagonal = true
	}
	allowedDistance := 2
	newX := tail.x
	newY := tail.y
	if tail.x+allowedDistance == head.x {
		newX++
	} else if tail.x-allowedDistance == head.x {
		newX--
	} else if tail.y+allowedDistance == head.y {
		newY++
	} else if tail.y-allowedDistance == head.y {
		newY--
	} else {
		return
	}

	if diagonal {
		if newX == tail.x {
			if head.x > newX {
				newX++
			} else if head.x < newX {
				newX--
			}
		} else if newY == tail.y {
			if head.y > newY {
				newY++
			} else if head.y < newY {
				newY--
			}
		}
	}
	tail.x = newX
	tail.y = newY
	updateHistory(tail)
	return
}

func moveHead(head *node, direction string) {
	if direction == "R" {
		head.x++
		updateHistory(head)
		return
	}
	if direction == "L" {
		head.x--
		updateHistory(head)
		return
	}
	if direction == "U" {
		head.y++
		updateHistory(head)
		return
	}
	if direction == "D" {
		head.y--
		updateHistory(head)
		return
	}
	panic("unknown direction")
}

func updateHistory(node *node) {
	key := strconv.Itoa(node.x) + "," + strconv.Itoa(node.y)
	_, ok := node.history[key]
	if ok {
		node.history[key]++
		return
	}
	node.history[key] = 1
}

func buildCommands(data string) []command {
	dataArray := strings.Split(data, "\n")
	commands := make([]command, 0)
	for _, s := range dataArray {
		parts := strings.Split(s, " ")
		if len(parts) != 2 {
			panic("expected length of parts is 2")
		}
		steps, err := strconv.Atoi(parts[1])
		if err != nil {
			panic(err)
		}
		commands = append(commands, command{direction: parts[0], steps: steps})
	}
	return commands
}

type command struct {
	direction string
	steps     int
}

type node struct {
	name    string
	x       int
	y       int
	history map[string]int
}
