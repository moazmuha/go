package main

import "fmt"

type contactInfo struct {
	email string
	number string
}

type person struct {
	firstName string
	lastName string
	contactInfo
}

func main() {

	moaz := 
		person{
			firstName:  "Moaz", 
			lastName: "Muhammad", 
			contactInfo: contactInfo{
						email: "moaz@gmail.com", 
						number: "123",
					},
		 }

	moaz.introduce()

	// fmt.Println(moaz)

	// var huzaifa person

	// huzaifa.firstName = "Huzaifa"
	// huzaifa.lastName = "Muhammad"

	// fmt.Printf("%+v", huzaifa)

	moaz.updateEmail("newemail")

	println("Moaz's new email is ", moaz.contactInfo.email)

	moaz.updateName("New Moaz") 

	println("Moaz's new name is ", moaz.firstName)
	
	l := []int{1,2}
	fmt.Println(l)
	swap(l)
	fmt.Println(l)

	s := "s"

	fmt.Println(&s)
	printPointer(&s)

}

func (p person) introduce() {
	fmt.Printf("Hi I am %s %s and my number is %s.\n", p.firstName, p.lastName, p.contactInfo.number)
}

func (p *person) updateEmail(email string) {
	p.contactInfo.email = email
}

func (p *person) updateName(name string){
	(*p).firstName = name
}

func swap(l []int){
	l[0], l[1] = l[1], l[0]
}

func printPointer(s *string){
	fmt.Println(s)
}