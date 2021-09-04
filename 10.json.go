package main

import(
	"fmt"
	"encoding/json"
)

// Profile declares `Profile` structure
type Profile struct {
	Username string `json:"uname,omitempty"`
	Followers int `json:"followers,omitempty,string"`
}

// Student declares `Student` structure
type Student struct {
	FirstName string `json:"fname"` // `fname` as field name
	LastName string `json:"lname,omitempty"` // discard if value is empty
	Email string `json:"-"` // always discard
	Age int `json:"-,"` // `-` as field name
	IsMale bool `json:",string"` // keep original field name, coerce to a string
	Profile Profile `json:""` // no effect
}

func main() {

	// define `john` struct (pointer)
	john := &Student{
		FirstName: "John",
		LastName: "", // empty
		Age: 21,
		Email: "john@doe.com",
		Profile: Profile{
			Username: "johndoe91",
			Followers: 1975,
		},
	}

	// encode `john` as JSON
	johnJSON, _ := json.MarshalIndent( john, "", "  " )

	// print JSON string
	fmt.Println( string(johnJSON) )
}