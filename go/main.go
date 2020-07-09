package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
)

type Customer struct {
	Latitude  string `json:"latitude"`
	User_id   int    `json:"user_id"`
	Name      string `json:"name"`
	Longitude string `json:"longitude"`
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
	fmt.Println("dist: ", dist)

	return dist
}

func convertToRadians(degrees float64) float64 {
	radians := degrees * (math.Pi / 180)
	return radians
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	file, err := os.Open("customers.json")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	f, err := os.Create("output.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	for scanner.Scan() {

		customerJson := scanner.Text()
		var customer Customer
		json.Unmarshal([]byte(customerJson), &customer)
		lat1, _ := strconv.ParseFloat(customer.Latitude, 64)
		lng1, _ := strconv.ParseFloat(customer.Longitude, 64)
		distance := calcuateDistance(lat1, lng1)
		if distance < 100 {
			fmt.Printf("%f Kilometers\n", distance)

			f.WriteString(customer.Name)
			f.WriteString("\n")

		}
	}

}
