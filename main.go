package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func main(){
	http.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		log.Println("Hello World")
		data, err := ioutil.ReadAll(request.Body)
		if err!=nil {
			http.Error(writer, "Oops! Error occured", http.StatusBadRequest)
		}
		log.Printf("data = %s", data)
		fmt.Fprintf(writer, "Hey %s", data)
	})
	
	http.ListenAndServe(":9090", nil)
}
