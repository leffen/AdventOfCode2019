package main

import (
	"testing"

	"gotest.tools/assert"
)

func TestInput(t *testing.T) {
	input := "1,9,10,3,2,3,11,0,99,30,40,50"

	items := strToIntA(input)
	assert.Equal(t, 12, len(items))
	assert.Equal(t, 1, items[0])
	assert.Equal(t, 50, items[11])

}


func TestRun1(t *testing.T) {
	input := "1,9,10,3,2,3,11,0,99,30,40,50"
	p := program{}
	p.prepItems(input)
	res := p.run()
	assert.Equal(t,3500,res)
}

func TestRun2(t *testing.T) {
	var tests = []struct {
		in  string
		out []int
	}{
		{"1,0,0,0,99", []int{2,0,0,0,99}},
		{"2,3,0,3,99",[]int{2,3,0,6,99}},
		{"2,4,4,5,99,0",[]int{2,4,4,5,99,9801}},
		{"1,1,1,4,99,5,6,0,99",[]int{30,1,1,4,2,5,6,0,99}},
	}

	for _, tt := range tests{
		t.Run(tt.in,func(tx *testing.T){
			p := program{}
			p.prepItems(tt.in)
			p.run()
			for idx,v:= range tt.out {
				assert.Equal(t,v,p.items[idx])
			}
		})

	}

}
