package file

import (
	"github.com/Cekretik/file-service/tree/main/models"
	"gorm.io/gorm"
)

type FileService struct {
	DB *gorm.DB
}

func NewFileService(db *gorm.DB) *FileService {
	return &FileService{DB: db}
}

func (f *FileService) AutoMigrate() {
	f.DB.AutoMigrate(&models.File{})
}
