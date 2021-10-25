package models

import (
    "gorm.io/gorm"
)

type Migration struct {
    gorm.Model
    Batch uint
    Migration string
}
