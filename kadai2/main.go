package main

import "fmt"

func tasu(a string, x float64, y float64) float64 {
	total := x + y
	return total
}

func hiku(a string, x float64, y float64) float64 {
	total := x - y
	return total
}

func kakeru(a string, x float64, y float64) float64 {
	total := x * y
	return total
}

func waru(a string, x float64, y float64) float64 {
	total := x / y
	return total
}

func main() {

	fmt.Println(tasu("+",4.4532, 2.109))
	fmt.Println(hiku("-",4.4532, 2.109))
	fmt.Println(kakeru("*",4.4532, 2.109))
	fmt.Println(waru("/",4.4532, 2.109))
}