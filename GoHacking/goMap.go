package main
import (
	"fmt"
)
func main() {
MyMap :=map[string]string{"FirstName":"Raja", "LastName":"Ahmed","FavColor":"Blue"}
fmt.Println(MyMap["FirstName"])
fmt.Println(MyMap["LastName"])
fmt.Println(MyMap["FavColor"])
fmt.Println(MyMap)
}