package handlers

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

type Hello struct {
	l *log.Logger
}

func NewHello(l *log.Logger) *Hello{
	return &Hello{l}
}

func (hello *Hello) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	hello.l.Println("Hello World")
	data, err := ioutil.ReadAll(r.Body)
	if err!=nil {
		http.Error(rw, "Oops! Error occured", http.StatusBadRequest)
	}
	hello.l.Printf("data = %s", data)
	fmt.Fprintf(rw, "Hey %s", data)
}
