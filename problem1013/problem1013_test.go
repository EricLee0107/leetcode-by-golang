package problem1013

import (
	"fmt"
)

func ExampleNumPairsDivisibleBy60(){
	tt := []int{3,3,6,5,-2,2,5,1,-9,4}
	num := canThreePartsEqualSum(tt)
	fmt.Println(num)
	tt1 := []int{0,2,1,-6,6,-7,9,1,2,0,1}
	num1 := canThreePartsEqualSum(tt1)
	fmt.Println(num1)
	tt2 := []int{1,-1,1,-1}
	num2 := canThreePartsEqualSum(tt2)
	fmt.Println(num2)
	// OutPut:
	// true
	// true
	// false
}


//[]