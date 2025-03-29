package database

import (
	"lab6/models"
	"log"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

// DB es la instancia global de conexión a la base de datos
var DB *gorm.DB

// Connect inicializa la conexión a la base de datos y realiza la migración automática del modelo Match
func Connect() {
	var err error

	// Abrir la conexión con una base de datos SQLite llamada matches.db
	DB, err = gorm.Open(sqlite.Open("matches.db"), &gorm.Config{})
	if err != nil {
		// Si hay un error al conectar, detener la aplicación
		log.Fatal("No se pudo conectar a la base de datos:", err)
	}

	// Ejecutar la migración automática del modelo Match (crear tabla si no existe)
	err = DB.AutoMigrate(&models.Match{})
	if err != nil {
		log.Fatal("No se pudo migrar el modelo:", err)
	}

	// Mensaje informativo si todo funcionó correctamente
	log.Println("📦 Conectado a la base de datos y migrado correctamente")
}
