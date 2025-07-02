package main

import (
	"fmt"
	"net/http"
)

func runChannels() {
	links := []string{
		"http://google.com",
		"http://facebook.com",
		"http://stackoverflow.com",
		"http://golang.org",
		"http://amazon.com",
	}

	c := make(chan string)
	//Establishes channel
	for _, link := range links {
		go checkLink(link, c)
	}
	//Responds to the channel
	// for i := 0; i < len(links); i++ {
	// 	fmt.Println(<-c)
	// }
	//Responds to the channel, infinite loop
	fmt.Println(<-c)
	for l := range c { //after channel returns value, assign it to l
		// go func() { //this is an anonymous function called a function literal
		// 	// time.sleep(5 * time.Second)
		// 	checkLink(l, c)
		// }()
		go func(link string) { //this is an anonymous function called a function literal
			// time.sleep(5 * time.Second)
			checkLink(l, c)
		}(l)
	}
}

// func checkLink(link string, c chan string) {
// 	_, err := http.Get(link)
// 	if err != nil {
// 		fmt.Println(link, "might be down!")
// 		c <- "Might be down I think"
// 		return
// 	}

//		fmt.Println(link, "is up!")
//		c <- "Yep its up"
//	}
func checkLink(link string, c chan string) {
	_, err := http.Get(link)
	if err != nil {
		fmt.Println(link, "might be down!")
		c <- link //sends the link text to the channel
		return
	}

	fmt.Println(link, "is up!")
	c <- link
}
