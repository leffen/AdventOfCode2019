package main

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestImport(t *testing.T) {
	items, err := importFile("input.txt")
	assert.Nil(t, err)
	assert.Equal(t, 1866, len(items))

	o := newOribitCalc(items)
	fmt.Printf("map: %d", len(o.mp))

	x := o.calcAllOrbits()
	assert.Equal(t, 1, x)
}

func TestOrbitals(t *testing.T) {
	data := `
	COM)B
B)C
C)D
D)E
E)F
B)G
G)H
D)I
E)J
J)K
K)L
	`
	items, err := itemFromText(data)
	assert.Nil(t, err)
	assert.NotNil(t, items)
	o := newOribitCalc(items)
	cnt := o.calcAllOrbits()
	assert.Equal(t, 42, cnt)
}

func TestMinimumHops(t *testing.T) {
	data := `COM)B
	B)C
	C)D
	D)E
	E)F
	B)G
	G)H
	D)I
	E)J
	J)K
	K)L
	K)YOU
	I)SAN
	`

	items, err := itemFromText(data)
	assert.Nil(t, err)
	assert.NotNil(t, items)
	o := newOribitCalc(items)
	assert.NotNil(t, o)
	hops := o.calcMinimunOrbTransfers("YOU", "SAN")
	assert.Equal(t, 4, hops)
}
