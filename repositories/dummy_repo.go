package repositories

import "fmt"

func GetUserGreeting(name string) string {
	return fmt.Sprintf("Welcome, %s!", name)
}
