package main

import (
	"fmt"

	supa "github.com/nedpals/supabase-go"
)

type Country struct {
	Name    string `json:"name"`
	Capital string `json:"capital"`
}

func main() {
	supabaseUrl := "<SUPABASE-URL>"
	supabaseKey := "<SUPABASE-KEY>"
	supabase := supa.CreateClient(supabaseUrl, supabaseKey)

	row := Country{
		Name:    "France",
		Capital: "Paris",
	}

	var results map[string]interface{}
	err := supabase.DB.From("countries").Update(row).Eq("id", "5").Execute(&results)
	if err != nil {
		panic(err)
	}

	fmt.Println(results) // Updated rows
}
