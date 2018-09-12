package learn

func Add(n, m int) int {
	return n + m
}

type Number int

func (n *Number) AddN(num Number) {
	*n += num
}
