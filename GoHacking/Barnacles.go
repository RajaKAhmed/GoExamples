package main

import (
    "errors"
	"fmt"
	"time"
)

func main() {
    err := errors.New("barnacles")
	fmt.Println("Sammy says:", err)
	
	//using time 
	err1 := fmt.Errorf("error occurred at: %v", time.Now())
    fmt.Println("An error happened:", err1)
}