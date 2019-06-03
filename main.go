package main

import (
	"log"
	"os"

	"github.com/betorvs/bigestresponsetimeicmp/config"
	"github.com/betorvs/bigestresponsetimeicmp/usecase"
)

func init() {
	log.SetOutput(os.Stdout)
}

func main() {
	//println("works: " + config.DestinationHost)
	usecase.CalculateBigestResponseTime(config.DestinationHost)

}
