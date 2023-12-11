package files

import "gorm.io/gorm"

type FileRepository interface {
	CreateFile(file *File) error
	//GetFileDetails()
}

type fileRepo struct {
	db *gorm.DB
}

func NewFileRepository(db *gorm.DB) FileRepository {
	return &fileRepo{db}
}

func (fr *fileRepo) CreateFile(newFile *File) error {
	return fr.db.Create(&newFile).Error
}
