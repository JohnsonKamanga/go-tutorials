package main

import (
	"fmt"
)

func twoSum(nums []int, target int) []int {
    //a map of sums to their indices
    var pair []int
    found := false
    //cycle through each combination of elements
    for i := 0; i < len(nums) ; i++ {
        // inner loop for finding sums with other numbers
        for j := i + 1 ; j < len(nums) ; j++ {
            //find sum
            sum := nums[i] + nums[j]
            
            if sum == target {
                pair = append(pair, i, j)
                found = true
                break
            }
        }
        if found {
            break
        }
    }

    return pair
}

func main(){

testCaseOne := []int {2,7,11,15} // target = 9
testCaseTwo := []int {3,2,4} // target = 6
testCaseThree := []int {3,3} // target = 6

resOne := twoSum(testCaseOne, 9) // expect [0,1]
resTwo := twoSum(testCaseTwo, 6) // expect [1,2]
resThree := twoSum(testCaseThree, 6) // expect [0,1]

fmt.Println("case one: ", resOne)
fmt.Println("case two: ", resTwo)
fmt.Println("case three: ", resThree)
}
