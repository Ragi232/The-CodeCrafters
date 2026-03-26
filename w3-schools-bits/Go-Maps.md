## GO maps 

A Go map is a built-in data type in the Go lauguage that stores data in kay value pairs.it is like a dictionary where you use a kay to get a value.

## E.g 

package main
import ("fmt")

func main() {
  var a = map[string]string{"brand": "Ford", "model": "Mustang", "year": "1964"}
  b := map[string]int{"Oslo": 1, "Bergen": 2, "Trondheim": 3, "Stavanger": 4}

  fmt.Printf("a\t%v\n", a)
  fmt.Printf("b\t%v\n", b)
}