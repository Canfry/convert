package convert

import "fmt"

func Convert(from string, to string, value float64) float64 {
	if from == "C" && to == "F" {
		result := value*1.8 + 32
		fmt.Printf("%v Celcius is equivalent to %v Farenheit\n", value, result)
		return result
	}
	return 0

}
