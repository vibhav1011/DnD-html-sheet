package main

import (
	"fmt"
	"net/http"
	"net/http/httputil"
	"encoding/json"
	"log"
)


func save_sheet(response http.ResponseWriter, request *http.Request) {
	fmt.Printf("got /save request\n")
	
	body, body_err := httputil.DumpRequest(request, true)
	if body_err != nil {
		fmt.Printf("could not read body: %s\n", body_err)
	}

	fmt.Printf("saving body to sheet : ", string(body))

	response_map := map[string]interface{}{"status": "ok"}
	response.Header().Set("Content-Type", "application/json")
	response.WriteHeader(http.StatusOK)

	json.NewEncoder(response).Encode(response_map)
}


func main() {
	http.HandleFunc("/save", save_sheet)

	port := 3333
	var server_status string = fmt.Sprintf("server started on port: %d", port)

	log.Println(server_status)
	log.Fatal( http.ListenAndServe(fmt.Sprintf(":%d", port), nil) )	

}
