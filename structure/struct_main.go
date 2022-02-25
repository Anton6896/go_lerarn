package main

import "fmt"

type User struct { // type like c types
	name string
	age  int
}

func (u User) printUser() { // add func to User struct 
	fmt.Printf("%#v\n", u)

	// u.age <- have access to all user data here 
}

func newUserConstructor(name string, age int) *User { // <- pointer
	return &User{ // <- reference
		name: name,
		age:  age,
	}
}

func main() {
	user1 := User{"user1", 11}
	fmt.Printf("%#v\n", user1) // #v as verbal full go syntax
	fmt.Println(user1.name)

	user2 := newUserConstructor("user2", 22)
	user2.printUser()

}
