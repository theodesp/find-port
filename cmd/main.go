package main

import (
	"fmt"
	"flag"
	"math/rand"
	"time"
	"os"
	"github.com/theodesp/find-port"
	"log"
)

var (
	port int
	version bool
)
func init()  {
	rand.Seed(time.Now().Unix())
	flag.BoolVar(&version, "version", false, "output version and exit")
}


func main()  {
	flag.Parse()
	if version {
		fmt.Println(findport.AppVersion)
		os.Exit(0)
	}

	openPort, err := findport.DetectOpenPort()
	if err != nil {
		log.Fatalf("Get available port failed with %v", err)
	}

	fmt.Printf("Found available port at: %v\n", openPort)
}
