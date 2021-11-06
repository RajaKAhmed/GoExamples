package main

import (
    "fmt"
	"net"
)

func main() {
    for i := 20; i <= 300; i++ {
        address := fmt.Sprintf("scanme.nmap.org:%d", i)
        conn, err := net.Dial("tcp", address)
        if err != nil {
            // port is closed or filtered.
			fmt.Printf(address, "IP Address Not found")
        }
        conn.Close()
        fmt.Printf("%d open\n", i)
    }
}