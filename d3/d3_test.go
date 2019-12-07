package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGrid(t *testing.T) {
	//input := "R75,D30,R83,U83,L12,D49,R71,U7,L72,U62,R66,U55,R34,D71,R55,D58,R83"
	input := "R8,U5,L5,D3\nU7,R6,D4,L4"
	p := &program{}
	p.run(input)
	p.grd.show()
	// Debug
	assert.True(t, false)
}
