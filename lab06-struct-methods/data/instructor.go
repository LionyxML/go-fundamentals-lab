package data

import "fmt"

type Instructor struct {
	Id        int
	FirstName string
	LastName  string
	Score     int
}

func (i Instructor) Print() string {
	return fmt.Sprintf("%v, %v (%d)", i.LastName, i.FirstName, i.Score)
}

// factory pattern
func NewInstructor(id int, firstName string, lastName string, score int) Instructor {
	return Instructor{id, firstName, lastName, score}
}
