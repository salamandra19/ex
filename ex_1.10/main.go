package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strconv"
	"time"
)

func main() {
	start := time.Now()
	ch := make(chan string)
	for filename, url := range os.Args[1:] {
		go fetch(strconv.Itoa(filename), url, ch)
	}
	for range os.Args[1:] {
		fmt.Println(<-ch)
	}
	fmt.Printf("%.2fs elapsed\n", time.Since(start).Seconds())
}
func fetch(filename string, url string, ch chan<- string) {
	start := time.Now()
	resp, err := http.Get(url)
	if err != nil {
		ch <- fmt.Sprint(err)
		return
	}
	file, err := os.Create(filename)
	if err != nil {
		log.Fatal(err)
	}
	nbytes, err := io.Copy(file, resp.Body)
	file.Close()
	resp.Body.Close()
	if err != nil {
		ch <- fmt.Sprintf("while reading %s: %v", url, err)
		return
	}
	secs := time.Since(start).Seconds()
	ch <- fmt.Sprintf("%.2fs %7d %s filename:%s", secs, nbytes, url, filename)
}
