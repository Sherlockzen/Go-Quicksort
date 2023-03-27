package main

import "fmt"

func main() {

	fmt.Printf("Contagem da lista [1 10 5 7 11] = %d\n", qsort([]int{7, 10, 5, 1, 11}))

}

// QUICKSORT FUNC
func qsort(arr []int) []int {

	//Base case
	if len(arr) < 2 {
		return arr
	}

	//Recursion
	pivot := arr[0]
	var minors []int
	var greaters []int
	for _, value := range arr[1:] {
		if value <= pivot {
			minors = append(minors, value)
		} else {
			greaters = append(greaters, value)
		}
	}

	minors = qsort(minors)
	greaters = qsort(greaters)

	return append(append(minors, pivot), greaters...)
}
