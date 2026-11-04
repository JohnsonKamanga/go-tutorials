package main

import (
	"fmt"
)

func longestCommonPrefix(strs []string) string {

    hasPrefix := func(s string, prefix string) bool{
        word := []rune(s)
            for i := 1 ; i <= len(word) ; i++{
                sub := string(word[:i])

                if sub == prefix {
                    return true
                }
            }
        return false
    }

    if len(strs) == 1{
        return strs[0]
    }
        
        
    possiblePrefixesCount := make(map[string]int)
    targetPrefix := ""
    first := []rune(strs[0])

    for i := 1 ; i <= len(first) ; i++ {
        prefix := string(first[:i])
        possiblePrefixesCount[prefix] = 1

        for j := 1 ; j < len(strs) ; j++ {
            if hasPrefix(strs[j], prefix) {
                possiblePrefixesCount[prefix]++ 
            }
        }
        if possiblePrefixesCount[targetPrefix] <= possiblePrefixesCount[prefix] && possiblePrefixesCount[prefix] == len(strs) {
            targetPrefix = prefix
        }

    }

    return targetPrefix
}

func main(){

fmt.Println("case one([\"flower\",\"flow\",\"flight\"] -> expect \"fl\"): ", longestCommonPrefix([]string{"flower","flow","flight"}))
fmt.Println("case two([\"dog\",\"racecar\",\"car\"] -> expect \"\"): ", longestCommonPrefix([]string{"dog","racecar","car"}))

}