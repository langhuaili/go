package main

import "fmt"

func main()  {

	var xiaoyaong student= student{
		people{180,"xiaoli",18},
		1,
		"上海中学",
		[]float32{1.1,2,3,4},
		favour{"运动":"篮球","电影":"泰坦尼克"},
	}
	fmt.Println(xiaoyaong.geAge(xiaoyaong))
}