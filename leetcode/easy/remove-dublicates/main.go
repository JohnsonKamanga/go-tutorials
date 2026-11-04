package main

import (
	"fmt"
)

func removeDuplicates(nums []int) int {
    length := 1

    for i := 1 ; i < len(nums) ; i++ {
        if nums[i] == nums[length - 1] {
            continue
        }

        target := nums[i]
        nums[i] = nums[length]
        nums[length] = target
        length++
    }
    return length
}

func main(){
fmt.Println("case one([1,1,2] -> expect [1,2] ): ", removeDuplicates([]int{1,1,2}))
fmt.Println("case two([0,0,1,1,1,2,2,3,3,4] -> expect [0,1,2,3,4]): ", removeDuplicates([]int{0,0,1,1,1,2,2,3,3,4}))

}