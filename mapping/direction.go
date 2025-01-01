package mapping

type Direction string

const (
	DirectionNorth = Direction("N")
	DirectionEast  = Direction("E")
	DirectionSouth = Direction("S")
	DirectionWest  = Direction("W")
	DirectionNone  = Direction("")
)

func (d Direction) Opposite() Direction {
	var opposites = map[Direction]Direction{
		DirectionNorth: DirectionSouth,
		DirectionEast:  DirectionWest,
		DirectionSouth: DirectionNorth,
		DirectionWest:  DirectionEast,
	}
	if result, ok := opposites[d]; ok {
		return result
	}
	panic("bad direction '" + d + "'")
}
