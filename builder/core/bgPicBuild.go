package core

import (
	"io/ioutil"
	"os"
)

func WritePic(picPath string, dest string) error {
	src, err := os.Open(picPath)
	if err != nil {
		return err
	}

	data, err := ioutil.ReadAll(src)
	if err != nil {
		return err
	}

	defer src.Close()

	err = ioutil.WriteFile(dest, data, 0666)

	return err
}
