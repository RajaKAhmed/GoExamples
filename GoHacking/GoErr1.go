package hello
import (
"errors"
"fmt"
)
func Hello(name string) (string, error) {
	if name == "" {
	return "", errors.New("No name was given")
	}
	message := fmt.Sprintf("Hi, %v. Welcome!", name)
return message, nil
}