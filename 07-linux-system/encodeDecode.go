package main

import (
	"encoding/json"
	"fmt"
)

type UseAll struct {
	Name    string `json:"username"`
	Surname string `json:"surname"`
	Year    int    `json:"created"`
}

func main() {
	useall := UseAll{
		Name:    "Mike",
		Surname: "Tsoukalos",
		Year:    2023,
	}
	// Convert Structure to JSON record with fields
	t, err := json.Marshal(&useall)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("Value %s\n", t) // Value {"username":"Mike","surname":"Tsoukalos","created":2023}
	}

	str := `{"username": "M.", "surname": "Ts", "created":2024}`
	// Convert string into a byte slice
	jsonRecord := []byte(str)
	temp := UseAll{}
	err = json.Unmarshal(jsonRecord, &temp)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("Data type: %T with value %v\n", temp, temp) // Data type: main.UseAll with value {M. Ts 2024}
	}
}
