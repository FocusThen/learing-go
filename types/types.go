package main

import "fmt"


type location string

func (origin location) destinationTo(destination location) string{
    fmt.Printf("Calculating route from %s to %s\n", origin, destination)
    return "route"
}


func locationTest(){
    ist := location("Istanbul")

    ist.destinationTo("Ankara")

    fmt.Println(ist)
}
