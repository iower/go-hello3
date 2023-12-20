package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"time"
)

type User struct {
	Id         int
	Name       string
	Occupation string
}

type User2 struct {
	name       string
	occupation string
	born       string
}

type Astronaut struct {
	Name  string
	Craft string
}

type people struct {
	Number  int
	People  []Astronaut
	Message string
}

func main() {
	u1 := User{1, "John Doe", "gardener"}
	jsonData, err := json.Marshal(u1)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(jsonData))

	users := []User{
		{Id: 2, Name: "Roger", Occupation: "driver"},
		{Id: 3, Name: "Lucy", Occupation: "teacher"},
	}
	jsonData2, err := json.Marshal(users)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(jsonData2))

	var u11 User
	err2 := json.Unmarshal(jsonData, &u11)
	if err2 != nil {
		log.Fatal(err2)
	}
	fmt.Println("Unmarshal", u11)

	var users2 []User
	err3 := json.Unmarshal(jsonData2, &users2)
	if err2 != nil {
		log.Fatal(err3)
	}
	fmt.Println("Unmarshal", users2)

	birds := map[string]interface{}{
		"sounds": map[string]string{
			"pigeon": "coo",
			"eagle":  "squak",
		},
	}
	data, err := json.MarshalIndent(birds, "", "    ")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(data))

	file, err := os.Open("data.json")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	data2, err2 := ioutil.ReadAll(file)
	if err2 != nil {
		log.Fatal(err2)
	}
	fmt.Println(string(data2))

	var result []User2
	jsonErr := json.Unmarshal(data2, &result)
	if jsonErr != nil {
		log.Fatal(jsonErr)
	}
	fmt.Println(result)

	//

	url := "http://api.open-notify.org/astros.json"
	var client = http.Client{
		Timeout: time.Second * 10,
	}
	res, err := client.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	fmt.Println("body: ", string(body))

	if err != nil {
		log.Fatal(err)
	}
	astros := people{}
	jsonErr2 := json.Unmarshal(body, &astros)
	if jsonErr2 != nil {
		log.Fatal(jsonErr2)
	}
	fmt.Println("astros:", astros)
}
