package main

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
	"time"

	"github.com/TiesB/aoc2024-go/utils"
)

func diff(a, b int) int {
	if a > b {
		return a - b
	}

	return b - a
}

func ReadInput() (list1 []int, list2 []int) {
	fs := utils.ScanInput(1)
	defer fs.Close()

	for fs.Scanner.Scan() {
		line := fs.Scanner.Text()
		parts := strings.Fields(line)
		i1, err := strconv.Atoi(parts[0])
		utils.CheckErr(err)
		i2, err := strconv.Atoi(parts[1])
		utils.CheckErr(err)
		list1 = append(list1, i1)
		list2 = append(list2, i2)
	}

	return
}

func SolvePart1() int {
	list1, list2 := ReadInput()

	sort.Ints(list1)
	sort.Ints(list2)

	sum := 0

	for i, v1 := range list1 {
		sum += diff(v1, list2[i])
	}

	return sum
}

func countInts(list []int) map[int]int {
	m := make(map[int]int)
	for _, v := range list {
		m[v] += 1
	}
	return m
}

func SolvePart2() int {
	list1, list2 := ReadInput()

	m := countInts(list2)

	result := 0

	for _, v := range list1 {
		result += v * m[v]
	}

	return result
}

func main() {
	fmt.Println("Day 01")

	start := time.Now()
	part1Result := SolvePart1()
	fmt.Printf("Part 1: %d (%s)\n", part1Result, time.Since(start))
	start = time.Now()
	part2Result := SolvePart2()
	fmt.Printf("Part 2: %d (%s)\n", part2Result, time.Since(start))
}
