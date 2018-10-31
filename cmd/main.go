package main

import (
	"flag"
	"fmt"
	"log"
	"math/rand"
	"os"
	"time"

	"github.com/theodesp/find-port"
)

var (
	AppVersion = "dev"
	version    bool
)

func init() {
	rand.Seed(time.Now().Unix())
	flag.BoolVar(&version, "version", false, "output version and exit")
}

func main() {
	flag.Parse()
	if version {
		fmt.Println(AppVersion)
		os.Exit(0)
	}

	openPort, err := findport.DetectOpenPort()
	if err != nil {
		log.Fatalf("Get available port failed with %v", err)
	}

	fmt.Printf("Found available port at: %v\n", openPort)
}
