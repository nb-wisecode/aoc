package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	// load puzzle input
	comms, err := loadSubCommands("input.txt")
	if err != nil {
		log.Fatalf("couldn't load puzzle input: %v", err)
	}
	result := solvePart1(comms)
	fmt.Printf("Answer for Part 1: %v\n", result)
	result = solvePart2(comms)
	fmt.Printf("Answer for Part 2: %v\n", result)
}

type SubCommand struct {
	Type   string
	Amount int
}

func loadSubCommands(filename string) ([]SubCommand, error) {
	var comms []SubCommand
	file, err := os.Open(filename)
	if err != nil {
		return comms, fmt.Errorf("couldn't open file, %v: %w", filename, err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.Split(line, " ")
		amt, err := strconv.Atoi(parts[1])
		if err != nil {
			return comms, fmt.Errorf("couldn't read amount %v: %w", parts[1], err)
		}
		comms = append(comms, SubCommand{Type: parts[0], Amount: amt})
	}

	if err := scanner.Err(); err != nil {
		return comms, fmt.Errorf("couldn't read file, %v: %w", filename, err)
	}

	return comms, nil
}

func solvePart1(input []SubCommand) int {
	var horiz, depth int
	for _, comm := range input {
		switch comm.Type {
		case "forward":
			horiz += comm.Amount
		case "up":
			depth -= comm.Amount
		case "down":
			depth += comm.Amount
		default:
			log.Fatalf("unrecognized command %v", comm.Type)
		}
	}
	return horiz * depth
}

func solvePart2(input []SubCommand) int {
	var horiz, depth, aim int
	for _, comm := range input {
		switch comm.Type {
		case "forward":
			horiz += comm.Amount
			depth += comm.Amount * aim
		case "up":
			aim -= comm.Amount
		case "down":
			aim += comm.Amount
		default:
			log.Fatalf("unrecognized command %v", comm.Type)
		}
	}
	return horiz * depth
}
