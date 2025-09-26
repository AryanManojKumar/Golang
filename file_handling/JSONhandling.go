package main

import (
	"encoding/json"
	"fmt"
	"log"
)

type usertable struct {
	Name string `json:"name"`
	UID  int    `json:"UID"`
	Api  int    `json:"Api"`
}

func main() {

	you := usertable{
		Name: "user",
		UID:  123,
		Api:  420}

	data, err := json.Marshal(you)
	if err != nil {
		log.Fatal("couldnt encode the data")
	}

	indented, err := json.MarshalIndent(you, " ", " ")

	fmt.Println(string(data))
	fmt.Print(string(indented))

	var result map[string]interface{}

	err = json.Unmarshal(data, &result)
	if err != nil {
		log.Print("couldnt unmarshal", err)
	}
	fmt.Print(result["name"])

}
