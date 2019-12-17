package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"

	"github.com/sirupsen/logrus"
)

func main() {
	data := importFile("input.txt")
	img := analyze(data, 25, 6)
	img.show()
}

func analyze(data []int, w, t int) *image {
	img := &image{layers: []*layer{}, minZ: 999999}
	numElements := w * t
	for i := 0; i < len(data); i += numElements {
		l := data[i : i+(w*t)]
		img.addLayer(&layer{data: l})
	}
	fmt.Printf("Num layers: %d\n", len(img.layers))
	fmt.Printf("Sum: %d\n", img.minLayer.num1digits*img.minLayer.num2digits)
	return img
}

type image struct {
	layers   []*layer
	w        int
	h        int
	minZ     int
	minLayer *layer
}

func (i *image) show() {
	//	imgLayer := &layer{}

}

func (i *image) addLayer(l *layer) {
	if i.layers == nil {
		i.layers = []*layer{}
		i.minZ = 999999
	}
	l.calc()
	i.layers = append(i.layers, l)
	if l.numZeros < i.minZ {
		i.minZ = l.numZeros
		i.minLayer = l
		fmt.Printf("UPDATED MIN NumZeros: %d num ones:%d num twos: %d\n", l.numZeros, l.num1digits, l.num2digits)
	}
}

type layer struct {
	data       []int
	numZeros   int
	num1digits int
	num2digits int
}

func (l *layer) calc() {
	for _, i := range l.data {
		if i == 0 {
			l.numZeros++
		}
		if i == 1 {
			l.num1digits++
		}
		if i == 2 {
			l.num2digits++
		}
	}
	//	fmt.Printf("NumZeros: %d num ones:%d num twos: %d\n", l.numZeros, l.num1digits, l.num2digits)
}

func importFile(fileName string) []int {
	fileBytes, err := ioutil.ReadFile(fileName)

	if err != nil {
		logrus.Fatal(err)
	}

	rc := []int{}
	for _, ch := range strings.TrimSpace(string(fileBytes)) {

		i, err := strconv.Atoi(string(ch))
		if err != nil {
			logrus.Fatalf("Error converting %v err:%s", ch, err)
		}
		rc = append(rc, i)
	}

	return rc
}
