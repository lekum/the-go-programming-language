package counter

type Counter struct{ n int }

func (c *Counter) N() int {
	return c.n
}
func (c *Counter) Increment() {
	c.n++
}
func (c *Counter) Reset() {
	c.n = 0
}
