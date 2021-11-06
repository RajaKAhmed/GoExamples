package main
import (
	"fmt"
	)
func main () {
	/*myArr := []string{"aik", "dau", "teen"}
	fmt.Println(myArr)
	*/
	mySlice := []int{}
	fmt.Println(mySlice)
	i := 0
	for i <= 20 {
	mySlice = append(mySlice, i)
	i ++
	fmt.Println(mySlice)
}

	fmt.Println("New MySlice:\n", mySlice)
}