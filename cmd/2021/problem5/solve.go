package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
	"strings"
)

func main() {
	// part 1
	p1s, p2s, err := loadInput("input.txt")
	if err != nil {
		log.Fatalf("error loading input.txt file: %v", err)
	}
	result := solvePart1(p1s, p2s)
	fmt.Printf("Part 1 answer: %v", result)
}

type Point struct {
	X int
	Y int
}

func (p *Point) String() string {
	return fmt.Sprintf("%d,%d", p.X, p.Y)
}

func (p *Point) Equal(op *Point) bool {
	return p.X == op.X && p.Y == op.Y
}

func NewPoint(tp string) (Point, error) {
	parts := strings.Split(tp, ",")
	x, err := strconv.Atoi(strings.TrimSpace(parts[0]))
	if err != nil {
		return Point{}, fmt.Errorf("error parsing x value, %v, from text point, %v: %w", parts[0], tp, err)
	}
	y, err := strconv.Atoi(strings.TrimSpace(parts[1]))
	if err != nil {
		return Point{}, fmt.Errorf("error parsing y value, %v, from text point, %v: %w", parts[1], tp, err)
	}
	return Point{X: x, Y: y}, nil
}

func updatePointMap(pm map[string]int, ps []Point) map[string]int {
	for _, p := range ps {
		k := p.String()
		v, ok := pm[k]
		if !ok {
			pm[k] = 1
		} else {
			pm[k] = v + 1
		}
	}
	return pm
}

func getLine(p1 Point, p2 Point) []Point {
	var ps []Point
	// x first
	x := diff(p1.X, p2.X)
	// y next
	y := diff(p1.Y, p2.Y)
	fmt.Println("x = ", x)
	fmt.Println("y = ", y)
	// only horiz lines so just iter over longer slice
	var tot int
	switch {
	case p1.Equal(&p2): // no line, just the same point
		return ps
	case p1.X != p2.X && p1.Y != p2.Y: // diag line
		tot = math.MaxInt
	case len(x) >= len(y):
		tot = len(x)
		for i := 1; i < len(x); i++ {
			y = append(y, p1.Y)
		}
	default:
		tot = len(y)
		for i := 1; i < len(y); i++ {
			x = append(x, p1.X)
		}
	}
	for i := 0; i < tot; i++ {
		ps = append(ps, Point{X: x[i], Y: y[i]})
	}
	return ps
}

func diff(x1 int, x2 int) []int {
	var x []int
	xDiff := int(math.Abs(float64(x1 - x2)))
	min := math.MinInt(x1, x2)
	for i := 0; i < xDiff+1; i++ {
		x = append(x, min+i)
	}
	return x
}

func loadInput(filename string) ([]Point, []Point, error) {
	var p1s, p2s []Point
	file, err := os.Open(filename)
	if err != nil {
		return p1s, p2s, fmt.Errorf("error opening file %v: %w", filename, err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()
		tPoints := strings.Split(line, "->")
		p1, err := NewPoint(tPoints[0])
		if err != nil {
			return p1s, p2s, fmt.Errorf("error creating point from text input, %v: %w", tPoints[0], err)
		}
		p2, err := NewPoint(tPoints[1])
		if err != nil {
			return p1s, p2s, fmt.Errorf("error creating point from text input, %v: %w", tPoints[1], err)
		}
		p1s = append(p1s, p1)
		p2s = append(p2s, p2)
	}
	if err := scanner.Err(); err != nil {
		return p1s, p2s, fmt.Errorf("error scanning file %v: %w", filename, err)
	}

	return p1s, p2s, nil
}

func solvePart1(p1s []Point, p2s []Point) int {
	fp1s, fp2s := filterLines(p1s, p2s)
	pointMap := make(map[string]int)
	// input slices are the same length from input file
	for i := 0; i < len(fp1s); i++ {
		line := getLine(fp1s[i], fp2s[i])
		pointMap = updatePointMap(pointMap, line)
	}
	var tot2ps int
	for _, v := range pointMap {
		if v >= 2 {
			tot2ps++
		}
	}
	return tot2ps
}

func filterLines(p1s []Point, p2s []Point) ([]Point, []Point) {
	var np1s, np2s []Point
	// input slices are the same length from input file
	for i := 0; i < len(p1s); i++ {
		if p1s[i].X == p2s[i].X || p1s[i].Y == p2s[i].Y {
			np1s = append(np1s, p1s[i])
			np2s = append(np2s, p2s[i])
		}
	}
	return np1s, np2s
}
