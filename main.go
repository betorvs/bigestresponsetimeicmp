package main

import (
	"log"
	"os"

	"github.com/betorvs/biggestresponsetimeicmp/config"
	"github.com/betorvs/biggestresponsetimeicmp/usecase"
)

func init() {
	log.SetOutput(os.Stdout)
}

func main() {
	//println("works: " + config.DestinationHost)
	usecase.CalculateBigestResponseTime(config.DestinationHost)

}
