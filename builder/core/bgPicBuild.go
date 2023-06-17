package core

import (
	"fmt"
	"io"
	"log"
	"os"
	"tangwei-site-build/builder/consts"
)

func WritePic(picPath string, dest string) error {
	src, err := os.Open(picPath)
	if err != nil {
		return err
	}

	data, err := io.ReadAll(src)
	if err != nil {
		return err
	}

	defer src.Close()

	err = os.WriteFile(dest, data, 0666)

	return err
}

func Copy24Pic(destFolder string) {
	for k, v := range pingyinMap {
		src := fmt.Sprintf("%s/%s.webp", consts.PicFoldPath, k)
		dest := fmt.Sprintf("%s/bg_%s.webp", destFolder, v)
		err := WritePic(src, dest)
		if err != nil {
			log.Println(err.Error())
		}
	}
}
