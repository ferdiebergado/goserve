package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/ferdiebergado/goserve/pkg"
)

func main() {
	const defaultPort = "8888"

	address := flag.String("a", "localhost", "network address to bind to")
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

	dir := "."

	if len(args) > 0 {
		dir = args[len(args)-1]
	}

	log.Printf("HTTP Server listening at %s:%s...\n", *address, *port)
	log.Fatal(http.ListenAndServe(*address+":"+*port, pkg.RequestLogger(http.FileServer(http.Dir(dir)))))
}
