package migration

import (
    "gorm.io/gorm"

    "example.com/go-gorm-exp/models"
)

func InitMigrationUp(db *gorm.DB) error {
    err := db.AutoMigrate(&models.Migration{})
    return err
}

func InitMigrationDown(db *gorm.DB) {
    db.Migrator().DropTable(&models.Migration{})
}
