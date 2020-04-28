package problem1010_test

import (
	"leetcode_go/problem1010"
	"fmt"
)

func ExampleNumPairsDivisibleBy60(){
	tt := []int{10,20,30,40,50,60}
	num := problem1010.NumPairsDivisibleBy60(tt)
	fmt.Println(num)
	// OutPut: 2
}


//[]