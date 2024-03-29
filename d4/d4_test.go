package main

import (
	"fmt"
	"testing"

	"gotest.tools/assert"
)

func TestHasDouble(t *testing.T) {
	var tests = []struct {
		num      int64
		isDouble bool
		isIncreasting bool
	}{
		{111111, true,true},
		{223450, true,false},
		{123789, false,true},
	}
	for _, tt := range tests {
		t.Run(fmt.Sprintf("Testing %d", tt.num), func(tx *testing.T) {
			exp := hasDouble(tt.num)
			assert.Equal(tx, tt.isDouble, exp)
			isInc := isIncreasting(tt.num)
			assert.Equal(tx, tt.isIncreasting, isInc)
		})
	}

}


func TestHasOnlyDouble(t *testing.T){
	var tests = []struct {
		num      int64
		exp bool
	}{
		{477778,false},
		{477779,false},
		{477788,true},
		{477789,false},
	}
	for _, tt := range tests {
		t.Run(fmt.Sprintf("Testing %d", tt.num), func(tx *testing.T) {
			exp := hasOnlyDouble(tt.num)
			assert.Equal(tx, tt.exp, exp)
		})
	}
}