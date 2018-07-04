package main

import "fmt"

type person struct {
	name string
	age int
}

func main(){
        //first demo array of strings
	var a[2] string

	a[0] = "Isaac"
	a[1] = "Lem"

	fmt.Println(a[0], a[1])
	fmt.Println(a)

	//second demo array of integers
	var b[2] int

	b[0] = 1
	b[1] = 5
	fmt.Println(b)
	
	
	//third demo calling a function
	fmt.Println(add(b[0], b[1]))
	
	//fouth demo structs
	st := person{name: "Isaac", age: 29}
	fmt.Println("Name: ", st.name)
	fmt.Println("Age: ", st.age)

}


func add(num1, num2 int) int {
	var result int
	result = num1 + num2
	return result
}