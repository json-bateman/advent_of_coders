package main

import (
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
	"strings"
)

func calculateFuel(mass int) int {
	fuel := int(math.Floor(float64(mass)/3.0)) - 2
	if fuel < 0 {
		return 0
	}
	return fuel
}

func Day1() {
	content, err := os.ReadFile("day1.txt")
	if err != nil {
		log.Fatal(err)
	}

	fileAsString := string(content)
	var lines []string = strings.Split(fileAsString, "\n")
	var massArr []int
	for i := 0; i < len(lines); i++ {
		if lines[i] == "" {
			continue
		}
		num, err := strconv.Atoi(lines[i])
		if err != nil {
			log.Fatal(err)
		}
		massArr = append(massArr, num)
	}

	var finalFuel1 int
	for i := 0; i < len(massArr); i++ {
		finalFuel1 += calculateFuel(massArr[i])
	}

	// massArr fuel + fuel fuel
	var finalFuel2 int
	for i := 0; i < len(massArr); i++ {
		fuelMass := calculateFuel(massArr[i])
		for fuelMass > 0 {
			finalFuel2 += fuelMass
			fuelMass = calculateFuel(fuelMass)
		}
	}

	fmt.Println(finalFuel1)
	fmt.Println(finalFuel2)
}
