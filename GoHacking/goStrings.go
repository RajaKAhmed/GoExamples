package main
import (
	"fmt"
)
func main () {
	fmt.Println("Please enter your name")
	var name string
	fmt.Scanln(&name)
	for i, c := range name {
	fmt.Printf("%d: %s\n", i, string(c))}	
	fmt.Printf("Hi, %s! I'm Go!\n", name)
	fmt.Println("Your name is ", len(name), " chars long")

}