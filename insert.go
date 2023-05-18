package main

import (
	"fmt"

	supa "github.com/nedpals/supabase-go"
)

type Country struct {
	ID      int    `json:"id"`
	Name    string `json:"name"`
	Capital string `json:"capital"`
}

func main() {
	supabaseUrl := "<SUPABASE-URL>"
	supabaseKey := "<SUPABASE-KEY>"
	supabase := supa.CreateClient(supabaseUrl, supabaseKey)

	row := Country{
		ID:      5,
		Name:    "Germany",
		Capital: "Berlin",
	}

	var results []Country
	err := supabase.DB.From("countries").Insert(row).Execute(&results)
	if err != nil {
		panic(err)
	}

	fmt.Println(results) // Inserted rows
}
