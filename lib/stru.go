package lib

import "fmt"

type Woman struct {
	name string
	age int
}
func HoldOn()  {
	fmt.Println("women I hold on")
}

func (*Woman)WomenSayHello()  {
	HoldOn()
}
