package main

import (
	"RahulStructs/data"
	"fmt"
)

func main() {

	rahul := data.Instructor{}
	rahul.FirstName = "Rahul"
	rahul.LastName = "Juliato"

	pru := data.Instructor{FirstName: "Pru", LastName: "Sanghi"}

	fmt.Println(rahul.Print())
	fmt.Println(pru.Print())

	navarro := data.NewInstructor(20, "Rodrigo", "Navaro", 100)
	fmt.Println(navarro.Print())

	goCourse := data.Course{Id: 2, Name: "Go fundamentals", Instructor: rahul}

	fmt.Printf("%v", goCourse)

	swiftWS := data.NewWorkshop("Swift on iOW", rahul)
	fmt.Printf("%v", swiftWS)

	var courses [2]data.Signable
	courses[0] = goCourse
	courses[1] = swiftWS

	for _, course := range courses {
		fmt.Println(course)
	}

}
