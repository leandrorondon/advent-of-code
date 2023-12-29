package cache

type TwoD struct {
	m map[int]map[int]int
}

func New2D() *TwoD {
	return &TwoD{make(map[int]map[int]int)}
}

func (c *TwoD) Get(a, b int) (int, bool) {
	m, ok := c.m[a]
	if !ok {
		c.m[a] = make(map[int]int)
	}

	if v, ok := m[b]; ok {
		return v, true
	}

	return 0, false

}

func (c *TwoD) Set(a, b, v int) {
	_, ok := c.m[a]
	if !ok {
		c.m[a] = make(map[int]int)
	}

	c.m[a][b] = v
}
