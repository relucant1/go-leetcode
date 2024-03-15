package main

import (
	"fmt"
	"time"
)

func main() {
	start := time.Now()
	defer func() {
		fmt.Printf("Time taken: %s\n", time.Since(start))
	}()

	evilNinjas := []string{"Tommy", "John", "Bob", "Andy"}
	for _, evilNinja := range evilNinjas {
		go attack(evilNinja)
	}
	time.Sleep(time.Second * 2)
}

func attack(target string) {
	fmt.Printf("Attacking %s\n", target)
	time.Sleep(1 * time.Second)
}
