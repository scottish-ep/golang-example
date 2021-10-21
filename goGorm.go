package main

import (
    "gorm.io/driver/mysql"
    "gorm.io/gorm"
)

type Product struct {
    gorm.Model
    Code  string
    Price uint
}

func main() {
    dsn := "root:@tcp(127.0.0.1:3306)/gorm_testing_db"
    db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
    if err != nil {
        panic("failed to connect database")
    }

    db.AutoMigrate(&Product{})
    
    // Create
    db.Create(&Product{Code: "D42", Price: 100})
}
