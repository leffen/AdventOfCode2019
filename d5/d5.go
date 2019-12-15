package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/sirupsen/logrus"
)

func part1() {
	input := "3,225,1,225,6,6,1100,1,238,225,104,0,1101,48,82,225,102,59,84,224,1001,224,-944,224,4,224,102,8,223,223,101,6,224,224,1,223,224,223,1101,92,58,224,101,-150,224,224,4,224,102,8,223,223,1001,224,3,224,1,224,223,223,1102,10,89,224,101,-890,224,224,4,224,1002,223,8,223,1001,224,5,224,1,224,223,223,1101,29,16,225,101,23,110,224,1001,224,-95,224,4,224,102,8,223,223,1001,224,3,224,1,223,224,223,1102,75,72,225,1102,51,8,225,1102,26,16,225,1102,8,49,225,1001,122,64,224,1001,224,-113,224,4,224,102,8,223,223,1001,224,3,224,1,224,223,223,1102,55,72,225,1002,174,28,224,101,-896,224,224,4,224,1002,223,8,223,101,4,224,224,1,224,223,223,1102,57,32,225,2,113,117,224,101,-1326,224,224,4,224,102,8,223,223,101,5,224,224,1,223,224,223,1,148,13,224,101,-120,224,224,4,224,1002,223,8,223,101,7,224,224,1,223,224,223,4,223,99,0,0,0,677,0,0,0,0,0,0,0,0,0,0,0,1105,0,99999,1105,227,247,1105,1,99999,1005,227,99999,1005,0,256,1105,1,99999,1106,227,99999,1106,0,265,1105,1,99999,1006,0,99999,1006,227,274,1105,1,99999,1105,1,280,1105,1,99999,1,225,225,225,1101,294,0,0,105,1,0,1105,1,99999,1106,0,300,1105,1,99999,1,225,225,225,1101,314,0,0,106,0,0,1105,1,99999,8,677,226,224,102,2,223,223,1006,224,329,101,1,223,223,107,677,677,224,1002,223,2,223,1006,224,344,101,1,223,223,8,226,677,224,102,2,223,223,1006,224,359,101,1,223,223,107,226,226,224,102,2,223,223,1005,224,374,1001,223,1,223,1108,677,226,224,1002,223,2,223,1006,224,389,101,1,223,223,107,677,226,224,102,2,223,223,1006,224,404,1001,223,1,223,1107,226,677,224,1002,223,2,223,1006,224,419,1001,223,1,223,108,677,677,224,102,2,223,223,1005,224,434,1001,223,1,223,1008,677,226,224,1002,223,2,223,1006,224,449,1001,223,1,223,7,226,677,224,1002,223,2,223,1006,224,464,1001,223,1,223,1007,677,677,224,102,2,223,223,1005,224,479,1001,223,1,223,1007,226,226,224,1002,223,2,223,1005,224,494,1001,223,1,223,108,226,226,224,1002,223,2,223,1005,224,509,1001,223,1,223,1007,226,677,224,1002,223,2,223,1006,224,524,101,1,223,223,1107,677,677,224,102,2,223,223,1005,224,539,101,1,223,223,1107,677,226,224,102,2,223,223,1005,224,554,1001,223,1,223,108,677,226,224,1002,223,2,223,1006,224,569,1001,223,1,223,1108,226,677,224,1002,223,2,223,1006,224,584,101,1,223,223,8,677,677,224,1002,223,2,223,1006,224,599,1001,223,1,223,1008,226,226,224,102,2,223,223,1006,224,614,101,1,223,223,7,677,677,224,1002,223,2,223,1006,224,629,101,1,223,223,1008,677,677,224,102,2,223,223,1005,224,644,101,1,223,223,7,677,226,224,1002,223,2,223,1005,224,659,101,1,223,223,1108,226,226,224,102,2,223,223,1006,224,674,1001,223,1,223,4,223,99,226"
	inpVals := []int{1}
	p := program{}
	p.prepItems(input)
	res := p.run(inpVals)
	fmt.Printf("RES=%d Outputs:%v\n", res, p.outputs)
	//logrus.Fatal("Unable to find result")
}

