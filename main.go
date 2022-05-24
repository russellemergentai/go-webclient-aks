package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"sync"
)

func call(wg *sync.WaitGroup, ch chan string) {

	resp, err := http.Get("http://webcode.me")

	if err != nil {
		log.Fatal(err)
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		log.Fatal(err)
	}

	ch <- string(body)

	wg.Done()
}

func main() {
	wg := new(sync.WaitGroup)
	N := 5
	wg.Add(N)

	ch := make(chan string, N)

	for i := 0; i < N; i++ {
		log.Println(fmt.Sprintf("starting web call %d\n", i))
		go call(wg, ch)
	}

	wg.Wait()

	for elem := range ch {
		fmt.Println(elem)
	}

}
