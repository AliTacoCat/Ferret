package main

import (
	"errors"
	"fmt"
)

func StopErroring() {
	var printValue string = "He, World!"
	printMe(printValue)

	var numerator int = 11
	var denominator int = 0
	var result, remainder, err = intDivision(numerator, denominator)
	if err != nil {
		fmt.Println(err.Error())
	} else if remainder == 0 {
		fmt.Printf(" The result of the interger division is %v with no remainder", result)
	} else {
		fmt.Printf(" The result of the interger division is %v with remainder %v", result, remainder)
	}

	fmt.Printf(" The result of the interger division is %v with remainder %v", result, remainder)
}

func printMe(printValue string) {
	fmt.Println(printValue)
}

func intDivision(numerator int, denominator int) (int, int, error) {
	var err error
	if denominator == 0 {
		err = errors.New(" Denominator cannot be zero")
		return 0, 0, err
	}
	var result int = numerator / denominator
	var remainder int = numerator % denominator
	return result, remainder, err
}
