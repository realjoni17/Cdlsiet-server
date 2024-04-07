package migrations

import (
	"github.com/realjoni17/Cdlsiet/models"
	"gorm.io/gorm"
)

func RunMigrations(db *gorm.DB) {
	db.AutoMigrate(models.Lecture{})
	db.AutoMigrate(models.Clas{})
	db.AutoMigrate(models.Notes{})
	db.AutoMigrate(models.Books{})
	db.AutoMigrate(models.Pyq{})
	db.AutoMigrate(models.Syllabus{})
	db.AutoMigrate(models.Post{})
}
