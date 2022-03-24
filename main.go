package main

import (
	"Logic_Test/maxnumber"
	"fmt"
)

func main() {
	fmt.Println(maxnumber.MaxValueNumber([]int{1, 2, 3, 8, 9, 3, 2, 1}))
	fmt.Println(maxnumber.MaxValueNumber([]int{1, 2, 1, 2, 2, 1}))
	fmt.Println(maxnumber.MaxValueNumber([]int{7, 1, 2, 9, 7, 2, 1}))

}
