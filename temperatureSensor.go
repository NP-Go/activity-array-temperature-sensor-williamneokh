package main

import "fmt"

func main() {
	var sumTemp float64
	var totalHour int

	var temperature = [24]float64{20.1, 24, 27.3, 30.1, 26.4, 22.2, 20.1, 24, 27.3, 30.1, 26.4, 20.1, 24, 27.3, 30.1, 26.4, 20.1, 24, 27.3, 30.1, 26.4, 20.1, 24, 27.3}
	for _, temp := range temperature {
		sumTemp = sumTemp + temp
		totalHour = totalHour + 1
	}
	averageTemp := sumTemp / float64(totalHour)
	fmt.Printf("The average temperature of a day is: %.2f\n", averageTemp)
	//Insert Code here
}
