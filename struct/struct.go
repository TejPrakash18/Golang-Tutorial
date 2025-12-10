package main

import "fmt"


type address struct{
	houseNumber int
	street string
	landmark string
}

// struct ambedding
type order struct{
	id int
	price int
	status string
	add address
}


type Person struct {
	name   string
	age    int
	job    string
	salary int
}

func (p *Person) changeName(name string){
	p.name = name
}

func (p Person) readName() string{
	return p.name
}

func printPerson(pers Person) {
	fmt.Println(pers.name)
	fmt.Println(pers.age)
	fmt.Println(pers.job)
	fmt.Println(pers.salary)
}

type intern struct{
	role string
	durection int
	isPaid bool
}

func newIntern(role string, durection int, isPaid bool) *intern{
	Myintern := intern{
		role: role,
		durection: durection,
		isPaid: isPaid,
	}
	return &Myintern

}

func main() {

	var per1 Person
	per1.name = "Tej"
	per1.age = 24
	per1.job = "Backend Engineer"
	per1.salary = 40000

	per2 := Person{
		name: "TP",
		age: 24,
		job: "SDE-1",
		salary: 45000,

	}

	printPerson(per1)
	fmt.Println(per2)

	per1.changeName("Tarun")
	printPerson(per1)
	var name = per1.readName()
	fmt.Println(name)

	//struct embedding
	myOrder := order{
		id: 1,
		price: 49,
		status: "received",
		add: address{
			houseNumber: 120,
			street: "Sasni Road",
			landmark: "Iglas",
		},
	}

	fmt.Println(myOrder)
 

	newIntern := newIntern("Backend",3,true)
	fmt.Println(newIntern)
}