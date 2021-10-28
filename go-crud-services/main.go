package main

import (
    "fmt"
)

type CRUDRepository struct {}

func (r CRUDRepository) create() {
    fmt.Println("create")
}

type ProductRepository struct {
    CRUDRepository
}

func (r ProductRepository) checkQuantity() {
    fmt.Println("checkQuantity")
}

func main() {
    var r ProductRepository = ProductRepository{}
    r.create()
    r.checkQuantity()
}
