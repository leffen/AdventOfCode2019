package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGrid(t *testing.T) {
	var tests = []struct {
		in  string
		out int
	}{
		{"R8,U5,L5,D3\nU7,R6,D4,L4", 6},
		{"R75,D30,R83,U83,L12,D49,R71,U7,L72\nU62,R66,U55,R34,D71,R55,D58,R83", 159},
		{"R98,U47,R26,D63,R33,U87,L62,D20,R33,U53,R51\nU98,R91,D20,R16,D67,R40,U7,R15,U6,R7", 135},
	}

	for _, tt := range tests {
		t.Run(tt.in, func(tx *testing.T) {
			p := program{}
			p.run(tt.in)
			dist := p.shortestPath1()
			if tt.out != dist {
				//	p.grd.show()
			}
			assert.Equal(t, tt.out, dist)
		})
	}
}

func TestShow(t *testing.T) {
	input := "R8,U5,L5,D3\nU7,R6,D4,L4"
	res := 6

	p := program{}
	p.run(input)
	dist := p.shortestPath1()
	p.grd.show()
	assert.Equal(t, res, dist)
	assert.True(t, false)
}
