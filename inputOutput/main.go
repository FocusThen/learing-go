package main
import ("example.com/go/io/data"; "fmt")


func calculateTax(price float32) (float32, float32) {
    return price * 0.08, price * 0.09
}

func calculateTaxWithName (price float32) (stateTax float32, fedaralTax float32) {
    stateTax = price * 0.08
    fedaralTax = price * 0.09
    return
}

func main() {
    // how to use functions and multiple return values
    _, federalTax := calculateTaxWithName(100)

    fmt.Println(federalTax)
    fmt.Println(data.MinSpeed)
	printData()
}
