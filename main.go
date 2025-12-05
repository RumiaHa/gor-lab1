package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())
	urls := []string{"google.com", "yandex.ru", "github.com", "stackoverflow.com"}

	var wg sync.WaitGroup
	duration := time.Duration(rand.Intn(3)) * time.Second


	wg.Add(len(urls))
	fmt.Println("---Going trough sites---")
	for _,url := range urls{
		
		go func(ur string){
			defer wg.Done()
			fmt.Printf("Checking %s\n",url)
			
			time.Sleep(duration)
			
			fmt.Printf("%s site is checked\n", url)
		}(url)
	}
	wg.Wait()
	fmt.Println("---Site checking is completed---")

}