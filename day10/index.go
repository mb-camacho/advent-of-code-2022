package main

import (
	"fmt"
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

func sumOfSignalStrengths(contents string) int {
	commands := buildCommands(contents)
	queueCommands := queueCommands(commands)
	registerX := 1
	strengthRegisters := make([]int, 0)
	fmt.Print("\n")
	for i, command := range queueCommands {
		cycle := (i + 1)
		pixelValue := ""
		if i%40 >= registerX-1 && i%40 <= registerX+1 {
			pixelValue = "#"
		} else {
			pixelValue = "."
		}
		fmt.Print(pixelValue)
		strengthRegister := registerX * (i + 1)
		strengthRegisters = append(strengthRegisters, strengthRegister)
		if command.keyword == "addx" {
			registerX += command.arg
		}
		if cycle%40 == 0 {
			fmt.Print("\n")
			// fmt.Print("###...###...###...###...###...###...###.\n")
			// fmt.Print("####....####....####....####....####....\n")
			// fmt.Print("#####.....#####.....#####.....#####.....\n")
			// fmt.Print("######......######......######......####\n")
			// fmt.Print("#######.......#######.......#######.....\n\n")
		}
	}
	fmt.Print("\n")
	return strengthRegisters[19] + strengthRegisters[59] + strengthRegisters[99] + strengthRegisters[139] + strengthRegisters[179] + strengthRegisters[219]
}

func queueCommands(commands []command) []command {
	queueCommands := make([]command, 0)
	for _, com := range commands {
		if com.keyword == "addx" {
			queueCommands = append(queueCommands, command{keyword: "next: " + com.keyword})
		}
		queueCommands = append(queueCommands, com)
	}
	return queueCommands
}

func buildCommands(data string) []command {
	dataArray := strings.Split(data, "\n")
	commands := make([]command, 0)
	for _, s := range dataArray {
		p := strings.Split(s, " ")
		if len(p) <= 0 {
			panic("expecting > 0. command: " + s)
		}
		if p[0] == "noop" {
			com := command{keyword: p[0]}
			commands = append(commands, com)
			continue
		}
		if p[0] == "addx" {
			if len(p) != 2 {
				panic("expecting len of 2 for addx. command" + s)
			}
			arg, err := strconv.Atoi(p[1])
			if err != nil {
				panic(err)
			}
			com := command{keyword: p[0], arg: arg}
			commands = append(commands, com)
			continue
		}
		panic("unknown command: " + s)
	}
	return commands
}

type command struct {
	keyword string
	arg     int
}
