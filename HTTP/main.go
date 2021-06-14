package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

func main() {
	resp, err := http.Get("https://www.google.com/")

	if err != nil {
		fmt.Print("Error:", err)
		os.Exit(1)
	}

	// use of read function

	// bs := make([]byte, 99999)
	// resp.Body.Read(bs)
	// fmt.Println(string(bs))

	// writer/reader interface
	io.Copy(os.Stdout, resp.Body)
}
