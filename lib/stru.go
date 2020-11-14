package lib

import "fmt"

type Woman struct {
	Name string
	Age int
}
func HoldOn()  {
	fmt.Println("women I hold on")
}

func (*Woman)WomenSayHello()  {
	HoldOn()
}
