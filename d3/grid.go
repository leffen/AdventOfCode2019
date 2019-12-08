package main

import "fmt"

type grid struct {
	cells map[string]*cell
}

func (g *grid) setXy(x, y, num int, mrk string, step int) {
	if g.cells == nil {
		g.cells = map[string]*cell{}
	}

	key := fmt.Sprintf("%d,%d", x, y)

	v, ok := g.cells[key]
	if ok {
		if v.num != num {
			fmt.Printf("XCROSS (%d,%d)  N:%d M:%s -> %s\n", x, y, num, mrk, v)
			v.cnt++
			v.mrk = "x"
			if v.xstep == nil {
				v.xstep = map[int]int{}
			}
			v.xstep[v.num] = v.step
			v.xstep[num] = step
		}
		return
	}

	v = &cell{x: x, y: y, cnt: 1, num: num, mrk: mrk, step: step}
	g.cells[key] = v
}

func (g *grid) count() int {
	return len(g.cells)
}

func (g *grid) getCell(x, y int) *cell {
	key := fmt.Sprintf("%d,%d", x, y)

	cell, ok := g.cells[key]
	if !ok {
		return nil
	}
	return cell
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
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
	for y := maxy + 1; y > miny-2; y-- {
		for x := minx - 1; x < maxx+2; x++ {
			cell := g.getCell(x, y)
			pos := "."
			if cell != nil {
				pos = cell.mrk
			}
			fmt.Print(pos)
		}
		fmt.Println("")
	}
	fmt.Println("")
}
