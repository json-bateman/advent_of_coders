package main

import (
	"fmt"
	"slices"
	"strings"

	"github.com/jacktrusler/goutils"
)

type response2 struct {
	Page   int      `json:"page"`
	Fruits []string `json:"fruits"`
}

func something(foo int) (hi int, bye int) {
	hi = foo
	bye = 2
	return hi, bye
}

func main() {
	var file string
	file = goutils.FileAsString("./daylmao.txt")

	calorieArr := strings.Split(file, "\n\n")

	totalCalories := []int{}
	for _, elf := range calorieArr {
		cArr := strings.Split(elf, "\n")
		intArr := goutils.StringArrAtoI(cArr)
		sum := 0
		for _, i := range intArr {
			sum += i
		}
		totalCalories = append(totalCalories, sum)
	}
	answer := slices.Max(totalCalories)

	fmt.Println(answer)
}
