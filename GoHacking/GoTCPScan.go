package main

import (
    "fmt"
    "net"
)

func main() {
    _, err := net.Dial("tcp", "192.168.17.128:80")
     if err != nil {
        
        fmt.Println("Connection Unsuccessful")
        return
                    }

        fmt.Println("Connection successful")  
}