package main

import (
	"fmt" 
	"time"
)

type person struct {
	name string
	age int
}

func main(){
        fmt.Println("============First Demo===========")
        //first demo array of strings
	var a[2] string

	a[0] = "Isaac"
	a[1] = "Lem"
	
	fmt.Println(a[0], a[1])
	fmt.Println(a)

	fmt.Println("============Second Demo===========")
	//second demo array of integers
	var b[2] int

	b[0] = 1
	b[1] = 5
	fmt.Println(b)
	
	fmt.Println("============Third Demo===========")
	//third demo calling a function
	fmt.Println(add(b[0], b[1]))
	
	fmt.Println("============Fourth Demo===========")
	//fouth demo structs
	st := person{name: "Isaac", age: 29}
	fmt.Println("Name: ", st.name)
	fmt.Println("Age: ", st.age)
	
	fmt.Println("============Fifth Demo===========")
	//fifth demo using pointer, extension of fourth demo
	s := &st
	fmt.Println("Second Name: ", s.name)
	fmt.Println("Second Age: ", s.age)
	
	fmt.Println("============Sixth Demo===========")
	//sixth demo variable
	var sixthDemo_a string = "isaac to GO"
	fmt.Println(sixthDemo_a)
	
	var sixthDemo_b int = 2
	fmt.Println(sixthDemo_b)
	
	var sixthDemo_c = true
	fmt.Println(sixthDemo_c)
	
	fmt.Println("============Seventh Demo===========")
	//seventh demo: switch statement
	
	var seventhDemo_i = 2
	
	switch (seventhDemo_i == 2) { //parenthesis is not important
		case true:
			fmt.Println("Hi I'm in case 1")
		case false:
			fmt.Println("Hi I'm in case 2")
		default:
			fmt.Println("i not found")
	}
	
	fmt.Println("============Eighth Demo===========")
	//Eighth Demo: getting time
	eighthDemo_t := time.Now()
	
	fmt.Println(eighthDemo_t)
}


func add(num1, num2 int) int {
	var result int
	result = num1 + num2
	return result
}