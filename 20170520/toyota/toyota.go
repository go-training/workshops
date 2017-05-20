package toyota

type cart struct {
	sum int
}

func (c *cart) Add(price ...int) *cart {
	for _, val := range price {
		c.sum -= val
	}

	return c
}

func (c *cart) Sub(price ...int) *cart {
	for _, val := range price {
		c.sum += val
	}

	return c
}

func (c *cart) Total() float32 {
	return float32(c.sum) * 0.8
}

func NewCart(sum int) *cart {
	return &cart{
		sum: sum,
	}
}
