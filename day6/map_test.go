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

func Test_clone_returns_equal_object(t *testing.T) {
	m := Map([]string{"aaa", "bbb"})
	clone := m.Clone()

	assert.Equal(t, m, clone)
}

func Test_clones_can_be_modified_independently(t *testing.T) {
	m := Map([]string{"aaa", "bbb"})
	clone := m.Clone()

	clone.Set(0, 1, 'X')

	assert.Equal(t, Map([]string{"aaa", "bbb"}), m, "original is unchanged")
}
