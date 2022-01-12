package main

import (
	"fmt"
	"time"
)

func addNumberToChan(chanName chan int) {
	for {
		chanName <- 1
		time.Sleep(1 * time.Second)
	}
}

func main() {
	var chan1 = make(chan int, 10)
	var chan2 = make(chan int, 10)

	go addNumberToChan(chan1)
	go addNumberToChan(chan2)

	for {
		select {
		case e := <-chan1 :
			fmt.Println("chan1 = %d", e)
		case e := <-chan2 :
			fmt.Println("chan2 = %d", e)
		default:
			fmt.Println("over!!!")
			time.Sleep(1 * time.Second)
		}
	}
}


func maxSubArray(nums []int) int {
	max := nums[0]
	for i := 1; i < len(nums); i++ {
		if nums[i] + nums[i-1] > nums[i] {
			nums[i] += nums[i-1]
		}
		if nums[i] > max {
			max = nums[i]
		}
	}
	return max
}

func maxProfit(prices []int) int {
	max := 0
	//买
	for i := 0; i < len(prices); i++ {
		//卖
		for j := i + 0; j < len(prices); j++ {
			sum := prices[j] - prices[i]
			if sum > max {
				max = sum
			}
		}
	}
	return max
}

