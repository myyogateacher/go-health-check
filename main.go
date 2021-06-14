package main

import (
	"fmt"
	"net/http"
	"os"
)

func main() {
	if _, err := http.Get("http://127.0.0.1/status"); err != nil {
		fmt.Println("error", err)
		os.Exit(1)
	}
}
