package chips

import (
	"fmt"
	"sort"
)

type Chips struct {
	current map[int]struct{}
}

func NewChips() *Chips {
	newMap := make(map[int]struct{})
	for i := 1; i <= 12; i++ {
		newMap[i] = struct{}{}
	}

	return &Chips{
		current: newMap,
	}
}

func NewChipsFromSlice(chips []int) *Chips {
	newMap := make(map[int]struct{})
	for _, chip := range chips {
		newMap[chip] = struct{}{}
	}

	return &Chips{
		current: newMap,
	}
}

func (c *Chips) Current() []int {
	result := make([]int, 0, len(c.current))
	for k := range c.current {
		result = append(result, k)
	}

	return result
}

func (c *Chips) Remove(chip int) {
	delete(c.current, chip)
}

func (c *Chips) Has(chip int) bool {
	_, ok := c.current[chip]
	return ok
}

func (c *Chips) HasMulti(chip1, chip2 int) bool {
	return c.Has(chip1) && c.Has(chip2)
}

func (c *Chips) RemoveMulti(chip1, chip2 int) {
	delete(c.current, chip1)
	delete(c.current, chip2)
}

func (c *Chips) String() string {
	s := c.Current()
	sort.Ints(s)
	return fmt.Sprintf("%v", s)
}
