package auth

import "fmt"

func UserAuthentication(token string) string {
	fmt.Println("Using user authentication")
	return fmt.Sprintf("User %s", token)
}
