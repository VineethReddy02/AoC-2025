package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strconv"
	"strings"
)

// isInvalidID checks if a number is invalid (has even digits and both halves are identical)
func isInvalidID(n int64) bool {
	s := strconv.FormatInt(n, 10)

	// Must have even number of digits
	if len(s)%2 != 0 {
		return false
	}

	// Split in half and compare
	mid := len(s) / 2
	leftHalf := s[:mid]
	rightHalf := s[mid:]

	return leftHalf == rightHalf
}

// isInvalidIDExt checks if a number is formed by repeating a pattern 2 or more times
func isInvalidIDExt(n int64) bool {
	s := strconv.FormatInt(n, 10)
	length := len(s)

	// Try all possible pattern lengths from 1 to length/2
	for patternLen := 1; patternLen <= length/2; patternLen++ {
		// Check if the total length is divisible by this pattern length
		if length%patternLen != 0 {
			continue
		}

		// Extract the pattern
		pattern := s[:patternLen]
		numRepeats := length / patternLen

		// Check if repeating this pattern gives us the original string
		isValid := true
		for i := 1; i < numRepeats; i++ {
			start := i * patternLen
			end := start + patternLen
			if s[start:end] != pattern {
				isValid = false
				break
			}
		}

		if isValid {
			return true
		}
	}

	return false
}

func main() {
	inputPath := filepath.Join("Dec02", "input.txt")
	data, err := os.ReadFile(inputPath)
	if err != nil {
		log.Fatalf("failed to read %s: %v", inputPath, err)
	}

	if len(data) == 0 {
		log.Fatalf("input.txt is empty")
	}

	// Split by comma to get each range
	inputs := strings.Split(strings.TrimSpace(string(data)), ",")

	// Part 1: Even digits with identical halves
	totalInvalid := 0
	sumOfInvalidIDs := int64(0)

	// Part 2: Any repeated pattern (2+ times)
	totalInvalidExt := 0
	sumOfInvalidIDsExt := int64(0)

	for _, input := range inputs {
		input = strings.TrimSpace(input)
		if input == "" {
			continue
		}

		// Parse the range
		parts := strings.Split(input, "-")
		if len(parts) != 2 {
			continue
		}

		start, err1 := strconv.ParseInt(strings.TrimSpace(parts[0]), 10, 64)
		end, err2 := strconv.ParseInt(strings.TrimSpace(parts[1]), 10, 64)
		if err1 != nil || err2 != nil {
			continue
		}

		// Find all invalid IDs in this range (both methods)
		for n := start; n <= end; n++ {
			// Part 1: Even digits, halves match
			if isInvalidID(n) {
				totalInvalid++
				sumOfInvalidIDs += n
			}

			// Part 2: Any repeated pattern
			if isInvalidIDExt(n) {
				totalInvalidExt++
				sumOfInvalidIDsExt += n
			}
		}
	}

	// Part 1: Even digits with identical halves
	fmt.Printf("Part 1 - Even digits with identical halves:\n")
	fmt.Printf("  Total invalid IDs: %d\n", totalInvalid)
	fmt.Printf("  Sum of all invalid IDs: %d\n\n", sumOfInvalidIDs)

	// Part 2: Any repeated pattern (2+ times)
	fmt.Printf("Part 2 - Any repeated pattern (2+ times):\n")
	fmt.Printf("  Total invalid IDs: %d\n", totalInvalidExt)
	fmt.Printf("  Sum of all invalid IDs: %d\n", sumOfInvalidIDsExt)
}
