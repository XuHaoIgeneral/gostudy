package main

import "fmt"

type Person struct {

}

var people *Person

func main()  {
	fmt.Printf("from %V",people)
}