package main
import (
    "fmt"
    "os"
)

func main() {
f, err := os.Open("countline1.txt")
if err != nil {
	fmt.Println(" File Not Found")
	return
}
fmt.Println(f.Name(), "success")
f.Close()	
}