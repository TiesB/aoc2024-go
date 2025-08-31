package day02

import (
	"fmt"
	"slices"
	"strconv"
	"strings"
	"time"

	"github.com/TiesB/aoc2024-go/utils"
)

func ReadInput() [][]int {
	var data [][]int
	fs := utils.ScanInput(2)
	defer fs.Close()

	for fs.Scanner.Scan() {
		line := fs.Scanner.Text()
		parts := strings.Fields(line)
		data = append(data, utils.MapSlice(parts, func(s string) int {
			i, _ := strconv.Atoi(s)
			return i
		}))
	}

	return data
}

func testRow(row []int, isSecond bool) (bool, int) {
	ascending := row[1] > row[0]
	for i := 1; i < len(row); i++ {
		a, b := row[i-1], row[i]
		diff := utils.DiffInt(a, b)
		if diff < 1 || diff > 3 {
			return false, i
		}
		if ascending && b < a {
			return false, i

		} else if !ascending && b > a {
			return false, i
		}
	}

	return true, -1
}

func SolvePart1() int {
	data := ReadInput()

	result := 0

	for _, row := range data {
		good, _ := testRow(row, false)

		if good {
			result += 1
		}
	}

	return result
}

func copySliceWithoutElement[T any](s []T, i int) []T {
	return slices.Delete(slices.Clone(s), i, i+1)
}

func SolvePart2() int {
	data := ReadInput()

	result := 0

	for _, row := range data {
		good, i := testRow(row, false)

		if good {
			result += 1
		} else {
			if i > 1 {
				if good, _ := testRow(copySliceWithoutElement(row, i-2), true); good {
					result += 1
					continue
				}
			}
			if good, _ := testRow(copySliceWithoutElement(row, i-1), true); good {
				result += 1
				continue
			}
			if good, _ := testRow(copySliceWithoutElement(row, i), true); good {
				result += 1
				continue
			}
		}
	}
	return result
}

func Run() {
	fmt.Println("Day 02")

	start := time.Now()
	part1Result := SolvePart1()
	fmt.Printf("Part 1: %d (%s)\n", part1Result, time.Since(start))
	start = time.Now()
	part2Result := SolvePart2()
	fmt.Printf("Part 2: %d (%s)\n", part2Result, time.Since(start))
}
