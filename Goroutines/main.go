package main

import (
	"fmt"
	"time"
)

func printMessage(text string) {
    for i := 0; i < 10; i++{
        fmt.Println(text)
        time.Sleep(800 * time.Microsecond)
    }
}

func main() {
    go printMessage("Go is great")
    printMessage("threading is great")
}
