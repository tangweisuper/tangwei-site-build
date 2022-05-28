package core

import (
	"fmt"
	"os"
	"tangwei-site-build/builder/consts"
)

func RunHugoBuild() error {
	res, err := Command("cd web && hugo")
	println(res)
	return err
}

func CopyHugoBuildFiles(dest string) error {
	err := mkdir(dest)
	if err != nil {
		return err
	}

	cmd := fmt.Sprintf("cp -r %s/* %s", consts.HugoBuildPath, dest)

	res, err := Command(cmd)
	println(res)
	return err
}

func mkdir(path string) error {
	_, err := os.Stat(path)
	if err != nil {
		notExist := os.IsNotExist(err)
		if notExist {
			err := os.MkdirAll(path, os.ModePerm)
			return err
		}
	}
	return err
}
