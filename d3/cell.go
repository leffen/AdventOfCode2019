package main

import "fmt"


type cell struct {
	x   int
	y   int
	num int
	cnt int
	mrk string
	step int
	xstep map[int]int
}

func (c *cell) manhattenDist() int {
	return abs(c.x) + abs(c.y)
}

func (c *cell) String() string {
	return fmt.Sprintf("(%d,%d) N:%d C:%d M:%s", c.x, c.y, c.num, c.cnt, c.mrk)
}
