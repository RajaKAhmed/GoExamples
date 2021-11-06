package main
import (
	"fmt"
)
var x int
func main() {
	x = 1
	fmt.Println(x)
	x += 1 					/*or x = x + 1 */
	fmt.Println("NewX", x)
	f()
}
func f() {
	fmt.Println("New X+1", x)
}