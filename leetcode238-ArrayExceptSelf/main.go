package main

import "fmt"

// for every singal element we need got the element behind and after the target element which is prefix and postfix
func main() {
	array := []int{23, 2, 3, 4}
	result := betterSoluation(array)
	fmt.Println(result)
	result1 := bruteForce(array)
	fmt.Println(result1)
}
func bruteForce(array []int) []int {
	var result = make([]int, len(array))
	for i := 0; i < len(array); i++ {
		var tmp int = 1

		// fmt.Println("**********", array[i])
		for j := 0; j < len(array); j++ {
			if array[j] == array[i] {
				// fmt.Println(array[j])
				continue
			}
			tmp = tmp * array[j]
		}
		result[i] = tmp
		// fmt.Println("######", result[i])
	}

	return result

}
func betterSoluation(array []int) []int {
	var result = make([]int, len(array))
	for i := range result {
		result[i] = 1
	}

	var pre = make([]int, len(array))
	for i := range pre {
		pre[i] = 1
	}
	// 1 2 3 4
	// 1 1 1 1
	// 1*1 array[2-1]*1 array[3-1]*1 array[4-1]*2
	// 1 1 2 6

	// 4  3  2  1
	// 1  2  3  4
	// 1  1  1  1
	// 			array[len(array)-1]
	// 24 12 4  1

	//24 12 8 6

	//fmt.Println("before pre is ", pre)
	for i := 1; i < len(array); i++ {
		pre[i] = array[i-1] * pre[i-1]
	}
	//fmt.Println("after pre is ", pre)
	var post = make([]int, len(array))
	for i := range result {
		post[i] = 1
	}
	//	fmt.Println("before post is ", post)
	/*
		4,3,2,1
		post 不用算last 因为last 等于1

		先算 last-1(post[last-1] index=3 ) = array[last](value is 4) * post[last](value is 1)

		再算 last-2(index=2) = array[last-1](value is 3) * post[last-1](value is 4 刚刚已经算出来了)

		再算 last-3(index=1) = array

	*/
	for i := len(array) - 1; i > 0; i-- {
		post[i-1] = array[i] * post[i]
	}
	//fmt.Println("after post is ", post)
	//fmt.Println("end post  ")
	for i := range result {
		result[i] = pre[i] * post[i]
	}
	return result
}
