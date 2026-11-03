package main

import (
	"fmt"
)

func isPalindrome(x int) bool {

    // all negative numbers are not palindromes
    if x < 0 {
        return false
    }

    //list of digits from right to left
    var digits []int

    divisor := 10

    //cycle through the array to find all digits
    for true {
        //find remainder
        rem := x % divisor
        previous_place_value := divisor/10

        //hard to explain so refer to this example eg 
        // we have 242 and the divisor is 100 
        //then we would be doing 242 - 42 
        digit := (rem - (x % (previous_place_value)))/previous_place_value

        digits = append(digits, digit)

        //stop condition: remainder is the same as x
        if rem == x {
            break
        }

        divisor = divisor * 10
    }

    for i, j := 0, len(digits) - 1 ; i <= j ; i, j = i + 1, j - 1 {
        if digits[i] != digits[j] {
            return false
        }
    }

    return true
}
func main(){

fmt.Println("case one(121 -> expect true): ", isPalindrome(121))
fmt.Println("case two(-121 -> expect false): ", isPalindrome(-121))
fmt.Println("case three(10 -> expect false): ", isPalindrome(10))
}
