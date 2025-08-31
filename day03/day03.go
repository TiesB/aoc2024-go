package day03

import (
	"fmt"
	"regexp"
	"strconv"
	"time"

	"github.com/TiesB/aoc2024-go/utils"
)

func ReadInput() (lines []string) {
	fs := utils.ScanInput(3)
	defer fs.Close()

	for fs.Scanner.Scan() {
		line := fs.Scanner.Text()
		lines = append(lines, line)
	}

	return
}

func SolvePart1() (result int) {
	lines := ReadInput()

	r := regexp.MustCompile(`mul\((?P<A>\d+),(?P<B>\d+)\)`)

	for _, line := range lines {
		matches := r.FindAllStringSubmatch(line, -1)

		for _, match := range matches {
			a, _ := strconv.Atoi(match[1])
			b, _ := strconv.Atoi(match[2])

			c := a * b

			result += c
		}
	}

	return
}

func SolvePart2() (result int) {
	lines := ReadInput()

	r := regexp.MustCompile(`mul\((?P<A>\d+),(?P<B>\d+)\)|don't\(\)|do\(\)`)

	enabled := true

	for _, line := range lines {
		matches := r.FindAllStringSubmatch(line, -1)

		for _, match := range matches {
			switch match[0] {
			case "do()":
				enabled = true
			case "don't()":
				enabled = false
			default:
				if !enabled {
					continue
				}

				a, _ := strconv.Atoi(match[1])
				b, _ := strconv.Atoi(match[2])

				c := a * b

				result += c
			}
		}
	}

	return
}

func Run() {
	fmt.Println("Day 03")

	start := time.Now()
	part1Result := SolvePart1()
	fmt.Printf("Part 1: %d (%s)\n", part1Result, time.Since(start))
	start = time.Now()
	part2Result := SolvePart2()
	fmt.Printf("Part 2: %d (%s)\n", part2Result, time.Since(start))
}
