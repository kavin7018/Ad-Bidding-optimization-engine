package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"personal-documents/zocket/app/engine"
	"personal-documents/zocket/app/healthmetrics"
	"personal-documents/zocket/app/mock"
	"time"

	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	exitChan := make(chan bool)

	// initialize self health metrics
	healthmetrics.Init()
	go func() {
		http.Handle("/metrics", promhttp.Handler())
		log.Fatal(http.ListenAndServe(":2112", nil))
	}()

	// Adding timer pulsar server to be up and running
	time.Sleep(20 * time.Second)
	go mock.Start(ctx, exitChan)

	// Adding timer to create producer topics in mock server
	time.Sleep(10 * time.Second)
	go engine.Start(ctx, exitChan)

	if <-exitChan {
		fmt.Println("Shutting down the server...")
		cancel()
	}

}
