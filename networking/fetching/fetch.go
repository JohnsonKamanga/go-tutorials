package main

import (
	"fmt"
	"io"
	"os"
	"net/http"
)

func main(){
	for _, url := range os.Args[1:] {
		res, err := http.Get(url)
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch: reading %s: %v \n", url, err)
			os.Exit(1)
		}
		body, err := io.ReadAll((res.Body))

		if err != nil {
		fmt.Fprintf(os.Stderr, "fetch: reading %s: %v\n", url, err)	
		os.Exit(1)
		}

		fmt.Printf("%s", body)
	}
}