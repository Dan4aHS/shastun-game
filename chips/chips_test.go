package chips

import "testing"

func TestInit(t *testing.T) {
	_ = NewChips()
}

func TestCurrentChips(t *testing.T) {
	c := NewChips()
	c.Current()
}
