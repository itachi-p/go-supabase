package main

import (
	"context"
	"fmt"

	supa "github.com/nedpals/supabase-go"
)

func main() {
	supabaseUrl := "https://bnbfhykebobfqjniuxfg.supabase.co"
	supabaseKey := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpc3MiOiJzdXBhYmFzZSIsInJlZiI6ImJuYmZoeWtlYm9iZnFqbml1eGZnIiwicm9sZSI6ImFub24iLCJpYXQiOjE2ODQyNDgxNzAsImV4cCI6MTk5OTgyNDE3MH0.zTap-sSV30qDlzIzjeh4aA12iZJS07z-4fLCyiSFjuM"
	supabase := supa.CreateClient(supabaseUrl, supabaseKey)

	ctx := context.Background()
	user, err := supabase.Auth.SignUp(ctx, supa.UserCredentials{
		Email:    "itachip38@gmail.com",
		Password: "2zGbWBwfJKdHjG9q",
	})
	if err != nil {
		panic(err)
	}

	fmt.Println(user)

	var results map[string]interface{}
	err = supabase.DB.From("users").Select("*").Single().Execute(&results)
	if err != nil {
		panic(err)
	}

	fmt.Println(results) // Selected rows
}
