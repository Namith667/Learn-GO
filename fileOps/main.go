package main

import (
	//"fmt"
	"io"
	"log"
	"os"
)



func main(){
	//var fileName []string 
	// fileName := os.Args[1]
	// fmt.Println(fileName[1])
	file, err := os.Open(os.Args[1])
	if err != nil {
		log.Println("Error opening file:", err)
		return
	}
	defer file.Close()
	
	io.Copy(os.Stdout, file)
	

}