package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	
	"github.com/sanumala/vue-go-magic/service/api"
)

func main() {
	http.HandleFunc("/api/thumbnail", handler)
	fileserver := http.FileServer(http.Dir("./frontend/dist"))
	http.Handle("/", fileserver)
	
	fmt.Println("Starting server and listening at 3000")
	log.Panic(
		http.ListenAndServe(":3000", nil),
	)
}

//thumbnail Handler
func handler(w http.ResponseWriter, r *http.Request) {
	var decodedRequest api.ThumbnailRequest
	
	err := json.NewDecoder(r.Body).Decode(&decodedRequest)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	fmt.Printf("Given url is: %s\n", decodedRequest.Url)
	
	var apiResponse = api.ScreenShotApi(decodedRequest)
	
	// Pass back the screenshot URL to the frontend.
	_, err = fmt.Fprintf(w, `{ "screenshot": "%s" }`, apiResponse.Screenshot)
	if err != nil {
		fmt.Println("Something went wrong!!!", err)
	}
}
