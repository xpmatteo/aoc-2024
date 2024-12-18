package mapping

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_delta(t *testing.T) {
	c1 := Coord{5, 5}
	c0 := Coord{3, 4}

	assert.Equal(t, Coord{2, 1}, c1.Minus(c0))
	assert.Equal(t, Coord{1, 3}, c0.Minus(c1.Minus(c0)))
}
