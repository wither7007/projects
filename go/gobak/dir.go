package main

import (
    "fmt"
    "os"
)

func main() {
    filepath := "data.txt"
    data, err := os.ReadFile(filepath)
    if err != nil {
        fmt.Println("File reading error", err)
        return
    }
    fmt.Println(string(data))
}
