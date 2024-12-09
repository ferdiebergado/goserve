package main

import (
	"errors"
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/ferdiebergado/goserve/pkg"
)

const defaultPort = "8888"
const defaultAddress = "localhost"
const defaultDir = "."
const readTimeout = 10
const writeTimeout = 10
const idleTimeout = 60

func main() {
	address := flag.String("a", defaultAddress, "network address to bind to")
	port := flag.String("p", defaultPort, "port to listen on")

	flag.Usage = func() {
		fmt.Fprintf(flag.CommandLine.Output(), "Usage: %s [flags] <path>\n", os.Args[0])
		fmt.Println("Flags:")
		flag.PrintDefaults()
		fmt.Println("\nArguments:")
		fmt.Println(`  <path>   Path to serve (default ".")`)
	}

	flag.Parse()

	args := flag.Args()

	dir := defaultDir

	if len(args) > 0 {
		dir = args[len(args)-1]
	}

	srv := http.Server{
		Addr:         *address + ":" + *port,
		Handler:      pkg.RequestLogger(pkg.DisableCache(http.FileServer(http.Dir(dir)))),
		ReadTimeout:  readTimeout * time.Second,
		WriteTimeout: writeTimeout * time.Second,
		IdleTimeout:  idleTimeout * time.Second,
	}

	log.Printf("HTTP Server listening at %s:%s...\n", *address, *port)

	if err := srv.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
		fmt.Fprintf(os.Stderr, "error listening and serving: %s\n", err)
	}
}
