package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	_, err := os.Open("nothing.txt")
	if err != nil {
		fmt.Println(err)
		log.Println("error opening file", err) // date and time of each logged message
		log.Fatal(err)
		//panic(err)
	}
}
