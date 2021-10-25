package main

import (
    "gorm.io/gorm"
    "gorm.io/driver/mysql"

    "example.com/go-gorm-exp/database/migration"
)

func main() {
    dsn := "root:@tcp(127.0.0.1:3306)/gorm_testing_db"
    db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
    if err != nil {
        panic("failed to connect database")
    }

    migration.RunMigration(db)
}
