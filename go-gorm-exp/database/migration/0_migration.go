package migration

import (
    "gorm.io/gorm"

    "example.com/go-gorm-exp/models"
)

func InitMigrationUp(db *gorm.DB) {
    db.AutoMigrate(&models.Migration{})
}

func InitMigrationDown(db *gorm.DB) {
    db.Migrator().DropTable(&models.Migration{})
}
