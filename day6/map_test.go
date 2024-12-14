package day6

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_set(t *testing.T) {
	m := Map([]string{"aaa"})
	m.Set(0, 1, 'X')
	assert.Equal(t, Map([]string{"aXa"}), m)
}
