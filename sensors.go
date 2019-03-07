package main

import (
	"fmt"
	"github.com/ssimunic/gosensors"
)

func main() {
	sensors, err := gosensors.NewFromSystem()
	// sensors, err := gosensors.NewFromFile("/path/to/log.txt")

	if err != nil {
		panic(err)
	}

	// Sensors implements Stringer interface,
	// so code below will print out JSON
	fmt.Println(sensors)

	// Also valid
	// fmt.Println("JSON:", sensors.JSON())

	// Iterate over chips
	for chip := range sensors.Chips {
		// Iterate over entries
		for key, value := range sensors.Chips[chip] {
			// If CPU or GPU, print out
			if key == "CPU" || key == "GPU" {
				fmt.Println(key, value)
			}
		}
	}
}
