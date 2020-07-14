package main

import (
	"fmt"
	"math"
	"regexp"
)

//These aren't sorts OR searches--they are just little study problems

// Write a function that determines if a string starts with an upper-case letter A-Z
func isUppercase1(s string) (bool, error) {
	matched, err := regexp.MatchString(`[A-Z]`, string(s[0]))
	return matched, err
}

func isUppercase2(s string) bool {
	matched, _ := regexp.MatchString(`[A-Z]`, string(s[0]))
	return matched
}

func isUppercase3(s string) bool {
	matched, err := regexp.MatchString(`[A-Z]`, string(s[0]))
	if err != nil {
		fmt.Print(err)
	}
	return matched
}

func TestIsUppercase() {
	tCase := "Dog"
	fCase := "cat"
	numberCase := "123"
	fmt.Printf("Testing %s, %s, %s \n", tCase, fCase, numberCase)
	fmt.Println(isUppercase1(tCase))
	fmt.Println(isUppercase2(fCase))
	fmt.Println(isUppercase3(numberCase))
}

// Write a function that determines the area of a circle given the radius
func calculateCircleArea(r float64) float64 {
	return math.Pi * math.Pow(r, 2)
}

func TestCalculateCircleArea() {
	oneTenth := calculateCircleArea(0.1)
	one := calculateCircleArea(1)
	oneHundred := calculateCircleArea(100)
	fmt.Printf("Testing %.8f, %.4f, %.2f \n", oneTenth, one, oneHundred)
	fmt.Println(oneTenth)
	fmt.Println(one)
	fmt.Println(oneHundred)
}

// Add up all the values in an array
func TestSumSlice() bool {
	return true
}

func main() {
	TestIsUppercase()
	TestCalculateCircleArea()
	// TestSumSlice()
}
