package main

import (
	"os"
	"log"
)

func main() {
	for i := 0; i < 5; i++ {
		f, err := os.Open("/home/xuhao/test/to/file")
		if err != nil {
			log.Fatal(err)
		}
		defer f.Close()
	}
}