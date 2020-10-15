package main

//A value can be more than one type

import(
	"fmt"

)

type name string

//<keyword> <identifier> <type>
type person struct{
	first name
	last name
}

type secretAgent struct{
	person
	ltk bool
}

func (s secretAgent) speak(){
	fmt.Println("I am", s.first ,s.last,"- the secretAgent speak")
}

func (p person) speak(){
	fmt.Println("I am", p.first ,p.last,"- the Person speak")
}

// anything which has method speak() now also become of type human
//A value can be more than one type
type human interface{
	speak()
}

func bar(h human) {
	switch h.(type) {
	case person:
		fmt.Println("I was passed into barrrr", h.(person).first)
	case secretAgent:
		fmt.Println("I was passed into barrr", h.(secretAgent).first)
	}
	fmt.Println("I was passed into bar", h)
}


func main(){

	sa1 := secretAgent{
		person: person{
			first: "James",
			last: "Bond",
		},
		ltk:true,
	}

	sa2 := secretAgent{
		person: person{
			first: "Jason",
			last: "Bourne",
		},
		ltk: true,
	}

	p1:= person{
		first: "Dr.",
		last: "who",
	} 

	fmt.Println(sa1)
	sa1.speak()
	sa2.speak()

	fmt.Println(p1)

	bar(sa1)
	bar(sa2)
	bar(p1)

}