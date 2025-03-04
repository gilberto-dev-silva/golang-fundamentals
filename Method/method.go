package main

import "fmt"

type user struct {
	name string
	age  uint8
}

func (u user) save() {
	fmt.Printf("Saving the data of User %s in the database\n", u.name)
}

func (u user) overage() bool {
	return u.age >= 18
}

func (u *user) makeBirthday() {
	u.age++
}

func main() {
	user1 := user{"User 1", 20}
	user1.save()

	user2 := user{"David", 30}
	user2.save()

	overage := user2.overage()
	fmt.Println(overage)

	user2.makeBirthday()
	fmt.Println(user2.age)

}
