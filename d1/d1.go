package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"

	"github.com/sirupsen/logrus"
)

func main() {

	intA := importFileAsIntA("input.txt")

	fuel := calcFuelForMass(intA)
	fmt.Printf("Fuel before ff %d\n", fuel)
	extraFuel := calcFuelForFuelReq(fuel)

	fmt.Printf("Fuel :%d extra fuel %d total: %d\n", fuel, extraFuel, fuel+extraFuel)
}

func calcFuelForMassInkl(intA []int) int {
	fuel := 0
	for _, mass := range intA {
		mfuel := fuelForMass(mass)
		fuel += mfuel + calcFuelForFuelReq(mfuel)
	}
	return fuel
}

func calcFuelForMass(intA []int) int {
	fuel := 0
	for _, mass := range intA {
		fuel += fuelForMass(mass)
	}
	return fuel
}

func calcFuelForFuelReq(startFuel int) int {
	fuelForFuel := fuelForMass(startFuel)
	fuel := 0
	for fuelForFuel > 0 {
		fuel += fuelForFuel
		fuelForFuel = fuelForMass(fuelForFuel)
	}
	return fuel
}

func fuelForMass(mass int) int {
	fuel := (mass / 3) - 2
	if fuel < 0 {
		fuel = 0
	}
	return fuel
}

func linesToIntA(lines []string) []int {
	rc := []int{}
	for _, l := range lines {
		i, err := strconv.Atoi(l)
		if err != nil {
			logrus.Fatal(err)
		}
		rc = append(rc, i)
	}
	return rc
}

func importFileAsIntA(fileName string) []int {
	lines, err := importFile(fileName)
	if err != nil {
		logrus.Fatal(err)
	}
	return linesToIntA(lines)
}

func importFile(fileName string) ([]string, error) {
	fileBytes, err := ioutil.ReadFile(fileName)

	if err != nil {
		return nil, err
	}

	return strings.Split(string(fileBytes), "\n"), nil
}
