package main

import "fmt"

func convertTemperature(celsius float64) []float64 {
	return []float64{celsius + 273.15, celsius*1.8 + 32}
}
func main() {
	celsius := 36.50
	fmt.Println(convertTemperature(celsius))
}
