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
	fmt.Printf("Part 1 : Found :%d\n", numFound)

	numFound = 0

	for i := min; i < max; i++ {
		if isIncreasting(i) && hasOnlyDouble(i) {
			numFound++
			//	fmt.Printf("%d\n", i)
		}
	}
	fmt.Printf("Part 2 : Found :%d\n", numFound)

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

func hasOnlyDouble(num int64) bool {
	n := strconv.FormatInt(num, 10)
	l := len(n)
	var prev byte

	for i := 0; i < l-1; i++ {
		if n[i] == n[i+1] {
			if (prev != n[i]) && (i+3 > l) {
				return true
			}

			if (prev != n[i]) && (n[i+1] != n[i+2]) {
				fmt.Printf("%d n[i]=%s n[i+1]=%s n[i+2]=%s\n", num, string(n[i]), string(n[i+1]), string(n[i+2]))
				return true
			}
		}
		prev = n[i]
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
