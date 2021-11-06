package main

import (
    "errors"
	"fmt"
	"time"
)

func boom() error {
	//return errors.New("Lulla")
	return fmt.Errorf(time.Now(), errors.New("lunndd"))
}

func main() {
    err := boom()

    if err != nil {
        fmt.Println("An error occurred:", err)
        return
    }
 
	fmt.Println("Anchors away!")
}