package main

import "fmt"

// custom --------------
type Age int

func (a Age) isAdult() bool {
	return a >= 18
}

// regular -------------  Can see this as class 
type User struct { // ! type like c types
	name string
	age  Age
}

func (u User) printUser() { // ! add func to User struct
	// value receiver will copy obj
	fmt.Printf("%#v\n", u)
	// u.age <- have access to all user data here
}

func (u *User) setName(name ...string) { // *args
	/* this one receive pointer to user object   */
	if len(name) > 0 { // actually its trick
		u.name = name[0] // ! <- reassign user name to user obj
	} else {
		u.name = "default_name"
	}

	fmt.Printf("user updated : %#v\n", u)
}

func newUserConstructor(name string, age Age) *User { // <- pointer
	return &User{ // <- reference
		name: name,
		age:  age,
	}
}

// -----------------------

func main() {
	user1 := User{"user1", 11}
	fmt.Printf("%#v\n", user1) // #v as verbal full go syntax
	fmt.Println(user1.name)

	user2 := newUserConstructor("user2", 22)
	user2.printUser()

	user2.setName("super")
	fmt.Println(user2.age.isAdult()) // true
	fmt.Println(user1.age.isAdult()) // false

}
