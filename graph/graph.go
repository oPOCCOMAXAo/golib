package graph

import (
	"math"
	"sort"
)

type Graph struct {
	points    []Point
	max       int64
	col       int64
	row       int64
	buffer    [][]rune
	rawbuffer []rune
	cstats    []ColStats
	maxvalue  int64
}

type Point struct {
	X int64
	Y int64
}

type ColStats struct {
	Max int64
	Min int64
	Avg int64
}

func CreateGraph(maxPts int64) Graph {
	res := Graph{
		points: make([]Point, 0),
		max:    maxPts,
		col:    80,
		row:    20,
	}
	res.rawbuffer = make([]rune, (res.col+1)*res.row)
	res.buffer = make([][]rune, res.row)
	for i := range res.buffer {
		min := int64(i) * (res.col + 1)
		max := min + res.col
		res.buffer[i] = res.rawbuffer[min:max]
		res.rawbuffer[max] = '\n'
	}
	res.rawbuffer = res.rawbuffer[:len(res.rawbuffer)]
	res.clear()
	res.cstats = make([]ColStats, res.col)
	return res
}

func ptsMinMaxAvg(pts []Point) (min int64, max int64, avg int64) {
	count := len(pts)
	min = math.MaxInt64
	for _, v := range pts {
		if v.Y < min {
			min = v.Y
		}
		if v.Y > max {
			max = v.Y
		}
		avg += v.Y
	}
	if count > 0 {
		avg /= int64(count)
	}
	return
}

func (g *Graph) Add(p Point) {
	g.points = append(g.points, p)
	g.sortPts()
	i := int64(len(g.points)) - g.max
	if i > 0 {
		g.points = g.points[i:]
	}
}

func (g *Graph) sortPts() {
	sort.Slice(g.points, func(i, j int) bool {
		return g.points[i].X < g.points[j].X
	})
}

const (
	HARD   = '▓'
	MEDIUM = '▒'
	LIGHT  = '░'
	FREE   = ' '
)

func (g *Graph) clear() {
	for _, v := range g.buffer {
		for j := range v {
			v[j] = ' '
		}
	}
}
func (g *Graph) calcColStats() {
	d := float64(len(g.points)) / float64(g.col)
	g.maxvalue = 0
	for i := range g.cstats {
		min := int(math.Floor(d * float64(i)))
		max := int(math.Ceil(d * float64(i+1)))
		g.cstats[i].Min, g.cstats[i].Max, g.cstats[i].Avg = ptsMinMaxAvg(g.points[min:max])
		if g.maxvalue < g.cstats[i].Max {
			g.maxvalue = g.cstats[i].Max
		}
		//fmt.Printf("%d, %d, %#v\n", min, max, v)
	}

	for i := range g.cstats {
		g.cstats[i].Avg = g.row - g.cstats[i].Avg*g.row/g.maxvalue
		g.cstats[i].Min = g.row - g.cstats[i].Min*g.row/g.maxvalue
		g.cstats[i].Max = g.row - g.cstats[i].Max*g.row/g.maxvalue
	}
}
func (g *Graph) Render() string {
	g.calcColStats()
	g.clear()
	var i, j int64
	for i = 0; i < g.col; i++ {
		stat := g.cstats[i]
		for j = 0; j < stat.Max; j++ {
			g.buffer[j][i] = FREE
		}
		for ; j < stat.Avg; j++ {
			g.buffer[j][i] = LIGHT
		}
		for ; j < stat.Min; j++ {
			g.buffer[j][i] = MEDIUM
		}
		for ; j < g.row; j++ {
			g.buffer[j][i] = HARD
		}
	}
	return string(g.rawbuffer)
}

func (g *Graph) MinMaxAvg() (min int64, max int64, avg int64) {
	return ptsMinMaxAvg(g.points)
}
