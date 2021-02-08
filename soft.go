package goproj

import "fmt"

func Hai(name string) string {
	message := fmt.Sprintf("Hai %v how are you?", name)
	return message
}
