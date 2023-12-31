package data

import "fmt"

var Countries [8]string
var Slice []int
var Codes map[string]int


func init() {
    Countries[0] = "USA"
    Countries[1] = "China"
    Countries[2] = "India"
    Countries[3] = "Germany"
    Countries[4] = "France"
    Countries[5] = "Spain"
    Countries[6] = "Italy"
    Countries[7] = "UK"
    fmt.Println("Countries initialized")
}
