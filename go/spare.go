package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"log"
	"os"
)

func main() {
	file, err := os.Open("customers.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		fmt.Println("hello")

		fmt.Println(scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}



/*package main

import (
	"encoding/json"
	"fmt"
)

type Book struct {
	Title  string `json:"title"`
	Author Author `json:"author"`
}

type Author struct {
	Name     string `json:"name"`
	Age      int    `json:age`
	Devloper bool   `json:"is_devloper"`
}

func main() {
	fmt.Println("hello world")

	author := Author{Name: "Ellio Forbes", Age: 25, Devloper: true}
	book := Book{Title: "Learning Concurancy", Author: author}

	fmt.Printf("%+v\n", book)

	//json.Marshall returns a byte array (which is why its been called a byte array)
	byteArray, err := json.MarshalIndent(book, "", " ")
	if err != nil {
		fmt.Println(err)
	}

	//because .Marshall returns a byte array we want to cast as a string
	fmt.Println(string(byteArray))
}

package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	file, err := os.Open("customers.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		fmt.Println("hello")

		fmt.Println(scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}







type SensorReading struct {
	Name       string `json:"name"`
	Capacity   int    `json:"capacity"`
	Time       string `json:"time"`
	Infomation Info   `json:"info"`
}

type Info struct {
	Description string `json:"desc"`
}
*/

/*
type Book struct {
	Title  string `json:"title"`
	Author Author `json:"author"`
}

type Author struct {
	Name     string `json:"name"`
	Age      int    `json:age`
	Devloper bool   `json:"is_devloper"`
}

func main() {
	fmt.Println("hello world")

	author := Author{Name: "Ellio Forbes", Age: 25, Devloper: true}
	book := Book{Title: "Learning Concurancy", Author: author}

	fmt.Printf("%+v\n", book)

	//json.Marshall returns a byte array (which is why its been called a byte array)
	byteArray, err := json.MarshalIndent(book, "", " ")
	if err != nil {
		fmt.Println(err)
	}

	//because .Marshall returns a byte array we want to cast as a string
	fmt.Println(string(byteArray))
}


*/

/*
	fmt.Println("hello world")

	content, err := ioutil.ReadFile("test.txt")
	fmt.Printf("File contents: %s", content)

	jsonString := `{"name": "battery sensor", "capacity": 40, "time": "2019-01-21T19:07:28Z", "info": { "desc": "a sensor reading"}}`
	var reading map[string]interface{}
	errr := json.Unmarshal([]byte(jsonString), &reading)
	if errr != nil {
		fmt.Println(err)
	}

	fmt.Printf("%+v\n", reading)

	resp, err := http.Get("https://s3.amazonaws.com/intercom-take-home-test/customers.txt")
	if err != nil {
		// handle err
	}
	fmt.Println(resp)
	defer resp.Body.Close()*/
