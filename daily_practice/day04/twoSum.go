package main

func twoSum() []int {
	nums := [4]int{2,7,11,15}
	target := 9
	hashTable := map[int]int{}
	for i, x := range nums {
		println(i)
		println(x)
		if p, ok := hashTable[target-x]; ok {
			return []int{p, i}
		}
		hashTable[x] = i

	}
	return nil
}

func main() {

	twoSum()
}