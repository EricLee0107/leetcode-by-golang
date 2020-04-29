package problem1013

func canThreePartsEqualSum(A []int) bool {
	var sum int64
	for _,v := range A{
		sum += int64(v)
	}
	if sum % 3 != 0{
		return false
	}else{
		sum1 := sum/3
		var i,n int
		var j int64
		for i,j=0,sum1;i<len(A); i++{
			j -= int64(A[i])
			if j == 0 {
				j=sum1
				n++
			}
		}
		if n ==3 || (n>3 && sum1==0){
			return true
		}else {
			return false
		}
	}
}