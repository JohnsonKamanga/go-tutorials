package main

import (
	"fmt"
	"os"
	"strconv"
)

func main(){
	args := os.Args[1:]
	set := []int{}
	minIndex := 0
	maxIndex := 0

	//check if enough args were provided

	if len(args) < 1 {
		fmt.Fprintf(os.Stderr, "No numbers were provided\n")
		os.Exit(-1)
	}

	// convert strings to numbers
	for _, val  := range args {
		res, err := strconv.Atoi(val)

		if err != nil {
			fmt.Fprintf(os.Stderr, "An error occured when trying to process user input: %v", err)
			os.Exit(-1)
		}

		set = append(set, res)
	}

	// find minimum and maximum numbers from the list
	for i := range set {
		//check if minimum
		if set[i] <= set[minIndex]{
			minIndex = i
		} else if set[i] > set[maxIndex] { // check if maximum
			maxIndex = i
		}
	}

	//print maximum and minimum
	fmt.Println("Minimum number: ", set[minIndex])

	fmt.Println("Maximum number: ", set[maxIndex])
}
