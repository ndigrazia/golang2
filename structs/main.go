package main
import "fmt"

type contactInfo struct {
	email   string
	zipCode int
}

type person struct {
	firstName string
	lastName  string
	contactInfo //Add a nested struct
}

//Another way to declare a embedded struct
/*type person struct {
	firstName string
	lastName  string
	contact contactInfo  //Use the embedded struct: contact: contactInfo{...}
}*/

/*
//Basic struct
type person struct {
	firstName string
	lastName  string
}*/

func main() {
	/*//Three ways to define a struct 
	//One
	nelson := person{"Nelson", "DiGrazia"} //No recommended
	//Two
	juan := person{firstName: "Juan", lastName: "Santoro"}
	
	fmt.Println(juan)
	fmt.Println(nelson)

	//Three
	var alex person
	alex.firstName = "Alex"
	alex.lastName = "Anderson"
	fmt.Printf("%+v", alex)*/


	jim := person{
		firstName: "Jim",
		lastName:  "Party",
		contactInfo: contactInfo{
			email:   "jim@gmail.com",
			zipCode: 94000,
		},
	}

	jim.updateName("jimmy")
	jim.print()

	//fmt.Printf("%+v", jim)
}

func (p *person) updateName(newFirstName string) {
	(*p).firstName = newFirstName
}

func (p person) print() {
	fmt.Printf("%+v", p)
}
