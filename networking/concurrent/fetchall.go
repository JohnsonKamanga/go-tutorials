package main


import (
	"fmt"
	"os"
	"net/http"
	"time"
	"io"
)

func main(){
	start := time.Now()
	ch := make(chan string)
	
	for _, url := range os.Args[1:]{
		go fetch(url, ch)
	}

	for range os.Args[1:]{
		fmt.Print(<-ch)
	}

	duration := time.Since(start).Seconds()

	fmt.Printf("%.2f elapsed", duration)


}

func fetch(url string, ch chan <- string){
	start := time.Now()
	res, err := http.Get(url)
	
	if err != nil {
		ch <- fmt.Sprint(err, "\n")
		return
	}

	nBytes, err := io.Copy(io.Discard, res.Body)

	if err != nil {
		ch <- fmt.Sprint(err)
		return
	}

	duration := time.Since(start).Seconds()

	ch <- fmt.Sprintf("%.2f %7d %s", duration, nBytes, url + "\n")
}