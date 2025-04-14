package main

import (
	"context"
	"flag"
	"log"
	"net/http"
	"os"
	"os/signal"
	"strings"
	"time"

	"github.com/denisdefreyne/dflag/server"
	"github.com/denisdefreyne/dflag/stores"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/collectors"
)

const (
	DEFAULT_PORT                         = "3000"
	DEFAULT_FEATURE_FLAGS_DATA_FILE_PATH = "data/feature_flags.hcl"
)

func getAddr() string {
	var addr string

	port := os.Getenv("PORT")
	if port == "" {
		port = DEFAULT_PORT
	}
	if strings.Contains(port, ":") {
		addr = port
	} else {
		addr = ":" + port
	}

	return addr
}

func main() {
	var check = flag.Bool("check", false, "check configuration and exit")
	flag.Parse()

	// Read feature flags
	featureFlagsPath := os.Getenv("FEATURE_FLAGS_DATA_FILE_PATH")
	if featureFlagsPath == "" {
		featureFlagsPath = DEFAULT_FEATURE_FLAGS_DATA_FILE_PATH
	}
	featureFlagsStore := stores.ReadFeatureFlags(featureFlagsPath)

	// Stop if only checking, not running
	if *check {
		log.Println("check ok")
		return
	}

	// Make server
	r := prometheus.NewRegistry()
	r.MustRegister(collectors.NewGoCollector())
	r.MustRegister(collectors.NewBuildInfoCollector())
	r.MustRegister(collectors.NewProcessCollector(collectors.ProcessCollectorOpts{}))
	server.RegisterMetrics(r)
	server := server.NewServer(getAddr(), r, &featureFlagsStore)

	// Start
	go func() {
		if err := server.ListenAndServe(); err != http.ErrServerClosed {
			panic(err)
		}
	}()
	log.Printf("Listening on %s\n", server.Addr)

	// Wait for shutdown signal
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)
	sig := <-quit

	// Shut down
	log.Printf("Shutting down (reason: %v)\n", sig)
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := server.Shutdown(ctx); err != nil {
		panic(err)
	}
}
