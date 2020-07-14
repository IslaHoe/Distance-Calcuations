package main

import (
	"flag"
	"fmt"
	"os"
	"sort"
)

var (
	latitude   = flag.Float64("lat", 53.339428, "Target latitude")
	longitude  = flag.Float64("long", -6.257664, "Target longitude")
	radius     = flag.Float64("radius", 100, "Target radius")
	outputFile = flag.String("output", "output.txt", "Write results to given filename")
	inputFile  = flag.String("input", "customers.txt", "Read customer list from given filename")
)

func main() {
	flag.Parse()

	customers, err := fetchCustomers(*inputFile)
	if err != nil {
		fmt.Printf("Error reading from input file %s: %s\n", *inputFile, err)
		os.Exit(1)
	}

	filteredCustomers, err := filterByDistance(customers, *latitude, *longitude, *radius)
	if err != nil {
		fmt.Printf("Error filtering customers: %s\n", err)
		os.Exit(1)
	}

	sort.Sort(ByID(filteredCustomers))

	if err := writeToFile(filteredCustomers, *outputFile); err != nil {
		fmt.Printf("Error writing file %s: %s\n", *outputFile, err)
		os.Exit(1)
	}
}
