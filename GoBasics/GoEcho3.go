// Echo3 prints its arguments using Join function.
package main
import (
"fmt"
"os"
)
func main() {
	// fmt.Println(os.Args[0:])
	var s string
	fmt.Println(s.Join(os.Args[1:], " "))
}
