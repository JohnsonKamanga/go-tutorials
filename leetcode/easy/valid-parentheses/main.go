package main

import (
	"fmt"
)

func isValid(s string) bool {
    isOpeningChar := func(c rune) bool{
        if c == '(' || c == '{' || c == '[' {
            return true
        } else {
            return false
        } 
    }

    isCorrectClosing := func(o rune, c rune) bool{
        switch o {
            case '(':
                return c == ')'
            case '[':
                return c == ']'
            case '{':
                return c == '}'
            default:
                return false
        }
    }

    var stack []rune
    head := 0
    chars := []rune(s)

    for i := 0 ; i < len(chars) ; i++ {
        //if is an opening character, push on to stack
        if isOpeningChar(chars[i]) {
            stack = append(stack, chars[i])
            head++
        } else {
            //check if previous key is the correct opening bracket
            if i > 0 && head > 0 && isCorrectClosing(stack[head - 1], chars[i]){
                //pop opening char off the stack
                stack = append(stack[:head - 1])
                head--
            } else {
                return false
            }
        }
    }

    if len(stack) > 0 {
        return false
    }

    return true
    
}

func main(){

fmt.Println("case one(\"()\" -> expect true): ", isValid("()"))
fmt.Println("case two(\"()[]{}\" -> expect true): ", isValid("()[]{}"))
fmt.Println("case three(\"(]\" -> expect false): ", isValid("(]"))
fmt.Println("case four(\"([])\" -> expect true): ", isValid("([])"))
fmt.Println("case five(\"([)]\" -> expect false): ", isValid("([)]"))
}
