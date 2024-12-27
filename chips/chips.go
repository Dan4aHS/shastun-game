package chips

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
