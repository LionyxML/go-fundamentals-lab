package main

type User struct {
	id   int
	name string
}

func (u User) PrettyPrint() string {
	return string(u.id) + ": " + u.name
}

func main() {
	var u1 User
	u1 = User{id: 1, name: "my cool name"}

	u2 := User{2, "cool name of two"}

	u1.PrettyPrint()
	u2.PrettyPrint()

}
