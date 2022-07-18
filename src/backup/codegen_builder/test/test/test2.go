package main

import (
	"encoding/json"
	"fmt"
)

type Enum struct {
	Name       string
	CustomType string
}

func main() {
	jsonStr := "[{\"owner_type\":[\"rental\",\"buying\",\"leasing\",\"\"]}]"

	var results []map[string][]string

	json.Unmarshal([]byte((jsonStr)), &results)

	var enums []Enum

	for _, m := range results {
		for name, values := range m {
			for _, v := range values {
				enum := Enum{Name: name, CustomType: v}
				enums = append(enums, enum)
			}
		}
	}

	fmt.Printf("results: %v\n", enums)
}
