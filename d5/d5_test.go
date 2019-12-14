package main

import (
	"fmt"
	"testing"

	"github.com/sirupsen/logrus"
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
	res := p.run([]int{})
	assert.Equal(t, 3500, res)
}

func TestRun2(t *testing.T) {
	var tests = []struct {
		in  string
		out []int
	}{
		{"1,0,0,0,99", []int{2, 0, 0, 0, 99}},
		{"2,3,0,3,99", []int{2, 3, 0, 6, 99}},
		{"2,4,4,5,99,0", []int{2, 4, 4, 5, 99, 9801}},
		{"1,1,1,4,99,5,6,0,99", []int{30, 1, 1, 4, 2, 5, 6, 0, 99}},
	}

	for _, tt := range tests {
		t.Run(tt.in, func(tx *testing.T) {
			p := program{}
			p.prepItems(tt.in)
			p.run([]int{})
			for idx, v := range tt.out {
				assert.Equal(t, v, p.items[idx])
			}
		})

	}
}

func TestInpAndOutput(t *testing.T) {
	logrus.SetLevel(logrus.DebugLevel)
	input := "3,0,4,0,99"
	p := program{}
	p.prepItems(input)
	p.run([]int{1})
	assert.Equal(t, 1, p.outputs[0])
}

func TestOpCode(t *testing.T) {
	var tests = []struct {
		in  int
		out opCode
	}{
		{1002, opCode{2, 0, 1, 0}},
		{11102, opCode{2, 1, 1, 1}},
		{10002, opCode{2, 0, 0, 1}},
		{2, opCode{2, 0, 0, 0}},
	}
	for _, tt := range tests {
		t.Run(fmt.Sprintf("OP %05d", tt.in), func(tx *testing.T) {
			op1 := parseOpcode(tt.in)
			assert.Equal(tx, tt.out.op, op1.op)
			assert.Equal(tx, tt.out.param1mode, op1.param1mode)
			assert.Equal(tx, tt.out.param2mode, op1.param2mode)
			assert.Equal(tx, tt.out.param3mode, op1.param3mode)
		})
	}

}

func TestOpcodeInp(t *testing.T) {
	logrus.SetLevel(logrus.DebugLevel)
	input := "1002,4,3,4,33"
	p := program{}
	p.prepItems(input)
	res := p.run([]int{})
	assert.Equal(t, 1002, res)
}
