package main


import (
	"fmt"
	//"leetcode/link"
)

type Category struct{
	Name string
}

type Cat struct{
	Name string
	*Category
}

func (cat *Cat)SetName(name string)  {
	cat.Name = name
}
func (cat Cat)SetName2(name string)  {
	cat.Name = name
}


func main()  {
	cat := Cat{Name:"kk"}
	fmt.Println(cat.Name)
	cat.SetName2("123")
	fmt.Println(cat.Name)
	cat.SetName("123")
	fmt.Println(cat.Name)
	var a int
	//b:=1
	fmt.Println(&a)


	cd := make([]int,10)
	cd[0] =1

	cd2 := cd

	cd2[0] =2
	
	fmt.Println(cd[0])
}