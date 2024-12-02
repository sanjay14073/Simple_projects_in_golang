package main

import (
	"fmt"
	"go/format"
	"os"
);

func main() {

	data, err := os.ReadFile("index.js")
	if err!=nil{
		fmt.Printf("There is an err",err);
	}
	formatedCode,err:=format.Source(data)
	if err!=nil{
		fmt.Printf("Error while formatting this code")
	}
	os.WriteFile("index.js",formatedCode,os.ModePerm.Perm())

}