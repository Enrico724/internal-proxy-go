package main

import (
	"fmt"
	"github.com/Enrico724/internal-proxy-go"
	"net/http"
)

func helloWorld(w http.ResponseWriter, _ *http.Request) {
	_, _ = w.Write([]byte("Hello World!"))
}

func main() {
	http.HandleFunc("/", helloWorld)
	err := internal_proxy.Serve()
	if err != nil {
		fmt.Printf("%v\n", err)
	}
}
