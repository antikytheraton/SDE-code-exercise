package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/peterbourgon/ff/v3"

	"code_exercise/shipments"
)

var fs = flag.NewFlagSet("shipments", flag.ContinueOnError)

const usage = `Usage: 
-destinations    - flag to set the destinations list file
-drivers         - flag to set the driver's names list file
`

func main() {
	flag.Usage = func() {
		fmt.Fprint(os.Stderr, usage)
	}
	var (
		destinationsFile = fs.String("destinations", "", "file with list of destinations")
		driversFile      = fs.String("drivers", "", "file with list of drivers")
	)
	err := ff.Parse(fs, os.Args[1:])
	if err != nil {
		log.Fatal(err)
	}
	if *destinationsFile == "" || *driversFile == "" {
		flag.Usage()
		log.Fatal("you must supply both destinations and drivers file names")
	}
	destinations, err := readFile(*destinationsFile)
	if err != nil {
		log.Fatal(err)
	}
	drivers, err := readFile(*driversFile)
	if err != nil {
		log.Fatal(err)
	}
	assignations := shipments.AssignRoutes(destinations, drivers)
	for _, assignation := range assignations {
		fmt.Printf("%+v\n", assignation)
	}
}
