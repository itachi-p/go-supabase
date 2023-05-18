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
	user, err := supabase.Auth.SignIn(ctx, supa.UserCredentials{
		Email:    "itachip38@gmail.com",
		Password: "qnB795jr@gfD#5V",
	})
	if err != nil {
		panic(err)
	}

	fmt.Println(user)
}
