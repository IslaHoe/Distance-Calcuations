package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"log"
	"math"
	"os"
	"sort"
	"strconv"
)

//Customer type based on paramaters in customers.txt
type Customer struct {
	Latitude  string `json:"latitude"`
	UserID    int    `json:"user_id"`
	Name      string `json:"name"`
	Longitude string `json:"longitude"`
}

//convert lat and lng coords to radians
func convertToRadians(degrees float64) float64 {
	radians := degrees * (math.Pi / 180)
	return radians
}

//calcuate the distance between the office at (53.339428,-6.257664) and the customer coords
func calculateDistance(lat1, lng1 float64) float64 {
	lat2 := 53.339428
	lng2 := -6.257664
	radius := 6373.0

	radlat1 := convertToRadians(lat1)
	radlat2 := convertToRadians(lat2)

	theta := float64(lng1 - lng2)
	radtheta := convertToRadians(theta)

	dist := math.Sin(radlat1)*math.Sin(radlat2) + math.Cos(radlat1)*math.Cos(radlat2)*math.Cos(radtheta)
	dist = math.Acos(dist)

	dist = dist * radius

	return dist
}

// calcuate the distance between two sets of coordnates
func checkDistance(customer Customer) float64 {
	lat1, _ := strconv.ParseFloat(customer.Latitude, 64)
	lng1, _ := strconv.ParseFloat(customer.Longitude, 64)
	distance := calculateDistance(lat1, lng1)
	return distance
}

//fetch customoers from customers.txt and place in array for ease of use
func fetchCustomers() (CustomerArray []Customer) {
	file, err := os.Open("customers.txt")
	if err != nil {
		log.Println(err)
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)
	if err != nil {
		log.Println(err)
	}

	var customer Customer
	for scanner.Scan() {
		customerLn := scanner.Text()
		json.Unmarshal([]byte(customerLn), &customer)
		CustomerArray = append(CustomerArray, customer)
	}
	return CustomerArray
}

// remove customers who live outside the 100k radius
func filterCustomer(customers []Customer) []Customer {

	var filteredCustomers []Customer
	for i := 1; i < len(customers); i++ {
		distance := checkDistance(customers[i])
		if distance < 100 {
			filteredCustomers = append(filteredCustomers, customers[i])
		}
	}

	return filteredCustomers

}

// write all remaining customers to output.txt
func writeToFile(customers []Customer) {

	f, err := os.Create("output.txt")
	if err != nil {
		log.Println(err)
	}

	for i := 0; i < len(customers); i++ {
		newLine := (strconv.Itoa(customers[i].UserID) + ": " + customers[i].Name)
		_, err = fmt.Fprintln(f, newLine)
		if err != nil {
			log.Println(err)
		}

	}

	err = f.Close()
	if err != nil {
		log.Println(err)
	}

}

// ByID implements sort.Interface for []Customer based on the ID field.
type ByID []Customer

// functions to sort the customers by their ID's
func (a ByID) Len() int           { return len(a) }
func (a ByID) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a ByID) Less(i, j int) bool { return a[i].UserID < a[j].UserID }

func main() {

	customers := fetchCustomers()
	customers = filterCustomer(customers)
	sort.Sort(ByID(customers))
	writeToFile(customers)
}
