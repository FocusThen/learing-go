package main

import "focusthen.com/server/data"

func main() {
	p := data.Person{Id: 1}
	p.FirstName = "Focus"

	print(p.Print())
}