func part2() {
	input := "3,225,1,225,6,6,1100,1,238,225,104,0,1101,48,82,225,102,59,84,224,1001,224,-944,224,4,224,102,8,223,223,101,6,224,224,1,223,224,223,1101,92,58,224,101,-150,224,224,4,224,102,8,223,223,1001,224,3,224,1,224,223,223,1102,10,89,224,101,-890,224,224,4,224,1002,223,8,223,1001,224,5,224,1,224,223,223,1101,29,16,225,101,23,110,224,1001,224,-95,224,4,224,102,8,223,223,1001,224,3,224,1,223,224,223,1102,75,72,225,1102,51,8,225,1102,26,16,225,1102,8,49,225,1001,122,64,224,1001,224,-113,224,4,224,102,8,223,223,1001,224,3,224,1,224,223,223,1102,55,72,225,1002,174,28,224,101,-896,224,224,4,224,1002,223,8,223,101,4,224,224,1,224,223,223,1102,57,32,225,2,113,117,224,101,-1326,224,224,4,224,102,8,223,223,101,5,224,224,1,223,224,223,1,148,13,224,101,-120,224,224,4,224,1002,223,8,223,101,7,224,224,1,223,224,223,4,223,99,0,0,0,677,0,0,0,0,0,0,0,0,0,0,0,1105,0,99999,1105,227,247,1105,1,99999,1005,227,99999,1005,0,256,1105,1,99999,1106,227,99999,1106,0,265,1105,1,99999,1006,0,99999,1006,227,274,1105,1,99999,1105,1,280,1105,1,99999,1,225,225,225,1101,294,0,0,105,1,0,1105,1,99999,1106,0,300,1105,1,99999,1,225,225,225,1101,314,0,0,106,0,0,1105,1,99999,8,677,226,224,102,2,223,223,1006,224,329,101,1,223,223,107,677,677,224,1002,223,2,223,1006,224,344,101,1,223,223,8,226,677,224,102,2,223,223,1006,224,359,101,1,223,223,107,226,226,224,102,2,223,223,1005,224,374,1001,223,1,223,1108,677,226,224,1002,223,2,223,1006,224,389,101,1,223,223,107,677,226,224,102,2,223,223,1006,224,404,1001,223,1,223,1107,226,677,224,1002,223,2,223,1006,224,419,1001,223,1,223,108,677,677,224,102,2,223,223,1005,224,434,1001,223,1,223,1008,677,226,224,1002,223,2,223,1006,224,449,1001,223,1,223,7,226,677,224,1002,223,2,223,1006,224,464,1001,223,1,223,1007,677,677,224,102,2,223,223,1005,224,479,1001,223,1,223,1007,226,226,224,1002,223,2,223,1005,224,494,1001,223,1,223,108,226,226,224,1002,223,2,223,1005,224,509,1001,223,1,223,1007,226,677,224,1002,223,2,223,1006,224,524,101,1,223,223,1107,677,677,224,102,2,223,223,1005,224,539,101,1,223,223,1107,677,226,224,102,2,223,223,1005,224,554,1001,223,1,223,108,677,226,224,1002,223,2,223,1006,224,569,1001,223,1,223,1108,226,677,224,1002,223,2,223,1006,224,584,101,1,223,223,8,677,677,224,1002,223,2,223,1006,224,599,1001,223,1,223,1008,226,226,224,102,2,223,223,1006,224,614,101,1,223,223,7,677,677,224,1002,223,2,223,1006,224,629,101,1,223,223,1008,677,677,224,102,2,223,223,1005,224,644,101,1,223,223,7,677,226,224,1002,223,2,223,1005,224,659,101,1,223,223,1108,226,226,224,102,2,223,223,1006,224,674,1001,223,1,223,4,223,99,226"
	inpVals := []int{1}
	p := program{}
	p.prepItems(input)
	res := p.run(inpVals)
	fmt.Printf("RES=%d Outputs:%v\n", res, p.outputs)
	//logrus.Fatal("Unable to find result")
}

