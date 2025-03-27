package database

import (
    "log"
    "gorm.io/driver/sqlite"
    "gorm.io/gorm"
    "lab6/models"
)

var DB *gorm.DB

func Connect() {
    var err error
    DB, err = gorm.Open(sqlite.Open("matches.db"), &gorm.Config{})
    if err != nil {
        log.Fatal("No se pudo conectar a la base de datos:", err)
    }

    err = DB.AutoMigrate(&models.Match{})
    if err != nil {
        log.Fatal("No se pudo migrar el modelo:", err)
    }

    log.Println("📦 Conectado a la base de datos y migrado correctamente")
}