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
	man person
	ltk bool
}

func (s secretAgent) speak(){
	fmt.Println("I am", s.man.first ,s.man.last,"- the secretAgent speak")
}

func (p person) speak(){
	fmt.Println("I am", p.first ,p.last,"- the Person speak")
}

// anything which has method speak() now also become of type human
// A value can be more than one type
type human interface{
	speak()
}

func bar(h human) {
	switch h.(type) {
		case person:
			fmt.Println("I was passed into bar",h.(person).first)
			h.speak()
		case secretAgent:
			fmt.Println("I was passed into bar",h.(secretAgent).man.first)
			h.speak()
	}
}


func main(){

	sa1 := secretAgent{
		man: person{
			first: "James",
			last: "Bond",
		},
		ltk: true,
	}

	sa2 := secretAgent{
		man: person{
			first: "Jason",
			last: "Bourne",
		},
		ltk: true,
	}

	p1:= person{
		first: "John",
		last: "wick",
	} 

	fmt.Println(sa1)
	sa1.speak()
	sa2.speak()

	fmt.Println(p1)

	bar(sa1)
	bar(sa2)
	bar(p1)

}