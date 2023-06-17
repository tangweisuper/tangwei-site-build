package core

import (
	"fmt"
	"gopkg.in/yaml.v2"
	"io"
	"log"
	"os"
	"tangwei-site-build/builder/consts"
)

type Intro struct {
	Title       Title       `yaml:"title"`
	Description Description `yaml:"description"`
}

type Title struct {
	Line []string `yaml:"line"`
}

type Description struct {
	Line []string `yaml:"line"`
}

func BuildIntro(intro Intro, savepath string) error {
	yamlData, err := yaml.Marshal(intro)
	if err != nil {
		return err
	}

	f, err := os.OpenFile(savepath, os.O_WRONLY&os.O_CREATE, 0666)
	if err != nil {
		log.Println(err.Error())
	}

	_, err = f.Write(yamlData)
	if err != nil {
		return err
	}

	return nil
}

func WriteYaml(srcPath string, dest string) error {
	src, err := os.Open(srcPath)
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

func Copy24Text(destFolder string) {
	err := os.MkdirAll(destFolder, os.ModePerm)
	if err != nil {
		fmt.Println("创建文件夹时出错:", err)
		return
	}
	for k, v := range pingyinMap {
		src := fmt.Sprintf("%s/%s.yaml", consts.TextFoldPath, k)
		dest := fmt.Sprintf("%s/%s.yaml", destFolder, v)
		err := WriteYaml(src, dest)
		if err != nil {
			log.Println(err.Error())
		}
	}
}
