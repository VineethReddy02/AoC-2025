package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strconv"
	"strings"
)

const (
	minValue = 0
	maxValue = 99
	modBase  = maxValue + 1
)

// countZeroClicks counts how many times we click through 0
func countZeroClicks(startPos, delta int) int {
	count := 0
	direction := 1
	if delta < 0 {
		direction = -1
		delta = -delta
	}

	// Click through each position one at a time
	for i := 0; i < delta; i++ {
		startPos += direction
		if startPos < minValue {
			startPos = maxValue
		}
		if startPos > maxValue {
			startPos = minValue
		}
		// If we land exactly on 0, we count it
		if startPos == 0 {
			count++
		}
	}

	return count
}

func main() {
	inputPath := filepath.Join("Dec01", "input.txt")
	file, err := os.Open(inputPath)
	if err != nil {
		log.Fatalf("failed to open %s: %v", inputPath, err)
	}
	defer file.Close()

	current := 50        // starting position
	counter := 0         // times we land exactly on 0 (Puzzle 1)
	totalZeroClicks := 0 // total times dial clicks through 0 (Puzzle 2)

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if line == "" {
			continue
		}

		dir := line[0]
		deltaStr := line[1:]
		delta, err := strconv.Atoi(deltaStr)
		if err != nil {
			log.Fatalf("invalid magnitude %q in line %q: %v", deltaStr, line, err)
		}

		var moveDistance int
		switch dir {
		case 'L', 'l':
			moveDistance = -delta
		case 'R', 'r':
			moveDistance = delta
		default:
			log.Fatalf("unknown direction %q in line %q", string(dir), line)
		}

		// Count landings on 0 for Puzzle 1
		if current == 0 {
			counter++
		}

		// Count zero crossings for Puzzle 2 by each click
		zeroCrossings := countZeroClicks(current, moveDistance)
		totalZeroClicks += zeroCrossings

		// Update current position
		current = current + moveDistance
		current = current % modBase
		if current < 0 {
			current += modBase
		}
	}

	fmt.Println("Puzzle 1: number of times we reach 0: ", counter)
	fmt.Println("Puzzle 2: total times dial clicks through 0 + landings on 0: ", totalZeroClicks)
	if err := scanner.Err(); err != nil {
		log.Fatalf("failed to read %s: %v", inputPath, err)
	}
}
