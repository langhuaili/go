package main

import (
	"GoStudy/lib"
	"fmt"
)

type human struct {
	lib.Woman
}
func main()  {
	//var kacy human
	//kacy.WomenSayHello()
	getJson()
}
func getJson()  {
	arr :=lib.JsonDeFomat()
	fmt.Println(arr)
}
