package util

import (
	"github.com/google/uuid"
	"io"
	"mime/multipart"
	"os"
)

const fileUploadDir = "/uploads"

func UploadImage(file *multipart.FileHeader) error {
	src, err := file.Open()
	if err != nil {
		return err
	}
	defer src.Close()

	dst, err := os.Create(uuid.New().String())
	if err != nil {
		return err
	}
	defer dst.Close()

	if _, err = io.Copy(dst, src); err != nil {
		return err
	}

	return nil
}
