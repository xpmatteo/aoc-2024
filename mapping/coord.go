package mapping

type Coord struct {
	Row, Col int
}

func (c0 Coord) Minus(c1 Coord) Coord {
	return Coord{
		Row: c0.Row - c1.Row,
		Col: c0.Col - c1.Col,
	}
}

func (c Coord) Plus(delta Coord) Coord {
	return Coord{
		Row: c.Row + delta.Row,
		Col: c.Col + delta.Col,
	}
}
