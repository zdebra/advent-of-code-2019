// https://adventofcode.com/2019/day/1
package main

import (
	"bufio"
	"fmt"
	"io"
	"math"
	"os"
	"strconv"
)

type Module struct {
	Mass float64
}

func (m Module) FuelRequiredToLaunch() float64 {
	calc := func(inp float64) float64 {
		return math.Trunc(	inp / 3 ) - 2
	}

	var total float64
	tmp := calc(m.Mass)
	for ; tmp > 0; tmp = calc(tmp) {
		total += tmp
	}
	return total
}

type Fleet []Module

func (f Fleet) FuelRequiredToLaunch() float64 {
	var total float64
	for _, module := range f {
		total += module.FuelRequiredToLaunch()
	}
	return total
}

func calculateFleetFuel(r io.Reader) (float64, error) {
	sc := bufio.NewScanner(r)
	f := Fleet{}
	for sc.Scan() {
		mass, err := strconv.ParseFloat(sc.Text(), 64)
		if err != nil {
			return 0, fmt.Errorf("failed to parse mass: %w", err)
		}
		f = append(f, Module{Mass:mass})
	}
	return f.FuelRequiredToLaunch(), nil
}

func main() {
	f, err := os.Open("day01/input")
	panicOnErr(err)

	fleetFuel, err := calculateFleetFuel(f)
	panicOnErr(err)

	fmt.Printf("total mass required for fleet is %f\n", fleetFuel)

}

func panicOnErr(err error) {
	if err != nil {
		panic(err)
	}
}
