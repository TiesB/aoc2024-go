package grid

type Position struct {
	X, Y int
}

func (p *Position) Transform(direction Direction) Position {
	return p.Add(directionDiff[direction])
}

func (p *Position) TransformScaled(direction Direction, scalor int) Position {
	return p.AddScaled(directionDiff[direction], scalor)
}

func (p *Position) Add(diff PositionDiff) Position {
	return Position{p.X + diff.DX, p.Y + diff.DY}
}

func (p *Position) AddScaled(diff PositionDiff, scalor int) Position {
	return Position{p.X + (diff.DX * scalor), p.Y + (diff.DY * scalor)}
}
