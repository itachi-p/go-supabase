package main

import (
	"context"
	"fmt"

	supa "github.com/nedpals/supabase-go"
)

func main() {
	supabaseUrl := "<SUPABASE-URL>"
	supabaseKey := "<SUPABASE-KEY>"
	supabase := supa.CreateClient(supabaseUrl, supabaseKey)

	ctx := context.Background()
	user, err := supabase.Auth.SignIn(ctx, supa.UserCredentials{
		Email:    "example@example.com",
		Password: "password",
	})
	if err != nil {
		panic(err)
	}

	fmt.Println(user)
}
