package uploader

import (
	"github.com/nekizz/telegram-bot/internal/enum"
)

// service
type service struct {
	uploader Uploader
}

// Uploader interface that define upload behavior
type Uploader interface {
	Upload(string, []byte, string) (string, error)
}

// Service interface define upload method
type Service interface {
	UploadFile(string, []byte, string) (string, error)
}

// New initialize uploader service
func New(bucketName, region, url string, uploaderType enum.UploadPlatform) Service {
	switch uploaderType {
	case enum.S3Uploader:
		return &service{}
	case enum.SelfHosted:
		//Initialize storage
		return nil
	default:
		return &service{}
	}
}

// UploadFile upload file to storage
func (s *service) UploadFile(fileName string, data []byte, contentType string) (string, error) {
	return s.uploader.Upload(fileName, data, contentType)
}
