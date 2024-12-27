package chips

type Chips struct {
}

func NewChips() *Chips {
	return &Chips{}
}

func (c *Chips) Current() []int {
	return []int{}
}
