package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"regexp"
)

// Google I/O 2012 - Go Concurrency Patterns
// Padrões recomendados pelo Rob Pike

//Padrão generator encapsula goroutine e gochannel, apenas chama-lo já resolve determinado situação

// <-chan - canal somente-leitura
func titulo(urls ...string) <-chan string {
	c := make(chan string)
	for _, url := range urls { //ignorando o indice _
		go func(url string) {
			resp, _ := http.Get(url)
			html, _ := ioutil.ReadAll(resp.Body)

			r, _ := regexp.Compile("<title>(.*?)<\\/title>")
			c <- r.FindStringSubmatch(string(html))[1]
		}(url)
	}
	return c
}

func main() {
	t1 := titulo("https://www.cod3r.com.br", "https://www.google.com")
	t2 := titulo("https://www.amazon.com", "https://www.youtube.com")
	fmt.Println("Primeiros:", <-t1, "|", <-t2)
	fmt.Println("Segunds:", <-t1, "|", <-t2)
}
