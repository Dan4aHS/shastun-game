package chips

type Chips struct {
	current []int
}

func NewChips() *Chips {
	return &Chips{
		current: []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12},
	}
}

func (c *Chips) Current() []int {
	return c.current
}
