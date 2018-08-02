package main

import (
	"io/ioutil"
	"fmt"
)

type Address struct {
	adId int
}

type VCard struct {
	name string
	id *Address
	data string
	funimage []byte
}

func main(){

	Tom:=VCard{"Tom",&Address{1001},"2018-7-31",readimage()}
	fmt.Println(Tom)
}

func readimage() []byte{
	b,err:=ioutil.ReadFile("test.png")
	if err!=nil {
		return nil
	}
	return b
}