package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestImport(t *testing.T) {
	data := importFile("input.txt")
	assert.Equal(t, 15000, len(data))
}
