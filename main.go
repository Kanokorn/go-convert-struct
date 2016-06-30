package main

import (
	"encoding/json"
	"fmt"
)

type User struct {
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName,omitempty"`
}

func main() {
	data := []byte(`{"firstName":"Weerasak"}`)
	u := User{FirstName: "Test", LastName: "Test"}
	json.Unmarshal(data, &u)
	fmt.Println(u)

	u2 := User{FirstName: "Weerasak"}
	b, _ := json.Marshal(u2)
	fmt.Println(string(b))
}
