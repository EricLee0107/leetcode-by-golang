package problem1010

func NumPairsDivisibleBy60(time []int) int {
	mm := make(map[int]int,60)
	for _,v := range time{
		mm[v%60] ++
	}
	num := (mm[0] * (mm[0] - 1)) /2  + (mm[30]*(mm[30] - 1))/2
	for j := 1;j<30;j++{
		if mm[j] != 0 {
			num += mm[60-j] * mm[j]
		}
	}
	return num
}