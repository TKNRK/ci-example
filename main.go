package main

import "fmt"

// slice に target が含まれるかどうか判定する
func include(slice []int, target int) bool {
	for _, v := range slice {
		if v == target {
			return true
		}
	}
	return false
}

func unusedFunc() {
	fmt.Println("unused!")
}

func main() {
	s1 := []int{1, 2, 3}
	if include(s1, 1) {
		fmt.Println("OK!")
	}
}
