package migration

import (
    "gorm.io/gorm"
)

func RunMigration(db *gorm.DB) {
    InitMigrationUp(db)
}

func RollbackMigration(db *gorm.DB) {
    InitMigrationDown(db)
}
