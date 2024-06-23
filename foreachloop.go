package main

import "fmt"

func main() {
    // A slice of strings
    fruits := []string{"apple", "banana", "cherry"}

    // Using range to iterate over the slice
    for _, value := range fruits {
        fmt.Println( "Value:", value)
    }
}
