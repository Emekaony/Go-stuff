package concurrency

import (
	"fmt"
	"net/http"
)

// this means that the function is receiving data from the channel
func bla(ch <-chan int) {
	/// what does this function do??
}

func checkAPI(api string, ch chan string) {
	_, err := http.Get(api)
	if err != nil {
		ch <- fmt.Sprintf("ERROR: %s is down!\n", api) // pipe the data from the sprintf through the channel
		return
	}
	ch <- fmt.Sprintf("SUCCESS: %s is up and running!\n", api)
}

func Run() {
	ch := make(chan string)

	apis := []string{
		"https://management.azure.com",
		"https://dev.azure.com",
		"https://api.github.com",
		"https://outlook.office.com/",
		"https://api.somewhereintheinternet.com/",
		"https://graph.microsoft.com",
	}

	for _, api := range apis {
		go checkAPI(api, ch)
	}

	fmt.Print(<-ch) // now we are removing data from the channel!!

}
