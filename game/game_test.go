package game

import "testing"

func TestInit(t *testing.T) {
	_ = NewGame()
}

func TestThrowCubes(t *testing.T) {
	g := NewGame()
	g.Throw(1, 1)
}
