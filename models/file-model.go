package models

import (
	"time"

	"gorm.io/gorm"
)

type File struct {
	ID         uint      `gorm:"primaryKey" json:"id"`
	FileName   string    `gorm:"not null" json:"file_name"`
	FileSize   int64     `gorm:"not null" json:"file_size"`
	UploadTime time.Time `gorm:"autoCreateTime" json:"upload_time"`
}

func (File) TableName() string {
	return "files"
}

type FileService struct {
	DB *gorm.DB
}

func NewFileService(db *gorm.DB) *FileService {
	return &FileService{DB: db}
}

func (f *FileService) AutoMigrate() {
	f.DB.AutoMigrate(&File{})
}
