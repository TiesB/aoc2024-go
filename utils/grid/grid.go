package grid

import (
	"errors"
	"iter"
)

type Grid []string

type PositionDiff struct {
	DX, DY int
}

type Direction int

const (
	North Direction = iota
	East
	South
	West

	NorthEast
	SouthEast
	SouthWest
	NorthWest
)

var Directions = [...]Direction{North, East, South, West, NorthEast, SouthEast, SouthWest, NorthWest}

var directionDiff = map[Direction]PositionDiff{
	North: {0, -1},
	East:  {1, 0},
	South: {0, 1},
	West:  {-1, 0},

	NorthEast: {1, -1},
	SouthEast: {1, 1},
	SouthWest: {-1, 1},
	NorthWest: {-1, -1},
}

func (g *Grid) Dimensions() (int, int, error) {
	if g == nil {
		return -1, -1, errors.New("grid is nil")
	}

	return len((*g)[0]), len(*g), nil
}

func (g *Grid) Positions() iter.Seq[Position] {
	return func(yield func(Position) bool) {
		height, width, err := g.Dimensions()

		if err != nil {
			return
		}

		for y := 0; y < height; y++ {
			for x := 0; x < width; x++ {
				if !yield(Position{X: x, Y: y}) {
					return
				}
			}
		}
	}
}

func (g *Grid) Items() iter.Seq2[Position, rune] {
	return func(yield func(Position, rune) bool) {
		for position := range g.Positions() {
			if item, err := g.At(position); err != nil || !yield(position, item) {
				return
			}
		}
	}
}

func (g *Grid) At(p Position) (rune, error) {
	if p.X < 0 || p.Y < 0 {
		return 0, errors.New("grid does not support negative indices")
	}

	if height, width, err := g.Dimensions(); err != nil || height <= p.Y || width <= p.X {
		if err != nil {
			return 0, err
		}

		return 0, errors.New("grid is too small")
	}

	return rune((*g)[p.Y][p.X]), nil
}

func (g *Grid) Neighbour(pos Position, dir Direction) (rune, error) {
	return g.At(pos.Transform(dir))
}

func (g *Grid) NeighboursInDirection(pos Position, dir Direction, n int) ([]rune, error) {
	var result []rune

	for i := 0; i < n; i++ {
		pos = pos.Transform(dir)
		r, err := g.At(pos)

		if err != nil {
			return nil, err
		}

		result = append(result, r)
	}

	return result, nil
}
