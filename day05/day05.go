package day05

import (
	"fmt"
	"slices"
	"strconv"
	"strings"
	"time"

	"github.com/TiesB/aoc2024-go/utils"
)

func ReadInput() (map[int]utils.Set[int], [][]int) {
	fs := utils.ScanInput(5)
	defer fs.Close()

	rules := make(map[int]utils.Set[int])
	var updates [][]int

	for fs.Scanner.Scan() {
		line := fs.Scanner.Text()

		if line == "" {
			break
		}

		parts := strings.Split(line, "|")

		first, err := strconv.Atoi(parts[0])
		if err != nil {
			panic(err)
		}

		second, err := strconv.Atoi(parts[1])
		if err != nil {
			panic(err)
		}

		if _, exists := rules[second]; !exists {
			rules[second] = make(utils.Set[int])
		}

		rules[second].Add(first)
	}

	for fs.Scanner.Scan() {
		line := fs.Scanner.Text()
		var update []int

		for _, part := range strings.Split(line, ",") {
			i, err := strconv.Atoi(part)
			if err != nil {
				panic(err)
			}

			update = append(update, i)
		}
		updates = append(updates, update)
	}

	return rules, updates
}

func SolvePart1() (result int) {
	rules, updates := ReadInput()

UpdatesLoop:
	for _, update := range updates {
		contents := make(utils.Set[int])
		contents.AddAll(update)
		seen := make(utils.Set[int])

		for _, i := range update {
			seen.Add(i)
			requirements, exists := rules[i]
			if !exists {
				continue
			}

			for req := range requirements {
				if contents.Has(req) && !seen.Has(req) {
					continue UpdatesLoop
				}
			}
		}

		middleItem := update[len(update)/2]
		result += middleItem
	}

	return
}

func SolvePart2() (result int) {
	rules, updates := ReadInput()

	for _, update := range updates {
		contents := make(utils.Set[int])
		contents.AddAll(update)
		seen := make(utils.Set[int])
		correct := true

	UpdateItemsLoop:
		for _, i := range update {
			seen.Add(i)
			requirements, exists := rules[i]
			if !exists {
				continue
			}

			for req := range requirements {
				if contents.Has(req) && !seen.Has(req) {
					correct = false
					continue UpdateItemsLoop
				}
			}
		}

		if correct {
			continue
		}

		slices.SortFunc(update, func(a, b int) int {
			if rules[a].Has(b) {
				return -1
			}
			if rules[b].Has(a) {
				return 1
			}
			return 0
		})

		middleItem := update[len(update)/2]
		result += middleItem
	}

	return
}

func Run() {
	fmt.Println("Day 05")

	start := time.Now()
	part1Result := SolvePart1()
	fmt.Printf("Part 1: %d (%s)\n", part1Result, time.Since(start))
	start = time.Now()
	part2Result := SolvePart2()
	fmt.Printf("Part 2: %d (%s)\n", part2Result, time.Since(start))
}
