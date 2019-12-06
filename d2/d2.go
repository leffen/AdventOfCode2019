package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/sirupsen/logrus"
)

func main() {
	input := "1,0,0,3,1,1,2,3,1,3,4,3,1,5,0,3,2,10,1,19,1,5,19,23,1,23,5,27,2,27,10,31,1,5,31,35,2,35,6,39,1,6,39,43,2,13,43,47,2,9,47,51,1,6,51,55,1,55,9,59,2,6,59,63,1,5,63,67,2,67,13,71,1,9,71,75,1,75,9,79,2,79,10,83,1,6,83,87,1,5,87,91,1,6,91,95,1,95,13,99,1,10,99,103,2,6,103,107,1,107,5,111,1,111,13,115,1,115,13,119,1,13,119,123,2,123,13,127,1,127,6,131,1,131,9,135,1,5,135,139,2,139,6,143,2,6,143,147,1,5,147,151,1,151,2,155,1,9,155,0,99,2,14,0,0"

	p := program{}
	p.prepItems(input)
	p.items[1] = 12
	p.items[2] = 2
	res := p.run()
	fmt.Printf("Res: %d\n", res)
}

type program struct {
	items []int
	curr  int
}

func (p *program) prepItems(input string) {
	p.items = strToIntA(input)
}

func (p *program) run() int {
	p.curr = 0

	for p.curr < len(p.items) {
		//		fmt.Printf("%d %v\n", p.curr, p.items)
		switch p.items[p.curr] {
		case 1:
			p.execAdd()
		case 2:
			p.execMultiply()
		case 99:
			return p.items[0]
		}

	}
	logrus.Fatal("Unexpected end of program")
	return -1
}

func (p *program) execAdd() {
	intA := p.val(p.curr + 1)
	intB := p.val(p.curr + 2)
	pos := p.items[p.curr+3]
	p.items[pos] = intA + intB
	p.curr += 4
}

func (p *program) execMultiply() {
	intA := p.val(p.curr + 1)
	intB := p.val(p.curr + 2)
	pos := p.items[p.curr+3]
	p.items[pos] = intA * intB
	p.curr += 4
}

func (p *program) val(pos int) int {
	return p.items[p.items[pos]]
}

func strToIntA(data string) []int {
	rc := []int{}
	items := strings.Split(data, ",")
	for _, i := range items {
		val, err := strconv.Atoi(i)
		if err != nil {
			logrus.Fatal(err)
		}
		rc = append(rc, val)
	}

	return rc
}
