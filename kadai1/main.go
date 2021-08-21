package main

import "fmt"

func hiku(x float64, y float64) float64 {
	total := x - y
	return total
}

func kakeru(x float64, y float64) float64 {
	total := x * y
	return total
}

func waru(x float64, y float64) float64 {
	total := x / y
	return total
}

func main() {

	fmt.Println(hiku(4.4532, 2.109))
	fmt.Println(kakeru(4.4532, 2.109))
	fmt.Println(waru(4.4532, 2.109))
}