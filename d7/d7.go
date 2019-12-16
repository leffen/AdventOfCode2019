package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/sirupsen/logrus"
)

func part1() {
	input := "3,8,1001,8,10,8,105,1,0,0,21,38,55,64,81,106,187,268,349,430,99999,3,9,101,2,9,9,1002,9,2,9,101,5,9,9,4,9,99,3,9,102,2,9,9,101,3,9,9,1002,9,4,9,4,9,99,3,9,102,2,9,9,4,9,99,3,9,1002,9,5,9,1001,9,4,9,102,4,9,9,4,9,99,3,9,102,2,9,9,1001,9,5,9,102,3,9,9,1001,9,4,9,102,5,9,9,4,9,99,3,9,1002,9,2,9,4,9,3,9,101,2,9,9,4,9,3,9,1002,9,2,9,4,9,3,9,1001,9,2,9,4,9,3,9,1001,9,2,9,4,9,3,9,101,1,9,9,4,9,3,9,1001,9,1,9,4,9,3,9,1001,9,2,9,4,9,3,9,101,1,9,9,4,9,3,9,1001,9,1,9,4,9,99,3,9,1002,9,2,9,4,9,3,9,101,2,9,9,4,9,3,9,1001,9,1,9,4,9,3,9,101,1,9,9,4,9,3,9,101,2,9,9,4,9,3,9,101,2,9,9,4,9,3,9,1001,9,1,9,4,9,3,9,101,1,9,9,4,9,3,9,102,2,9,9,4,9,3,9,101,2,9,9,4,9,99,3,9,1002,9,2,9,4,9,3,9,101,2,9,9,4,9,3,9,102,2,9,9,4,9,3,9,101,2,9,9,4,9,3,9,1001,9,2,9,4,9,3,9,1002,9,2,9,4,9,3,9,1002,9,2,9,4,9,3,9,101,2,9,9,4,9,3,9,1001,9,2,9,4,9,3,9,101,1,9,9,4,9,99,3,9,102,2,9,9,4,9,3,9,1001,9,2,9,4,9,3,9,1002,9,2,9,4,9,3,9,102,2,9,9,4,9,3,9,102,2,9,9,4,9,3,9,101,2,9,9,4,9,3,9,101,1,9,9,4,9,3,9,101,1,9,9,4,9,3,9,1001,9,1,9,4,9,3,9,102,2,9,9,4,9,99,3,9,101,1,9,9,4,9,3,9,1002,9,2,9,4,9,3,9,102,2,9,9,4,9,3,9,1002,9,2,9,4,9,3,9,101,1,9,9,4,9,3,9,102,2,9,9,4,9,3,9,1002,9,2,9,4,9,3,9,1002,9,2,9,4,9,3,9,101,1,9,9,4,9,3,9,102,2,9,9,4,9,99"
	p := program{}
	p.prepItems(input)

	currSeq := []int{}
	phase := 0
	max := 0

	initSeq := 0
	maxSeq := 4
	for i := initSeq; i <= maxSeq; i++ {
		for j := initSeq; j <= maxSeq; j++ {
			for k := initSeq; k <= maxSeq; k++ {
				for l := initSeq; l <= maxSeq; l++ {
					for m := initSeq; m <= maxSeq; m++ {
						if i == j || i == k || i == l || i == m || j == k || j == l || j == m || k == l || k == m || l == m {
							continue
						}
						phase = 0
						seq := []int{i, j, k, l, m}
						for _, s := range seq {
							inpVals := []int{s, phase}
							p.run(inpVals)
							phase = p.lastOutput()
						}
						if phase > max {
							max = phase
							currSeq = seq
						}
					}
				}
			}
		}
	}

	fmt.Printf("Max=%d Seq:%v\n", max, currSeq)
	//logrus.Fatal("Unable to find result")
}

func part2() {
	input := "3,8,1001,8,10,8,105,1,0,0,21,38,55,64,81,106,187,268,349,430,99999,3,9,101,2,9,9,1002,9,2,9,101,5,9,9,4,9,99,3,9,102,2,9,9,101,3,9,9,1002,9,4,9,4,9,99,3,9,102,2,9,9,4,9,99,3,9,1002,9,5,9,1001,9,4,9,102,4,9,9,4,9,99,3,9,102,2,9,9,1001,9,5,9,102,3,9,9,1001,9,4,9,102,5,9,9,4,9,99,3,9,1002,9,2,9,4,9,3,9,101,2,9,9,4,9,3,9,1002,9,2,9,4,9,3,9,1001,9,2,9,4,9,3,9,1001,9,2,9,4,9,3,9,101,1,9,9,4,9,3,9,1001,9,1,9,4,9,3,9,1001,9,2,9,4,9,3,9,101,1,9,9,4,9,3,9,1001,9,1,9,4,9,99,3,9,1002,9,2,9,4,9,3,9,101,2,9,9,4,9,3,9,1001,9,1,9,4,9,3,9,101,1,9,9,4,9,3,9,101,2,9,9,4,9,3,9,101,2,9,9,4,9,3,9,1001,9,1,9,4,9,3,9,101,1,9,9,4,9,3,9,102,2,9,9,4,9,3,9,101,2,9,9,4,9,99,3,9,1002,9,2,9,4,9,3,9,101,2,9,9,4,9,3,9,102,2,9,9,4,9,3,9,101,2,9,9,4,9,3,9,1001,9,2,9,4,9,3,9,1002,9,2,9,4,9,3,9,1002,9,2,9,4,9,3,9,101,2,9,9,4,9,3,9,1001,9,2,9,4,9,3,9,101,1,9,9,4,9,99,3,9,102,2,9,9,4,9,3,9,1001,9,2,9,4,9,3,9,1002,9,2,9,4,9,3,9,102,2,9,9,4,9,3,9,102,2,9,9,4,9,3,9,101,2,9,9,4,9,3,9,101,1,9,9,4,9,3,9,101,1,9,9,4,9,3,9,1001,9,1,9,4,9,3,9,102,2,9,9,4,9,99,3,9,101,1,9,9,4,9,3,9,1002,9,2,9,4,9,3,9,102,2,9,9,4,9,3,9,1002,9,2,9,4,9,3,9,101,1,9,9,4,9,3,9,102,2,9,9,4,9,3,9,1002,9,2,9,4,9,3,9,1002,9,2,9,4,9,3,9,101,1,9,9,4,9,3,9,102,2,9,9,4,9,99"
	p := program{}
	p.prepItems(input)

	currSeq := []int{}
	phase := 0
	max := 0

	initSeq := 5
	maxSeq := 9
	for i := initSeq; i <= maxSeq; i++ {
		for j := initSeq; j <= maxSeq; j++ {
			for k := initSeq; k <= maxSeq; k++ {
				for l := initSeq; l <= maxSeq; l++ {
					for m := initSeq; m <= maxSeq; m++ {
						if i == j || i == k || i == l || i == m || j == k || j == l || j == m || k == l || k == m || l == m {
							continue
						}
						seq := []int{i, j, k, l, m}
						for _, s := range seq {
							inpVals := []int{s, phase}
							p.run(inpVals)
							phase = p.lastOutput()
						}
						if phase > max {
							max = phase
							currSeq = seq
						}
					}
				}
			}
		}
	}

	fmt.Printf("Max=%d Seq:%v\n", max, currSeq)
	//logrus.Fatal("Unable to find result")
}

func main() {
	part1()
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
	if p.curr+1 >= len(p.items) {
		logrus.Fatalf("Unable to get index %d from inputs with length %d", p.curr+1, len(p.items))
	}
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
