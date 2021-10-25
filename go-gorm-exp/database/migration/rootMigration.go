package migration

import (
    "strings"
    "fmt"
    // "reflect"
    "os"
    "gorm.io/gorm"
    "path/filepath"
    "github.com/iancoleman/strcase"

    "example.com/go-gorm-exp/models"
)

func GetMigrationFile() []string {
    var files []string

    migrationFolder := "./database/migration"
    err := filepath.Walk(migrationFolder, func(path string, info os.FileInfo, err error) error {
        if info.IsDir() {
            return nil
        }
        var fileNameFull = info.Name()
        var extension = filepath.Ext(fileNameFull)
        var fileName = fileNameFull[0:len(fileNameFull)-len(extension)]
        if (fileName != "rootMigration") {
            files = append(files, fileName)
        }
        return nil
    })
    if err != nil {
        panic(err)
    }
    return files
}

func ParseFuncName(name string) string {
    split := strings.Index(name, "_")
    return strcase.ToCamel(name[split+1:len(name)])
}

func RunMigrationFile(fileName string, migrationIndex uint, isUp bool) {
    var funcName = ParseFuncName(fileName)
    if (isUp) {
        funcName = funcName + "Up"
    } else {
        funcName = funcName + "Down"
    }
    fmt.Println(fileName)
    fmt.Println(funcName)
    fmt.Println(migrationIndex)
}

func RunMigration(db *gorm.DB) {
    var migration = models.Migration{}
    var files = GetMigrationFile()
    db.Order("batch desc").First(&migration).Select("batch")
    var migrationIndex = migration.Batch + 1;
    // if (err != nil) {
    //     fmt.Println(migration)
    //     fmt.Println(migration.Batch)
    // } else {
    //     fmt.Println(migration)
    // }
    for _, file := range files {
        RunMigrationFile(file, migrationIndex, true)
    }
}

func RollbackMigration(db *gorm.DB) {
    InitMigrationDown(db)
}
