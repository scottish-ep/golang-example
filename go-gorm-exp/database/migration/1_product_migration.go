package migration

import (
    "gorm.io/gorm"

    "example.com/go-gorm-exp/models"
)

func ProductMigrationUp(db *gorm.DB) {
    db.AutoMigrate(&models.Product{})
}

func ProductMigrationDown(db *gorm.DB) {
    db.Migrator().DropTable(&models.Product{})
}

func ProductMigration() {
    MigrationStorage["ProductMigrationUp"] = ProductMigrationUp
    MigrationStorage["ProductMigrationDown"] = ProductMigrationDown
}
