package file

import (
	"file-service/models"
	"fmt"
	"io"
	"mime/multipart"
	"os"

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

func (f *FileService) UploadFile(fileHeader *multipart.FileHeader) (*models.File, error) {
	// Создание начальной записи о файле в базе данных
	fileRecord := &models.File{
		FileName: fileHeader.Filename,
		FileSize: fileHeader.Size,
	}
	if err := f.DB.Create(fileRecord).Error; err != nil {
		return nil, err
	}

	// Открыть файл для чтения
	file, err := fileHeader.Open()
	if err != nil {
		return nil, err
	}
	defer file.Close()

	// Путь для сохранения файла, включая ID
	filePath := fmt.Sprintf("path/to/storage/%d-%s", fileRecord.ID, fileHeader.Filename)

	// Сохранение файла
	dst, err := os.Create(filePath)
	if err != nil {
		return nil, err
	}
	defer dst.Close()

	// Копирование содержимого файла
	if _, err = io.Copy(dst, file); err != nil {
		return nil, err
	}

	// Обновление записи в базе данных с полным путем к файлу
	fileRecord.Path = filePath
	if err := f.DB.Save(fileRecord).Error; err != nil {
		return nil, err
	}

	return fileRecord, nil
}
