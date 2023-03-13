package main

import "fmt"

func superAdd(numbers ...int) int{
	total := 0
	for _, numbers := range numbers{
		total = total + numbers
	}
	return total
}

func main() {
	result := superAdd(1, 2, 3, 4, 5, 6)
	fmt.Println(result)
}