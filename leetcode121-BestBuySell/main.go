package main

import "fmt"

func main() {
	a := []int{5, 1, 2, 3, 4}
	result := bruteAlg(a)
	fmt.Println("result is ", result)
}

func bruteAlg(a []int) int {
	var profit int
	for i := 0; i < len(a)-1; i++ {
		for j := i + 1; j < len(a); j++ {
			profit = max(a[j]-a[i], profit)
		}
	}
	return profit
}

// optimizedAlg
func optimizedAlg() {

}
