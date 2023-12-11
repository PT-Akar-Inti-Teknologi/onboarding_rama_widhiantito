package files

import "fmt"

type FileService interface {
	Create(newOrder *File) (*File, error)
}

type fileService struct {
	fileRepository FileRepository
}

func NewFileService(fileRepo FileRepository) FileService {
	return &fileService{
		fileRepository: fileRepo,
	}
}

func (fs *fileService) Create(newFile *File) (*File, error) {
	if err := fs.fileRepository.CreateFile(newFile); err != nil {
		return nil, fmt.Errorf("failed to create file: %w", err)
	}

	return newFile, nil
}
