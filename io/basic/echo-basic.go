package main
//a simple echo copycat, prints its command line arguments
import (
	"fmt"
	"os"
)

func main(){
	command := os.Args[0]
	sep, s := "", ""

	for _, arg := range os.Args[1:]{
		s += sep + arg
		sep = " "
	}

	fmt.Println(command + " " + s)
}