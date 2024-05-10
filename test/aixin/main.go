package main

import (
	"fmt"
	"math"
)

func main() {
	printHeart()
}

func printHeart() {
	width := 20
	height := 10

	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			if isHeart(x, y, width, height) {
				fmt.Print("*")
			} else {
				fmt.Print(" ")
			}
		}
		fmt.Println()
	}
}

func isHeart(x, y, width, height int) bool {
	heartWidth := float64(width)
	heartHeight := float64(height)

	normalizedX := 2.0*float64(x)/heartWidth - 1.0
	normalizedY := 2.0*float64(y)/heartHeight - 1.0

	d := math.Pow(normalizedX*normalizedX+normalizedY*normalizedY-1.0, 3) -
		normalizedX*normalizedX*math.Pow(normalizedY, 3)

	return d < 0
}
