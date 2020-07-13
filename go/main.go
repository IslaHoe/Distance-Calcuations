package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"math"
	"os"
	"strconv"
)

type Customer struct {
	Latitude  string `json:"latitude"`
	UserID    int    `json:"user_id"`
	Name      string `json:"name"`
	Longitude string `json:"longitude"`
}

type CustomerJSON struct {
	Name string
	ID   int
}

//place error checking in seperate function to make the full code more consise
func check(e error) {
	if e != nil {
		fmt.Println(e)
		return
	}
}
func calcuateDistance(lat1, lng1 float64) float64 {
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

func convertToRadians(degrees float64) float64 {
	radians := degrees * (math.Pi / 180)

	return radians
}

func swap(a, b CustomerJSON) (CustomerJSON, CustomerJSON) {
	var temp CustomerJSON
	temp = a
	a = b
	b = temp

	return a, b
}

func quickSort(CustomerArray []CustomerJSON) []CustomerJSON {

	for j := 1; j < len(CustomerArray); j++ {
		for i := 1; i < len(CustomerArray); i++ {
			if CustomerArray[i-1].ID > CustomerArray[i].ID {
				CustomerArray[i-1], CustomerArray[i] = swap(CustomerArray[i-1], CustomerArray[i])
			}
		}

	}

	writeToFile(CustomerArray)
	return CustomerArray
}

func writeToFile(CustomerArray []CustomerJSON) {

	f, err := os.Create("output.txt")
	check(err)

	for i := 1; i < len(CustomerArray); i++ {
		newLine := CustomerArray[i]
		_, err = fmt.Fprintln(f, newLine)
		check(err)
	}

	err = f.Close()
	check(err)
}

func writeCustArray(CustomerArray []CustomerJSON) []CustomerJSON {
	file, err := os.Open("customers.json")
	check(err)
	defer file.Close()

	scanner := bufio.NewScanner(file)
	check(err)
	i := 0
	for scanner.Scan() {

		customerJSON := scanner.Text()
		var customer Customer
		json.Unmarshal([]byte(customerJSON), &customer)
		lat1, _ := strconv.ParseFloat(customer.Latitude, 64)
		lng1, _ := strconv.ParseFloat(customer.Longitude, 64)
		distance := calcuateDistance(lat1, lng1)
		if distance < 100 {
			customerInfo := CustomerJSON{Name: customer.Name, ID: customer.UserID}
			CustomerArray = append(CustomerArray, customerInfo)
			i++
		}
	}

	return CustomerArray

}

func main() {

	var CustomerArray []CustomerJSON
	CustomerArray = writeCustArray(CustomerArray)
	quickSort(CustomerArray)
}
