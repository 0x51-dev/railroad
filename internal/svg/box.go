package svg

type Point struct {
	X, Y float64
}

type Box struct {
	Position Point
	Size     Point
}

func NewBox(position Point, size Point) Box {
	if size.X < 0 {
		position.X += size.X
		size.X = -size.X
	}
	if size.Y < 0 {
		position.Y += size.Y
		size.Y = -size.Y
	}
	return Box{
		Position: position,
		Size:     size,
	}
}
func (b Box) Combine(a Box) Box {
	if a.Position.X < b.Position.X {
		b.Size.X += b.Position.X - a.Position.X
		b.Position.X = a.Position.X
	}
	if a.Position.Y < b.Position.Y {
		b.Size.Y += b.Position.Y - a.Position.Y
		b.Position.Y = a.Position.Y
	}
	if a.Position.X+a.Size.X > b.Position.X+b.Size.X {
		b.Size.X = a.Position.X + a.Size.X - b.Position.X
	}
	if a.Position.Y+a.Size.Y > b.Position.Y+b.Size.Y {
		b.Size.Y = a.Position.Y + a.Size.Y - b.Position.Y
	}
	return b
}
