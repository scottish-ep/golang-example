package main

import "fmt"

func Concat4(args ...interface{}) string {
    a := "default-a"
    b := 5

    for _, arg := range args {
        fmt.Println(arg)
        switch t := arg.(type) {
        case string:
            a = t
        case int:
            b = t
        default:
            panic("Unknown argument")
        }
    }

    return fmt.Sprintf("%s%d", a, b)
}

func main() {
    fmt.Println(Concat4(1, 2))
    fmt.Println(Concat4(1, 2, 3, 4))
    fmt.Println(Concat4(1))
}
