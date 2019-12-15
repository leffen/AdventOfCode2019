package main

import (
	"fmt"
	"io/ioutil"
	"strings"

	"github.com/sirupsen/logrus"
)

func main() {
	part1()
}

func part1() {
	items, err := importFile("input.txt")
	if err != nil {
		logrus.Fatal(err)
	}
	o := newOribitCalc(items)
	cnt := o.calcAllOrbits()

	fmt.Printf("Num orbits=%d\n", cnt)

}

type item struct {
	name         string
	orbitsAround string
	numParents   int
}

func importFile(fileName string) (map[string]*item, error) {
	fileBytes, err := ioutil.ReadFile(fileName)

	if err != nil {
		return nil, err
	}
	return itemFromText(string(fileBytes))
}

func itemFromText(data string) (map[string]*item, error) {
	rc := map[string]*item{}

	for _, line := range strings.Split(string(data), "\n") {
		if strings.TrimSpace(line) == "" {
			continue
		}
		fields := strings.Split(line, ")")
		var prev *item
		for _, f := range fields {
			name := strings.TrimSpace(f)
			i, ok := rc[name]
			if !ok {
				i = &item{name: strings.TrimSpace(f)}
			}

			if prev != nil {
				i.orbitsAround = strings.TrimSpace(prev.name)
			}
			rc[name] = i
			prev = i
		}
	}
	return rc, nil
}

type oribitCalc struct {
	mp map[string]*item
}

func newOribitCalc(items map[string]*item) *oribitCalc {
	o := &oribitCalc{mp: items}

	totalOrbitals := 0
	for _, i := range items {
		if i.orbitsAround != "" {
			i.numParents = o.numParents(i.orbitsAround, 1)
			//fmt.Printf("updating %v\n", i)
			totalOrbitals += i.numParents
		}
	}
	//fmt.Printf("FOUND %d orbitals \n", totalOrbitals)

	return o
}

func (o *oribitCalc) numParents(name string, startNum int) int {
	i, ok := o.mp[name]
	if !ok || i.orbitsAround == "" {
		//	fmt.Printf("No parents found for %s\n", name)
		return startNum
	}
	return o.numParents(i.orbitsAround, startNum+1)
}

func (o *oribitCalc) calcAllOrbits() int {
	rc := 0
	for _, i := range o.mp {
		rc += i.numParents
	}

	return rc
}