func main() {
	part2()
}

type program struct {
	items   []int
	curr    int
	outputs []int
	input   string
}

type opCode struct {
	op         int
	param1mode int
	param2mode int
	param3mode int
}

func (o *opCode) String() string {
	return fmt.Sprintf(" opc:%d m1:%d m2:%d m3:%d", o.op, o.param1mode, o.param2mode, o.param3mode)
}

func (o *opCode) Code() string {
	switch o.op {
	case 1:
		return "ADD"
	case 2:
		return "MUL"
	case 3:
		return "STO"
	case 4:
		return "GET"
	case 5:
		return "JUMP TRUE"
	case 6:
		return "JUMP FALSE"
	case 7:
		return "LESS THAN"
	case 8:
		return "EQUAL"
	case 99:
		return "STOP"
	}
	logrus.Fatal("Unknown opcode")
	return "UNKNOWN"
}

func (o *opCode) Modes() string {
	return fmt.Sprintf("m1:%d m2:%d m3:%d", o.param1mode, o.param2mode, o.param3mode)
}

func (p *program) prepItems(input string) {
	p.input = input
	p.items = strToIntA(input)
}

func (p *program) assignItems() {
	p.items = strToIntA(p.input)
}

func (p *program) showItems() {
	fmt.Printf("Items: %v\n", p.items)
	for idx, i := range p.items {
		fmt.Printf(" %02d %d\n", idx, i)
	}
}

func (p *program) run(inputs []int) int {
	logrus.Debugf("RUNNING with %v", inputs)
	p.outputs = []int{}
	p.assignItems()
	p.curr = 0

	for p.curr < len(p.items) {
		//		fmt.Printf("%d %v\n", p.curr, p.items)
		cmd := parseOpcode(p.items[p.curr])
		logrus.Debugf("  %s[%d] modes:%v", cmd.Code(), p.curr, cmd.Modes())
		switch cmd.op {
		case 1:
			p.execAdd(cmd)
		case 2:
			p.execMultiply(cmd)
		case 3:
			p.store(inputs[0])
			inputs = inputs[1:]
		case 4:
			v := p.getVal(cmd)
			p.outputs = append(p.outputs, v)
		//fmt.Printf("Output: %d\n", v)
		case 5:
			p.jumpIfTrue(cmd)
		case 6:
			p.jumpIfFalse(cmd)
		case 7:
			p.lessThan(cmd)
		case 8:
			p.equal(cmd)

		case 99:
			logrus.Debugf("    Exit with %d outputs: %v", p.items[0], p.outputs)
			return p.items[0]
		default:
			logrus.Fatalf("Invalid opcode %v", cmd)

		}

	}
	logrus.Fatal("Unexpected end of program")
	return -1
}

// Op 1
func (p *program) execAdd(op *opCode) {
	intA := p.valByMode(p.curr+1, op.param1mode)
	intB := p.valByMode(p.curr+2, op.param2mode)
	v := intA + intB
	pos := p.items[p.curr+3]
	if op.param3mode == 1 {
		pos = p.curr + 3
	}

	p.items[pos] = v

	logrus.Debugf("      i1: %d i2: %d items[%d]=%d items: %v ", intA, intB, pos, v, p.itemsForDisplay())

	p.curr += 4
}

// Op 2
func (p *program) execMultiply(op *opCode) {
	intA := p.valByMode(p.curr+1, op.param1mode)
	intB := p.valByMode(p.curr+2, op.param2mode)
	v := intA * intB
	pos := p.items[p.curr+3]
	if op.param3mode == 1 {
		pos = p.curr + 3
	}
	p.items[pos] = v

	logrus.Debugf("      i1: %d * %d = %v items[%d]=%d  items: %v op:%v", intA, v, pos, v, intB, p.itemsForDisplay(), op)

	p.curr += 4
}

