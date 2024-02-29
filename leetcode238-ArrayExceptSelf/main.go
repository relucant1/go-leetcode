package main

import "fmt"

func main() {
	array := []int{1, 2, 3, 4, 3, 0}
	var result = make([]int, len(array))
	for i := range result {
		result[i] = 1
	}
	var pre int = 1
	var post int = 1
	for i := range array {
		result[i] = pre
		pre = array[i] * pre
	}
	for i := len(array) - 1; i > 0; i-- {
		result[i] = post * result[i]
		post = array[i] * post
	}
	fmt.Println(result)
}
