package mapping

type Coord struct {
	Row, Col int
}

func (c Coord) Minus(c1 Coord) Coord {
	return Coord{
		Row: c.Row - c1.Row,
		Col: c.Col - c1.Col,
	}
}

func (c Coord) Plus(delta Coord) Coord {
	return Coord{
		Row: c.Row + delta.Row,
		Col: c.Col + delta.Col,
	}
}

func (c Coord) OrthoNeighbors() []Coord {
	return []Coord{
		c.North(),
		c.East(),
		c.South(),
		c.West(),
	}
}

func (c Coord) North() Coord {
	return Coord{
		Row: c.Row - 1,
		Col: c.Col,
	}
}

func (c Coord) East() Coord {
	return Coord{
		Row: c.Row,
		Col: c.Col + 1,
	}
}

func (c Coord) South() Coord {
	return Coord{
		Row: c.Row + 1,
		Col: c.Col,
	}
}

func (c Coord) West() Coord {
	return Coord{
		Row: c.Row,
		Col: c.Col - 1,
	}
}

func (c Coord) NorthWest() Coord {
	return Coord{
		Row: c.Row - 1,
		Col: c.Col - 1,
	}
}

func (c Coord) NorthEast() Coord {
	return Coord{
		Row: c.Row - 1,
		Col: c.Col + 1,
	}

}

func (c Coord) SouthWest() Coord {
	return Coord{
		Row: c.Row + 1,
		Col: c.Col - 1,
	}
}
