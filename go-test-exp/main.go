package main

import "fmt"

func Sum(x int, y int) int {
    return x + y
}

func Sub(x int, y int) int {
    return x - y
}

func main() {
    sum := Sum(5, 5)
    fmt.Println(sum)
}
