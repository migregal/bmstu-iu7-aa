package graph

import (
	"math"
	"math/rand"
	"time"
)

func (c *AntColony) SearchAnt(d int) []int {
	r := make([]int, len(c.w))

	for i := 0; i < d; i++ {
		for j := 0; j < len(c.w); j++ {
			a := c.CreateAnt(j)
			a.moveAnt()
			cur := a.getDistance()

			if (r[j] == 0) || (cur < r[j]) {
				r[j] = cur
			}
		}
	}

	return r
}

func (a *Ant) moveAnt() {
	for {
		prob := a.getProb()
		way := getWay(prob)
		if way == -1 {
			break
		}
		a.follow(way)
		a.updatePh()
	}
}

func (a *Ant) getProb() []float64 {
	var sum float64

	p := make([]float64, 0)

	for i, l := range a.vis[a.pos] {
		if l == 0 {
			p = append(p, 0)
		} else {
			d := math.Pow((1.0/float64(l)), a.col.a) *
				math.Pow(a.col.ph[a.pos][i], a.col.b)
			p = append(p, d)
			sum += d
		}
	}

	for _, l := range p {
		l /= sum
	}

	return p
}

func (a *Ant) getDistance() int {
	d := 0

	for i, j := range a.isv {
		for k, z := range j {
			if z {
				d += a.col.w[i][k]
			}
		}
	}

	return d
}

func (a *Ant) updatePh() {
	delta := 0.0

	for i := 0; i < len(a.col.ph); i++ {
		for j, ph := range a.col.ph[i] {
			if a.col.w[i][j] != 0 {
				if a.isv[i][j] {
					delta = a.col.q / float64(a.col.w[i][j])
				} else {
					delta = 0
				}
				a.col.ph[i][j] = (1 - a.col.p) * (float64(ph) + delta)
			}

			if a.col.ph[i][j] <= 0 {
				a.col.ph[i][j] = 0.1
			}
		}
	}
}

func (a *Ant) follow(path int) {
	for i := range a.vis {
		a.vis[i][a.pos] = 0
	}
	a.isv[a.pos][path] = true
	a.pos = path
}

func getWay(p []float64) int {
	var sum, rn float64
	var r *rand.Rand

	for _, j := range p {
		sum += j
	}

	r = rand.New(rand.NewSource(time.Now().UnixNano()))
	rn = r.Float64() * sum
	sum = 0

	for i, val := range p {
		if rn > sum && rn < sum+val {
			return i
		}
		sum += val
	}

	return -1
}
