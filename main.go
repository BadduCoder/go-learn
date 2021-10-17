package main

import (
	"context"
	"golearn/handlers"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"
)

func main(){
	logger := log.New(os.Stdout, "go-learn", log.LstdFlags)
	helloHandler := handlers.NewHello(logger)

	sm := http.NewServeMux()
	sm.Handle("/", helloHandler)

	server := &http.Server{
		Addr : ":9090",
		ReadTimeout: 1 * time.Second,
		WriteTimeout: 1 * time.Second,
		IdleTimeout: 120 * time.Second,
		Handler: sm,
	}

	go func() {
		err := server.ListenAndServe()
		if err!=nil {
			log.Fatal(err)
		}
	}()

	sigChan := make(chan os.Signal)
	signal.Notify(sigChan, os.Interrupt)
	signal.Notify(sigChan, os.Kill)

	sig := <- sigChan
	logger.Println("Received signal, graceful shutdown", sig)

	tc, _ := context.WithTimeout(context.Background(), 10 * time.Second)

	server.Shutdown(tc)
}
