package main

import (
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
)


func save_sheet(response http.ResponseWriter, request *http.Request) {
	fmt.Printf("got /save request\n")
	
	body, body_err := ioutil.ReadAll(request.Body)
	if body_err != nil {
		fmt.Printf("could not read body: %s\n", body_err)
	}

	fmt.Printf("saving %s to sheet ...\n", body)
	io.WriteString(response, "saving to sheet ...\n")
}


func main() {
	http.HandleFunc("/save", save_sheet)

	err := http.ListenAndServe(":3333", nil)
	if errors.Is(err, http.ErrServerClosed) {
		fmt.Printf("server closed\n")
	} else if err != nil {
		fmt.Printf("error starting server: %s\n", err)
		os.Exit(1)
	}
}
