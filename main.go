package main

import (
	"fmt"
	"time"

	"github.com/TiesB/aoc2024-go/day01"
	"github.com/TiesB/aoc2024-go/day02"
	"github.com/TiesB/aoc2024-go/day03"
	"github.com/TiesB/aoc2024-go/day04"
	"github.com/TiesB/aoc2024-go/day05"
)

func main() {
	start := time.Now()
	day01.Run()
	day02.Run()
	day03.Run()
	day04.Run()
	day05.Run()
	fmt.Printf("Total time: %s\n", time.Since(start))
}
