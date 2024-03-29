package data

import "fmt"

type Person struct {
    Id int
    FirstName string
    LastName string
    Age int
}

func (p Person) Print() string{
    return fmt.Sprintf("%v %v (%d)", p.FirstName, p.LastName, p.Age)
}

func (p Person) String() string {
    return fmt.Sprintf("---- %v ----- %v", p.FirstName, p.LastName)
}
