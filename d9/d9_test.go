package main

import (
	"fmt"
	"testing"

	"github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"
)

func TestPart1(t *testing.T) {

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
		{2002, opCode{2, 0, 2, 0}},
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

func TestD7Part2(t *testing.T) {
	input := "3,26,1001,26,-4,26,3,27,1002,27,2,27,1,27,26,27,4,27,1001,28,-1,28,1005,28,6,99,0,0,5"
	seq := []int{9, 8, 7, 6, 5}

	max := analyzeSeq(input, seq)
	assert.Equal(t, 139629729, max)
}

func TestRelBase(t *testing.T) {
	input := "109,1,204,-1,1001,100,1,100,1008,100,16,101,1006,101,0,99"
	logrus.SetLevel(logrus.DebugLevel)

	p := program{}
	p.prepItems(input)
	p.run([]int{})
	assert.Equal(t, 109, p.outputs[0])
	assert.Equal(t, 16, len(p.outputs))
}
