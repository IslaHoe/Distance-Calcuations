package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

type Users struct {
	Users []User
}

type User struct {
	Latitude  float64 `json:"latitude,string"`
	UserID    int
	Name      string
	Longitude float64 `json:"longitude,string"`
}

func main() {
	// Open our jsonFile
	jsonFile, err := os.Open("customers.json")
	// if we os.Open returns an error then handle it
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("Successfully customers.json")
	// defer the closing of our jsonFile so that we can parse it later on
	defer jsonFile.Close()

	// read our opened xmlFile as a byte array.
	byteValue, _ := ioutil.ReadAll(jsonFile)

	// we initialize our Users array
	var users Users

	fmt.Println(len(users.Users))
	// we unmarshal our byteArray which contains our
	// jsonFile's content into 'users' which we defined above
	json.Unmarshal(byteValue, &users)

	// we iterate through every user within our users array and
	// print out the user Type, their name, and their facebook url

	fmt.Println("Lat: " + users.Users[0].Latitude)
	fmt.Println("Lng: " + users.Users[0].Longitude)
	fmt.Println("Name: " + users.Users[0].Name)
	fmt.Println("ID: ", users.Users[0].UserID)

}
