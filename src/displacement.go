package main

import (
	"os"
	"bufio"
	"fmt"
	"math"
	"strconv"
)

func GenDisplaceFn(a, v0, s0 float64) func(float64) float64 {
	return func(t float64) float64 {
		return 0.5*a*math.Pow(t, 2) + v0*t + s0
	}
}

func main() {

	var acce, vInit, sInit, t_in float64

	scanner := bufio.NewScanner(os.Stdin)

	fmt.Print("Acceleration: ")
	scanner.Scan()
	acce, _ = strconv.ParseFloat(scanner.Text(), 64)
	fmt.Print("Initial Velocity: ")
	scanner.Scan()
	vInit, _ = strconv.ParseFloat(scanner.Text(), 64)
	fmt.Print("Initial Displacement: ")
	scanner.Scan()
	sInit, _ = strconv.ParseFloat(scanner.Text(), 64)
	fmt.Print("At time (enter a value for t): ")
	scanner.Scan()
	t_in, _ = strconv.ParseFloat(scanner.Text(), 64)

	fn := GenDisplaceFn(acce, vInit, sInit)

	fmt.Printf("Displacememtn is %.2f\n", fn(t_in))

}
