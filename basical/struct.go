package main

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



/*给student结构挂载一个属性方法*/
func (student) geAge (studentobj student) int  {
	return studentobj.age
}
func (student) getFavour (studentobj student) favour  {
	return studentobj.favour
}