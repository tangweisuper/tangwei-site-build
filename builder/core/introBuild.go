package core

import (
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"log"
	"os"
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

	data, err := ioutil.ReadAll(src)
	if err != nil {
		return err
	}

	defer src.Close()

	err = ioutil.WriteFile(dest, data, 0666)

	return err
}
