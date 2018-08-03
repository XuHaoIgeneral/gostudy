package main

import (
	"errors"
	"fmt"
	"math"
)

var errNotFOund error =errors.New("Not found error")

type PathError struct {
	Op string    // "open", "unlink", etc.
	Path string  // The associated file.
	Err error  // Returned by the system call.
}


func main()  {
	fmt.Println("error: %v",errNotFOund)

	testerr:=&PathError{"open","/root",errNotFOund}
	fmt.Println(testerr)
	if _, err := Sqrt(-1); err != nil {
		fmt.Printf("Error: %s\n", err)
	}
}

func Sqrt(f float64) (float64, error) {
	if f < 0 {
		return 0, errors.New ("math - square root of negative number")
	}
	re:=math.Sqrt(f*f)
	return re,nil
}

func (e *PathError) String() string {
	return e.Op + " " + e.Path + ": "+ e.Err.Error()
}
