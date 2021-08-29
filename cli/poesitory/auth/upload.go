package auth

import "fmt"

func UploadAuthentication(token string) string {
	fmt.Println("Using upload token authentication")
	return fmt.Sprintf("Plugin %s", token)
}
