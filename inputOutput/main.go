package main
import ("example.com/go/io/data"; "fmt")


func calculateTax(price float32) (float32, float32) {
    return price * 0.08, price * 0.09
}

func main() {
    // how to use functions and multiple return values
    stateTax, federalTax := calculateTax(100)

    fmt.Println(stateTax, federalTax)
    fmt.Println(data.MinSpeed)
	printData()
}
