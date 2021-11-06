package main

import (
   "fmt"
   "log"
   "os"
)

var (
   f os.FileInfo
   err error
)

func main() {
   // Stat returns file info. It will return
   // an error if there is no file.
   f, err = os.Stat("Bash-Beginners-Guide.pdf")
   if err != nil {
      log.Fatal(err)
   }
   fmt.Println("File name:", f.Name())
   fmt.Println("Size in bytes:", f.Size())
   fmt.Println("Permissions:", f.Mode())
   fmt.Println("Last modified:", f.ModTime())
   fmt.Println("Is Directory: ", f.IsDir())
   fmt.Printf("System interface type: %T\n", f.Sys())
   fmt.Printf("System info: %+v\n\n", f.Sys())
}