package util

import (
	"github.com/google/uuid"
	"io"
	"mime/multipart"
	"os"
	"path"
)

const fileUploadDir = "uploads"

func UploadImage(file *multipart.FileHeader) error {
	src, err := file.Open()
	if err != nil {
		return err
	}
	defer src.Close()

	filePath := path.Join(fileUploadDir, uuid.New().String()+file.Filename)
	dst, err := os.Create(filePath)
	if err != nil {
		return err
	}
	defer dst.Close()

	if _, err = io.Copy(dst, src); err != nil {
		return err
	}

	return nil
}
