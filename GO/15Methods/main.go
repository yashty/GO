package main

import "fmt"

func main() {
	fmt.Println("Structus in golang")

	//no inheritance in golang: No super or parent

	yash := User{"yash", "yasg@gmail.com", true, 43}
	fmt.Println(yash)
	fmt.Printf("yash details are: %+v\n", yash)
	fmt.Printf("Name is %v and email is %v.", yash.Name, yash.Email)
	yash.GetStatus()
	yash.NewMail()

}

type User struct {
	Name   string
	Email  string
	Status bool
	Age    int
}

func (u User) GetStatus() {
	fmt.Println("Is User active:", u.Status)
}

func (u User) NewMail() {
	u.Email = "test@go.dev"
	fmt.Println("Eamil of this user is:", u.Email)
}
