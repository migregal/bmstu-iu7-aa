package graph

type AntColony struct {
	w          [][]int
	ph         [][]float64
	a, b, q, p float64
}

type Ant struct {
	col *AntColony
	vis [][]int
	isv [][]bool
	pos int
}

func (c *AntColony) CreateAnt(pos int) *Ant {
	a := Ant{
		col: c,
		vis: make([][]int, len(c.w)),
		isv: make([][]bool, len(c.w)),
		pos: pos,
	}
	for i := 0; i < len(c.w); i++ {
		a.vis[i] = make([]int, len(c.w))
		for j := 0; j < len(c.w[i]); j++ {
			a.vis[i][j] = c.w[i][j]
		}
	}
	for i := range a.isv {
		a.isv[i] = make([]bool, len(c.w))
	}
	return &a
}

func CreateColony(w [][]int) *AntColony {
	c := AntColony{
		w:  w,
		ph: make([][]float64, len(w)),
		a:  3.0, b: 7.0, q: 20.0, p: 0.6,
	}
	for i := 0; i < len(c.ph); i++ {
		c.ph[i] = make([]float64, len(c.w[i]))
		for j := range c.ph[i] {
			c.ph[i][j] = 0.5
		}
	}
	return &c
}
