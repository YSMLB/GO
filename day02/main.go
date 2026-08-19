package main

import "fmt"

func main() {
	var s []int // создали пустой срез

	for i := 0; i < 30; i++ {
		s = append(s, i)
	}

	filterEven := filterEven(s)
	fmt.Println(filterEven)
}

func filterEven(nums []int) []int {
	i := 0 // write
	for j := 0; j < len(nums); j++ { // read
		if nums[j]%2 == 0 {
			nums[i] = nums[j]
			i++
		}
		fmt.Printf("len: %2d; cap: %2d; addres: %p\n", len(nums), cap(nums), &nums[0])

	}

	nums = nums[:i]
	return nums
}
