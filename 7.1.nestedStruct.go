package main

import(
	"fmt"
)
type name string

type Salary struct{
	basic int
	insurance int
	allowance int
}

type Employee struct{
	first name
	last name
	salary Salaried
	bool
}


type Salaried interface{
	getSalary() int 
}

func (s Salary) getSalary() int{
	return s.basic + s.insurance + s.allowance
}

func main(){
	bob:= Employee{
		first: "Bob",
		last: "dylan",
		bool: true,
		salary: Salary{
			basic: 30000,
			insurance: 7000,
			allowance: 60000,
		},
	}

	fmt.Println(bob)

	fmt.Printf("Employee %v's salary is:%v\n",bob.first, bob.salary.getSalary())
}