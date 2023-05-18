package main

import (
	"fmt"
	"time"

	supa "github.com/nedpals/supabase-go"
)

type User struct {
	ID        int
	UUID      string
	Name      string
	Email     string
	PassWord  string
	CreatedAt time.Time
	//Todos     []Todo
}

func main() {
	supabaseUrl := "https://bnbfhykebobfqjniuxfg.supabase.co"
	supabaseKey := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpc3MiOiJzdXBhYmFzZSIsInJlZiI6ImJuYmZoeWtlYm9iZnFqbml1eGZnIiwicm9sZSI6ImFub24iLCJpYXQiOjE2ODQyNDgxNzAsImV4cCI6MTk5OTgyNDE3MH0.zTap-sSV30qDlzIzjeh4aA12iZJS07z-4fLCyiSFjuM"
	supabase := supa.CreateClient(supabaseUrl, supabaseKey)

	var results map[string]interface{}
	err := supabase.DB.From("users").Select("*").Single().Execute(&results)
	if err != nil {
		panic(err)
	}

	fmt.Println(results) // Selected rows
}
