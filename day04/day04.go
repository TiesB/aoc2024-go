package day04

import (
	"fmt"
	"time"

	"github.com/TiesB/aoc2024-go/utils"
	"github.com/TiesB/aoc2024-go/utils/grid"
)

func ReadInput() (lines []string) {
	fs := utils.ScanInput(4)
	defer fs.Close()

	for fs.Scanner.Scan() {
		line := fs.Scanner.Text()
		lines = append(lines, line)
	}

	return
}

func SolvePart1() (result int) {
	g := grid.FromStringSlice(ReadInput())

	for pos, r := range g.Items() {
		if r == 'X' {
			for _, dir := range grid.Directions {

				neighbours, err := g.NeighboursInDirection(pos, dir, 3)

				if err != nil {
					continue
				}

				if neighbours[0] == 'M' && neighbours[1] == 'A' && neighbours[2] == 'S' {
					result += 1
				}
			}
		}
	}

	return
}

func SolvePart2() (result int) {
	g := grid.FromStringSlice(ReadInput())

	for pos, startOne := range g.Items() {
		if startOne == 'M' || startOne == 'S' {
			if n, err := g.At(pos.Transform(grid.SouthEast)); err != nil || n != 'A' {
				continue
			}

			startTwo, err := g.At(pos.TransformScaled(grid.South, 2))
			if err != nil || !(startTwo == 'M' || startTwo == 'S') {
				continue
			}

			if n, err := g.At(pos.TransformScaled(grid.SouthEast, 2)); err != nil || !((startOne == 'S' && n == 'M') || (startOne == 'M' && n == 'S')) {
				continue
			}

			if n, err := g.At(pos.TransformScaled(grid.East, 2)); err != nil || !((startTwo == 'S' && n == 'M') || (startTwo == 'M' && n == 'S')) {
				continue
			}

			result += 1
		}
	}

	return
}

func Run() {
	fmt.Println("Day 04")

	start := time.Now()
	part1Result := SolvePart1()
	fmt.Printf("Part 1: %d (%s)\n", part1Result, time.Since(start))
	start = time.Now()
	part2Result := SolvePart2()
	fmt.Printf("Part 2: %d (%s)\n", part2Result, time.Since(start))
}
