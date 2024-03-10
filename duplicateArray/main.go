package main

import (
	"fmt"
	"slices"
)

func main() {
	a := []int{5, 19, 2, 333, 2313, 290, 35, 21263, 313, 1, 312, 35, 3, 213, 12, 312312, 3123123}
	result := bruteForce(a)
	fmt.Printf("\nresult is %v ", result)
	result1 := optimalSolution(a)
	fmt.Printf("\nresult1 is %v ", result1)
}

func bruteForce(a []int) bool {
	var count int
	slices.Sort(a)
	fmt.Printf("a:", a)
	var flag bool
	flag = false
	for i := 0; i < len(a)-1; i++ {
		for j := i + 1; j < len(a); j++ {
			count++
			if a[i] == a[j] {
				fmt.Printf("\nvalue is %v ", a[i])
				flag = true
			}
		}
	}
	fmt.Printf("\ncount is%v", count)
	return flag
}

// optimizedAlg
func optimalSolution(array []int) bool {
	var count int
	var flag bool = false
	var m = make(map[int]int)
	for i := 0; i < len(array); i++ {
		count++
		_, ok := m[array[i]]
		if ok {
			fmt.Printf("\nduplicate value is%v ", array[i])
			flag = true
		} else {
			m[array[i]] = i
		}

	}
	fmt.Printf("\nmap is%v ", m)
	fmt.Printf("\ncount is%v", count)
	return flag

}
