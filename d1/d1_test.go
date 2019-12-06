package main

import (
	"github.com/stretchr/testify/assert"
	"strconv"
	"testing"
)

func TestImport(t *testing.T) {
	lines, err := importFile("input.txt")
	assert.Nil(t, err)
	assert.NotNil(t, lines)
	assert.Equal(t, 100, len(lines))
}

func TestFuelForMass(t *testing.T) {
	var tests = []struct {
		in  int
		out int
	}{
		{12, 2},
		{14, 2},
		{1969, 654},
		{100756, 33583},
	}

	for _, tt := range tests {
		t.Run(strconv.Itoa(tt.in), func(tx *testing.T) {
			assert.Equal(tx, tt.out, fuelForMass(tt.in))
		})
	}
}

func TestLinesConv(t *testing.T) {
	lines := []string{"1", "2", "3"}
	il := linesToIntA(lines)
	assert.Equal(t, 3, len(il))
}

func TestInitalFuel(t *testing.T) {
	lines, _ := importFile("input.txt")
	intA := linesToIntA(lines)
	fuel := calcFuelForMass(intA)
	assert.Equal(t, 3372695, fuel)
}

func TestCalcFuelReq(t *testing.T) {
	t1 := 100756
	t2 := calcFuelForFuelReq(t1)
	assert.Equal(t, 50346, t2)

	t1 = 1969
	t2 = calcFuelForFuelReq(t1)
	assert.Equal(t, 966, t2)

}
