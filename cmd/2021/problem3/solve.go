package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	inputs, err := loadInput("input.txt")
	if err != nil {
		log.Fatalf("error loading input data: %v", err)
	}
	powerComp, err := solvePart1(inputs)
	if err != nil {
		log.Fatalf("error running solvePart1 func: %v", err)
	}
	fmt.Printf("Answer to part 1: %v", powerComp)
}

func loadInput(filename string) ([]string, error) {
	var input []string
	file, err := os.Open(filename)
	if err != nil {
		return input, fmt.Errorf("couldn't open file, %v: %w", filename, err)
	}
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		input = append(input, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		return input, fmt.Errorf("error while scanning file, %v: %w", filename, err)
	}

	return input, nil
}

func solvePart1(inputs []string) (int, error) {
	postMap := make(map[int]map[string]int)
	for i := 0; i < len(inputs[0]); i++ {
		postMap[i] = map[string]int{
			"0": 0,
			"1": 0,
		}
	}

	for _, nbr := range inputs {
		for j, char := range nbr {
			postMap[j][string(char)] += 1
		}
	}
	// fmt.Printf("postMap: %v\n", postMap)
	var gamstr, epstr string
	for i := 0; i < len(inputs[0]); i++ {
		v := postMap[i]
		if v["1"] > v["0"] {
			gamstr += "1"
			epstr += "0"
		} else {
			gamstr += "0"
			epstr += "1"
		}
	}
	gamma, err := strconv.ParseInt(gamstr, 2, 64)
	if err != nil {
		return 0, fmt.Errorf("couldn't parse gamma binary string, %v: %w", gamstr, err)
	}
	epsilon, err := strconv.ParseInt(epstr, 2, 64)
	if err != nil {
		return 0, fmt.Errorf("couldn't parse epsilon binary string, %v: %w", epstr, err)
	}

	return int(gamma * epsilon), nil
}
