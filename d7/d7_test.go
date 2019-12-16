package main

import (
	"fmt"
	"testing"

	"github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"
)

func Init() {
	logrus.SetLevel(logrus.WarnLevel)
}

func TestInput(t *testing.T) {
	input := "1,9,10,3,2,3,11,0,99,30,40,50"

	items := strToIntA(input)
	assert.Equal(t, 12, len(items))
	assert.Equal(t, 1, items[0])
	assert.Equal(t, 50, items[11])

}

func TestRun1(t *testing.T) {
	logrus.SetLevel(logrus.DebugLevel)
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

func TestOp5678(t *testing.T) {
	//	logrus.SetLevel(logrus.DebugLevel)
	input := "3,9,8,9,10,9,4,9,99,-1,8"
	p := program{}
	p.prepItems(input)
	res := p.run([]int{77})
	fmt.Printf("Outputs: %v\n", p.outputs)
	assert.Equal(t, 3, res)
	assert.Equal(t, 0, p.lastOutput())

	res = p.run([]int{8})
	assert.Equal(t, 1, p.lastOutput())

	var tests = []struct {
		name    string
		in      int
		prg     string
		lastOut int
	}{
		{"t1", 99, "3,9,8,9,10,9,4,9,99,-1,8", 0},
		{"t1", 8, "3,9,8,9,10,9,4,9,99,-1,8", 1},

		{"t2", 8, "3,9,7,9,10,9,4,9,99,-1,8", 0},
		{"t2", 7, "3,9,7,9,10,9,4,9,99,-1,8", 1},
		{"t2", 77, "3,9,7,9,10,9,4,9,99,-1,8", 0},

		{"t3", 0, "3,3,1108,-1,8,3,4,3,99", 0},
		{"t3", 7, "3,3,1108,-1,8,3,4,3,99", 0},
		{"t3", 8, "3,3,1108,-1,8,3,4,3,99", 1},
		{"t3", 999, "3,3,1108,-1,8,3,4,3,99", 0},

		{"t4", 0, "3,3,1107,-1,8,3,4,3,99", 1},
		{"t4", 7, "3,3,1107,-1,8,3,4,3,99", 1},
		{"t4", 8, "3,3,1107,-1,8,3,4,3,99", 0},
		{"t4", 999, "3,3,1107,-1,8,3,4,3,99", 0},

		{"t5", 0, "3,12,6,12,15,1,13,14,13,4,13,99,-1,0,1,9", 0},
		{"t5", 99, "3,12,6,12,15,1,13,14,13,4,13,99,-1,0,1,9", 1},

		{"t6", 0, "3,3,1105,-1,9,1101,0,0,12,4,12,99,1", 0},
		{"t6", 99, "3,3,1105,-1,9,1101,0,0,12,4,12,99,1", 1},
	}
	for _, tt := range tests {
		t.Run(fmt.Sprintf("OP %05d", tt.in), func(tx *testing.T) {
			p := program{}
			p.prepItems(tt.prg)
			p.run([]int{tt.in})
			//fmt.Printf("Outputs: %v\n", p.outputs)
			assert.Equal(t, tt.lastOut, p.lastOutput())
		})
	}

}

func TestJump(t *testing.T) {
	//logrus.SetLevel(logrus.DebugLevel)
	input := "3,12,6,12,15,1,13,14,13,4,13,99,-1,0,1,9"
	p := program{}
	p.prepItems(input)
	p.showItems()
	res := p.run([]int{0})
	fmt.Printf("Outputs: %v\n", p.outputs)
	assert.Equal(t, 3, res)
	assert.Equal(t, 0, p.lastOutput())

	res = p.run([]int{8})
	assert.Equal(t, 1, p.lastOutput())
}

func TestJumpX(t *testing.T) {
	// logrus.SetLevel(logrus.DebugLevel)
	input := "3,3,1105,-1,9,1101,0,0,12,4,12,99,1"
	p := program{}
	p.prepItems(input)
	res := p.run([]int{0})
	// fmt.Printf("Outputs: %v\n", p.outputs)
	assert.Equal(t, 3, res)
	assert.Equal(t, 0, p.lastOutput())

	p.showItems()
	res = p.run([]int{1})
	fmt.Printf("Outputs: %v\n", p.outputs)
	assert.Equal(t, 1, p.lastOutput())
}

func TestJump2(t *testing.T) {

	input := "3,21,1008,21,8,20,1005,20,22,107,8,21,20,1006,20,31,1106,0,36,98,0,0,1002,21,125,20,4,20,1105,1,46,104,999,1105,1,46,1101,1000,1,20,4,20,1105,1,46,98,99"
	p := program{}
	p.prepItems(input)
	//	p.showItems()
	res := p.run([]int{0})
	assert.Equal(t, 3, res)
	assert.Equal(t, 999, p.lastOutput())

	res = p.run([]int{8})
	assert.Equal(t, 1000, p.lastOutput())

	res = p.run([]int{789})
	assert.Equal(t, 1001, p.lastOutput())
}
