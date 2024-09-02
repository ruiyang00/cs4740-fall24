// All files start with a package declaration
package main

// Import statements, one package on each line
import (
      "fmt"
      "lec1"
)

// Main method will be called when the Go executable is run
func main() {
      fmt.Println("Hello world!")
      lec1.Basic()
      lec1.Add(1, 2)
      //divide(3, 4)
      //loops()
      //slices()
      //maps()
      //sharks()
}

