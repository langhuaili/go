package main

import "fmt"

/*结构体继承*/
type people struct {
	tall int
	name string
	age int
}

type student struct {
	people
	level int
	school string
	scores []float32
	favour favour
}

type favour map[string]string

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

/*给student结构挂载一个属性方法*/
func (student) geAge (studentobj student) int  {
	return studentobj.age
}
