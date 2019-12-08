package main

import (
	"fmt"
	"strconv"
)

func main() {
	min := int64(372304)
	max := int64(847060)
	numFound := 0
	for i := min; i < max; i++ {
		if isIncreasting(i) && hasDouble(i) {
			numFound++
		}
	}
	fmt.Printf("Found :%d\n", numFound)
}

func hasDouble(num int64) bool {
	n := strconv.FormatInt(num, 10)

	for i := 0; i < len(n)-1; i++ {
		if n[i] == n[i+1] {
			return true
		}
	}

	return false
}

func isIncreasting(num int64) bool {
	n := strconv.FormatInt(num, 10)

	for i := 0; i < len(n)-1; i++ {
		if n[i] > n[i+1] {
			return false
		}
	}
	return true
}
