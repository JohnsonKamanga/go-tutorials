package main 

import (
	"fmt"
)

func romanToInt(s string) int {

    // we start by creating some helper functions to make the logic easier to manage
    // this function simple returns the value of an individual numeral
    getValue := func(numeral rune) int {
        switch numeral {
            case 'I':
                return 1
            case 'V':
                return 5
            case 'X':
                return 10
            case 'L':
                return 50
            case 'C':
                return 100
            case 'D':
                return 500
            case 'M':
                return 1000
            default :
                return -1 // indicates invalid character
        }
    }

    roman := []rune(s)
    sum := 0 // running sum variable

    for i := 0 ; i < len(roman) ; i++ {
        firstValue := getValue(roman[i])

        if i == len(roman) - 1 {
            sum += firstValue
            continue
        }

        secondValue := getValue(roman[i + 1])

        if firstValue < secondValue {
            sum += secondValue - firstValue
            i++ //advance i by one so that we skip the secondValue character
        } else {
            sum += firstValue
        }
    }

    return sum
}

func main(){

fmt.Println("case one(III -> expect 3): ", romanToInt("III"))
fmt.Println("case two(LVIII -> expect 58): ", romanToInt("LVIII"))
fmt.Println("case three(MCMXCIV -> expect 1994): ", romanToInt("MCMXCIV"))
}