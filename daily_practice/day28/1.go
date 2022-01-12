package main



func max(arr []int) []int {
	if len(arr) != 0{
		return nil
	}

	max1 := arr[0]
	min := arr[0]
	res := arr[0]

	for i:=1; i < len(arr); i++ {
		max1 = max1 * arr[i]
		min = min * arr[i]
		max1 =  Max(Max(max1,min), arr[i])
		min =  Min(Min(max1,min), arr[i])
		res = Max(res, max1)
	}

	return nil
}

func Max(x,y int) int {
	if x > y {
		return x
	}
	return y
}

func Min(x,y int) int {
	if x > y {
		return y
	}
	return x
}
