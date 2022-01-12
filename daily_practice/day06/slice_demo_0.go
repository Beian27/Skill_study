package main 

import "fmt"

func main() {
  numbers := []int{0,1,2,3,4,5,6,7,8,9}
  printSlice(&numbers)

  fmt.Println(numbers[1:4])

  numbers_1 := make([]int,0,5)
  for i := 0; i < 100; i++ {
     numbers_1 = append(numbers_1,i)
  }
  printSlice(&numbers_1)
}

func printSlice(num *[]int ){
  fmt.Printf("len=%d cap=%d slice=%v\n", len(*num), cap(*num), *num)
}
