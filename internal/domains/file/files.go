package files

import "gorm.io/gorm"

type File struct {
	gorm.Model
	Name        string `json:"name"`
	MimeType    string `json:"mime_type"`
	FilePath    string `json:"file_path"`
	IsUploaded  bool   `json:"is_uploaded"`
	IsResized   bool   `json:"is_resized"`
	ResizedSize int    `json:"resized_size"`
}
