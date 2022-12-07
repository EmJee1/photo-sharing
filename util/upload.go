package util

import (
	"errors"
	"github.com/google/uuid"
	"gorm.io/gorm/utils"
	"io"
	"mime/multipart"
	"os"
	"path"
	"path/filepath"
)

const fileUploadDir = "uploads"

func isSupportedFileExtension(fileExt string) bool {
	supportedExts := []string{".jpg", ".jpeg", ".png", ".webp"}
	return utils.Contains(supportedExts, fileExt)
}

func UploadImage(file *multipart.FileHeader) (string, error) {
	src, err := file.Open()
	if err != nil {
		return "", err
	}
	defer src.Close()

	fileExt := filepath.Ext(file.Filename)
	if isSupported := isSupportedFileExtension(fileExt); isSupported != true {
		return "", errors.New("that image-type is not supported")
	}

	filename := uuid.New().String() + fileExt
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
