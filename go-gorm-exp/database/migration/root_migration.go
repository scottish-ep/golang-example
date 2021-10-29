package migration

import (
    "fmt"
    "strings"
    "reflect"
    "errors"
    "os"
    "gorm.io/gorm"
    "path/filepath"
    "github.com/iancoleman/strcase"

    "example.com/go-gorm-exp/models"
)

type migrationMapping map[string]interface{}

var MigrationStorage = migrationMapping{}


func Call(funcName string, params ... interface{}) (result interface{}, err error) {
	f := reflect.ValueOf(MigrationStorage[funcName])
	if len(params) != f.Type().NumIn() {
		err = errors.New("The number of params is out of index.")
		return
	}
	in := make([]reflect.Value, len(params))
	for k, param := range params {
		in[k] = reflect.ValueOf(param)
	}
	result = f.Call(in)
	return
}

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
        if (fileName != "root_migration" && fileName != "register_migration") {
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

func RunMigrationFile(db *gorm.DB, fileName string, migrationIndex uint, isUp bool) {
    var funcName = ParseFuncName(fileName)
    var migration = models.Migration{}
    if (isUp) {
        funcName = funcName + "Up"
    } else {
        funcName = funcName + "Down"
    }
    Call(funcName, db)
    if (isUp) {
        result := db.Where(&models.Migration{Migration: fileName}).First(&migration)
        if (result.RowsAffected == 0) {
            fmt.Println(migration)
            migration = models.Migration{
                Batch: migrationIndex,
                Migration: fileName,
            }
            db.Create(&migration)
        }
    }
}

func RunMigration(db *gorm.DB, isUp bool) {
    var migration = models.Migration{}
    var files = GetMigrationFile()
    db.Order("batch desc").First(&migration).Select("batch")
    var migrationIndex = migration.Batch + 1;
    for _, file := range files {
        RunMigrationFile(db, file, migrationIndex, isUp)
    }
}
