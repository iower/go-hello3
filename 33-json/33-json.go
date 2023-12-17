package main

import (
	"encoding/json"
	"fmt"
	"log"
)

type User struct {
	Id         int
	Name       string
	Occupation string
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
}
