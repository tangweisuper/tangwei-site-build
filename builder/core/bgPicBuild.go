package core

import (
	"io"
	"os"
)

func WritePic(picPath string, dest string) error {
	src, err := os.Open(picPath)
	if err != nil {
		return err
	}

	file, err := os.OpenFile(dest, os.O_WRONLY|os.O_CREATE, os.ModePerm)
	if err != nil {
		return err
	}

	defer src.Close()
	defer file.Close()

	_, err = io.Copy(file, src)

	return err
}
