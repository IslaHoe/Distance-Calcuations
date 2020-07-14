package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"math"
	"os"
	"strconv"
)

//Customer type based on paramaters in customers.txt
type Customer struct {
	Latitude  string `json:"latitude"`
	UserID    int    `json:"user_id"`
	Name      string `json:"name"`
	Longitude string `json:"longitude"`
}

// ByID implements sort.Interface for []Customer based on the ID field.
type ByID []Customer

// functions to sort the customers by their ID's
func (a ByID) Len() int           { return len(a) }
func (a ByID) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a ByID) Less(i, j int) bool { return a[i].UserID < a[j].UserID }

//convert lat and lng coords in degrees to radians
func convertToRadians(degrees float64) float64 {
	radians := degrees * (math.Pi / 180)
	return radians
}

//calcuate the distance between the office at (53.339428,-6.257664) and the customer coords
func calculateDistance(lat1, lng1, lat2, lng2 float64) float64 {
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

//fetch customoers from customers.txt and place in array for ease of use
func fetchCustomers(filename string) ([]Customer, error) {
	var CustomerArray []Customer

	file, err := os.Open(filename)
	if err != nil {
		return CustomerArray, err
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)
	if err != nil {
		return CustomerArray, err
	}

	var customer Customer
	for scanner.Scan() {
		customerLn := scanner.Text()
		if err := json.Unmarshal([]byte(customerLn), &customer); err != nil {
			return CustomerArray, err
		}

		CustomerArray = append(CustomerArray, customer)
	}
	return CustomerArray, nil
}

// remove customers who live outside the 100k radius
// sourceLat and sourceLng are the coordnates from where
func filterByDistance(customers []Customer, sourceLat, sourceLng, targetDistance float64) ([]Customer, error) {
	var filteredCustomers []Customer

	for _, customer := range customers {
		lat1, err := strconv.ParseFloat(customer.Latitude, 64)
		if err != nil {
			return filteredCustomers, err
		}

		lng1, err := strconv.ParseFloat(customer.Longitude, 64)
		if err != nil {
			return filteredCustomers, err
		}

		distance := calculateDistance(lat1, lng1, sourceLat, sourceLng)
		if distance < targetDistance {
			filteredCustomers = append(filteredCustomers, customer)
		}
	}

	return filteredCustomers, nil
}

// write all remaining customers to output.txt
func writeToFile(customers []Customer, filename string) error {

	f, err := os.Create(filename)
	if err != nil {
		return err
	}

	for _, customer := range customers {
		newLine := fmt.Sprintf("%d: %s", customer.UserID, customer.Name)
		_, err = fmt.Fprintln(f, newLine)
		if err != nil {
			return err
		}
	}

	return f.Close()
}
