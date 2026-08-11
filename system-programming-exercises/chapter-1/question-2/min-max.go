package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	set := []int{}
	minIndex, maxIndex := 0, 0
	reader := bufio.NewReader(os.Stdin)

	//prompt for user input
	fmt.Print("Enter your list of numbers separated by a single space: ")

	// get user input
	input, err := reader.ReadString('\n')

	//print string

	fmt.Println("\n", input)

	if err != nil {
		fmt.Fprintf(os.Stderr, "An error occured when reading user input: %v\n", err)
		os.Exit(-1)
	}

	// split string by spaces. NB the new line character remains in the input string as the last character
	strInputs := strings.Split(input, " ")

	//helper function which remove news line character from numbers
	normalizeString := func(x string) string {
		if strings.Contains(x, "\n") {
			return x[:len(x)-1]
		} else {
			return x
		}
	}

	// convert each string to a number
	for _, val := range strInputs {
		//convert each string into a number
		res, err := strconv.Atoi(normalizeString(val))

		if err != nil {
			fmt.Fprintf(os.Stderr, "An error occured when parsing user input: %v\n", err)
			os.Exit(-1)
		}

		set = append(set, res)

	}

	// find minimum and maximum numbers in set

	for i := range set {
		// check if minimum
		if set[i] <= set[minIndex] {
			minIndex = i
		} else if set[i] > set[maxIndex] { // check if maximum
			maxIndex = i
		}
	}

	//print maximum and minimum
	fmt.Println("Minimum number: ", set[minIndex])

	fmt.Println("Maximum number: ", set[maxIndex])

}
