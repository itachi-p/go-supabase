package main

import (
	"fmt"
	"time"

	supa "github.com/nedpals/supabase-go"
)

type User struct {
	ID        int       `json:"id"`
	UUID      string    `json:"uuid"`
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	PassWord  string    `json:"password"`
	CreatedAt time.Time `json:"created_at"`
	// Todos     []Todo
}

func main() {
	supabaseUrl := "https://bnbfhykebobfqjniuxfg.supabase.co"
	supabaseKey := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpc3MiOiJzdXBhYmFzZSIsInJlZiI6ImJuYmZoeWtlYm9iZnFqbml1eGZnIiwicm9sZSI6ImFub24iLCJpYXQiOjE2ODQyNDgxNzAsImV4cCI6MTk5OTgyNDE3MH0.zTap-sSV30qDlzIzjeh4aA12iZJS07z-4fLCyiSFjuM"
	supabase := supa.CreateClient(supabaseUrl, supabaseKey)

	row := User{
		//		ID:      1,
		UUID:      "UUID0001",
		Name:      "supa_testuser1",
		Email:     "sptest1@example.com",
		PassWord:  "sptest1",
		CreatedAt: time.Now(),
	}

	var results []User
	err := supabase.DB.From("users").Insert(row).Execute(&results)
	if err != nil {
		panic(err)
	}

	fmt.Println(results) // Inserted rows
}