// Opcode 3
func (p *program) store(inp int) {
	pos := p.items[p.curr+1]
	p.items[pos] = inp
	logrus.Debugf("      save %d in %d items[%d]=%d items: %v ", inp, pos, pos, inp, p.itemsForDisplay())

	p.curr += 2
}

// Opcode 4
func (p *program) getVal(op *opCode) int {
	pos := p.items[p.curr+1]
	if op.param1mode == 1 {
		pos = p.curr + 1
	}
	v := p.items[pos]
	logrus.Debugf("      Get val in pos: %d returns: %d items: %v", pos, v, p.itemsForDisplay())
	p.curr += 2
	return v
}

// Opcode 5
func (p *program) jumpIfTrue(op *opCode) int {
	v := p.valByMode(p.curr+1, op.param1mode)
	pos := p.valByMode(p.curr+2, op.param2mode)

	logrus.Debugf("      Jump if %v != 0. jump_to:%d items: %v", v, pos, p.itemsForDisplay())

	if v != 0 {
		p.curr = pos
	} else {
		p.curr += 3
	}
	return p.curr
}

// Opcode 6
func (p *program) jumpIfFalse(op *opCode) int {
	v := p.valByMode(p.curr+1, op.param1mode)
	pos := p.valByMode(p.curr+2, op.param2mode)

	logrus.Debugf("      Jump to %d if %d == 0 items: %v", pos, v, p.itemsForDisplay())

	if v == 0 {
		p.curr = pos
	} else {
		p.curr += 3
	}
	return p.curr
}

// Opcode 7
func (p *program) lessThan(op *opCode) int {
	v1 := p.valByMode(p.curr+1, op.param1mode)
	v2 := p.valByMode(p.curr+2, op.param2mode)
	pos := p.posByMode(p.curr+3, op.param3mode)
	v := int(0)
	if v1 < v2 {
		v = 1
	}
	p.items[pos] = v
	logrus.Debugf("        Store %d in %d (1 if %d < %d else 0) items[%d]=%d items: %v ", v, pos, v1, v2, pos, p.items[pos], p.itemsForDisplay())
	p.curr += 4
	return p.curr
}

// Opcode 8
func (p *program) equal(op *opCode) int {
	v1 := p.valByMode(p.curr+1, op.param1mode)
	v2 := p.valByMode(p.curr+2, op.param2mode)
	pos := p.posByMode(p.curr+3, op.param3mode)
	v := int(0)
	if v1 == v2 {
		v = 1
	}
	logrus.Debugf("      Store %d in %d (1 if %d== %d else 0) items: %v", v, pos, v1, v2, p.itemsForDisplay())
	p.items[pos] = v
	p.curr += 4
	return p.curr
}

func parseOpcode(code int) *opCode {
	cmd := fmt.Sprintf("%05d", code)
	op, err := strconv.Atoi(cmd[3:5])
	if err != nil {
		logrus.Fatal(err)
	}
	//	fmt.Printf("OP %d cmd:%s part: %s\n", op, cmd, string(cmd[3:5]))
	return &opCode{op: op, param1mode: posToInt(cmd, 2), param2mode: posToInt(cmd, 1), param3mode: posToInt(cmd, 0)}
}

func posToInt(cmd string, pos int) int {
	i, err := strconv.Atoi(string(cmd[pos]))
	if err != nil {
		logrus.Fatal(err)
	}
	return i
}

func (p *program) lastOutput() int {
	return p.outputs[len(p.outputs)-1]
}

func (p *program) itemsForDisplay() []int {
	maxL := len(p.items) - 1
	if maxL < 20 {
		return p.items
	}
	toAddr := p.curr + 20
	if toAddr > maxL {
		toAddr = maxL
	}

	return p.items[p.curr:toAddr]
}
func (p *program) val(pos int) int {
	return p.items[p.items[pos]]
}

func (p *program) valByMode(pos, mode int) int {
	// Imidiate mode
	if mode == 1 {
		return p.items[pos]
	}
	return p.items[p.items[pos]]
}

func (p *program) posByMode(pos, mode int) int {
	if mode == 0 {
		return p.items[pos]
	}
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
