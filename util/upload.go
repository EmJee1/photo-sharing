package util

import (
	"errors"
	"github.com/google/uuid"
	"gorm.io/gorm/utils"
	"io"
	"mime/multipart"
	"net/http"
	"os"
	"path"
	"path/filepath"
)

const fileUploadDir = "uploads"

func getSupportedImageFormats() []string {
	return []string{"image/jpeg", "image/png", "image/webp"}
}

func isSupportedImageFormat(file multipart.File) bool {
	// Docs state that DetectContentType only takes the first 512 bytes into consideration
	buff := make([]byte, 512)
	if _, err := file.Read(buff); err != nil {
		return false
	}

	contentType := http.DetectContentType(buff)
	return utils.Contains(getSupportedImageFormats(), contentType)
}

func UploadImage(file *multipart.FileHeader) (string, error) {
	src, err := file.Open()
	if err != nil {
		return "", err
	}
	defer src.Close()

	if isSupported := isSupportedImageFormat(src); isSupported != true {
		return "", errors.New("that image-type is not supported")
	}

	filename := uuid.New().String() + filepath.Ext(file.Filename)
	filePath := path.Join(fileUploadDir, filename)
	dst, err := os.Create(filePath)
	if err != nil {
		return "", err
	}
	defer dst.Close()

	if _, err = io.Copy(dst, src); err != nil {
		return "", err
	}

	return filePath, err
}
