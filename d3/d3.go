package main

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"strconv"
	"strings"
)

func main() {

}

type cell struct {
	x   int
	y   int
	num int
	cnt int
}

type grid struct {
	cells map[string]*cell
}

func (g *grid) setXy(x, y, num int) {
	if g.cells == nil {
		g.cells = map[string]*cell{}
	}

	key := fmt.Sprintf("%d,%d", x, y)

	v, ok := g.cells[key]
	if ok {
		v.cnt++
		v.num = 0
		return
	}

	v = &cell{x: x, y: y, cnt: 1, num: num + 1}
	g.cells[key] = v
}

func (g *grid) getCell(x, y int) *cell {
	key := fmt.Sprintf("%d,%d", x, y)

	cell, ok := g.cells[key]
	if !ok {
		return nil
	}
	return cell
}

func (g *grid) show() {
	maxx := -9999
	minx := 9999

	maxy := -9999
	miny := 9999

	for _, c := range g.cells {
		if c.x < minx {
			minx = c.x
		}
		if c.x > maxx {
			maxx = c.x
		}
		if c.y < miny {
			miny = c.y
		}
		if c.y > maxy {
			maxy = c.y
		}
	}
	fmt.Printf("miny: %d maxy: %d minx:%d maxx:%d\n", miny, maxy, minx, maxx)
	for y := maxy + 1; y > miny-1; y-- {
		for x := minx - 1; x < maxx+2; x++ {
			cell := g.getCell(x, y)
			pos := "."
			if cell != nil {
				pos = fmt.Sprintf("%d", cell.num)
			}
			fmt.Print(pos)
		}
		fmt.Println("")
	}
	fmt.Println("")

}

type program struct {
	grd  grid
	posX int
	posY int
}

func (p *program) walk(dist int, walker func(i int)) {
	for i := 0; i < dist; i++ {
		walker(i)
	}
}

func (p *program) performWalk(nr int, steps []string) {
	p.posX = 0
	p.posY = 0
	for _, step := range steps {
		dist, err := strconv.Atoi(step[1:])
		if err != nil {
			logrus.Fatal(err)
		}

		switch step[0] {
		case 'R':
			p.walk(dist, func(i int) {
				p.posX++
				p.grd.setXy(p.posX, p.posY, nr)
			})
		case 'L':
			p.walk(dist, func(i int) {
				p.posX--
				p.grd.setXy(p.posX, p.posY, nr)
			})
		case 'D':
			p.walk(dist, func(i int) {
				p.posY--
				p.grd.setXy(p.posX, p.posY, nr)
			})
		case 'U':
			p.walk(dist, func(i int) {
				p.posY++
				p.grd.setXy(p.posX, p.posY, nr)
			})
		}
	}
}

func (p *program) run(input string) int {
	p.grd = grid{}
	walks := strings.Split(input, "\n")
	for num, w := range walks {
		steps := strings.Split(w, ",")
		p.performWalk(num, steps)
	}
	return 0
}
