package main

import (
	"fmt"

	supa "github.com/nedpals/supabase-go"
)

func main() {
	supabaseUrl := "<SUPABASE-URL>"
	supabaseKey := "<SUPABASE-KEY>"
	supabase := supa.CreateClient(supabaseUrl, supabaseKey)

	var results map[string]interface{}
	err := supabase.DB.From("countries").Select("*").Single().Execute(&results)
	if err != nil {
		panic(err)
	}

	fmt.Println(results) // Selected rows
}
