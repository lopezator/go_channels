package main

import (
	"net/http"
	"fmt"
	"time"
)

func main() {
	links := []string {
		"http://facebook.com",
		"http://stackoverflow.com",
		"http://golang.org",
		"http://amazon.com",
		"http://google.com",
	}

	c := make(chan string)

	for _, link := range links {
		go checkLink(link, c)
	}

	/*for i := 0; i < len(links); i++ {
		//Esta linea es bloqueante
		//Es por ello que se ejecuta en otro for distinto
		//Si lo ejecutamos en el if de arriba las llamadas no serías concurrentes
		//Porque esperaría a que terminase cada go func
		//Antes de lanza la siguiente
		fmt.Println(<- c)
	}*/

	// Loop infinito, el primer argumento cuela como
	// String al ser un canal de tipo string
	for l := range c {
		go func(link string) {
			time.Sleep(5 * time.Second)
			checkLink(link, c)
		}(l)
	}
}

func checkLink(link string, c chan string) {
	_, err := http.Get(link)

	if err != nil {
		fmt.Println(link, "might be down!")
		c <- link
		return
	}

	fmt.Println(link, "is up!")
	c <- link
}
