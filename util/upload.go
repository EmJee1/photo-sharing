package util

import (
	"github.com/google/uuid"
	"io"
	"mime/multipart"
	"os"
	"path"
	"path/filepath"
)

const fileUploadDir = "uploads"

func UploadImage(file *multipart.FileHeader) (string, error) {
	src, err := file.Open()
	if err != nil {
		return "", err
	}
	defer src.Close()

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
