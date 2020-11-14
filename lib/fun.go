package lib

import "fmt"

type man struct {
	name string
	age int
}
func (*man)holdOn()  {
	fmt.Println("Man hold on")
}

func SayHello()  {
	fmt.Printf("hello")
}
