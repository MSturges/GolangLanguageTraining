package main

import "fmt"

//find project euler problem and solve it

// Even Fibonacci numbers
// Problem 2
// Each new term in the Fibonacci sequence is generated by adding the previous two terms. By starting with 1 and 2, the first 10 terms will be:

// 1, 2, 3, 5, 8, 13, 21, 34, 55, 89, ...

// By considering the terms in the Fibonacci sequence whose values do not exceed four million, find the sum of the even-valued terms.

func main() {
	fmt.Println(sumFibonacciEvenToFourMillion())
}

func sumFibonacciEvenToFourMillion() int {
	slice := []int{1, 2}

	for ind := 2; slice[ind-1]+slice[ind-2] <= 4000000; ind++ {
		slice = append(slice, slice[ind-1]+slice[ind-2])
	}

	evenSum := 0

	for _, val := range slice {
		if val%2 == 0 {
			evenSum += val
		}
	}

	return evenSum
}