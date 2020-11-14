package main

import (
	"GoStudy/lib"
	)

type human struct {
	lib.Woman
}
func main()  {
	var kacy human
	kacy.WomenSayHello()
}
