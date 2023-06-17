package build

import (
	"fmt"
	"tangwei-site-build/builder/consts"
	"tangwei-site-build/builder/core"
)

func BuildFlaskTemplate() error {
	//清除dist文件夹
	output, err := core.Command(fmt.Sprintf("rm -rf %s/*", consts.BuildDistPath))
	println(output)
	if err != nil {
		return err
	}

	intro := consts.FlaskTemplateIntroPath

	err = core.WriteYaml(intro, consts.HugoIntroPath)
	if err != nil {
		println(err.Error())
		return err
	}

	err = core.RunHugoBuild()
	if err != nil {
		println(err.Error())
		return err
	}

	err = core.CopyHugoBuildFiles(consts.FlaskTemplateBuildDistPath)
	if err != nil {
		println(err.Error())
		return err
	}

	//编译完之后处理

	core.Copy24Pic(consts.FlaskTemplateBuildDistPath + "/images")

	core.Copy24Text(consts.FlaskTemplateBuildDistPath + "/texts")

	return nil
}

func Build24() error {
	//清除dist文件夹
	output, err := core.Command(fmt.Sprintf("rm -rf %s/*", consts.BuildDistPath))
	println(output)
	if err != nil {
		return err
	}

	jieqiList := core.Get24JieQi()
	for _, v := range jieqiList {
		intro := v.IntroTextPath
		pic := v.PicPath

		err := core.WriteYaml(intro, consts.HugoIntroPath)
		if err != nil {
			println(err.Error())
			return err
		}

		err = core.WritePic(pic, consts.HugoBGPicPath)
		if err != nil {
			println(err.Error())
			return err
		}

		err = core.RunHugoBuild()
		if err != nil {
			println(err.Error())
			return err
		}

		err = core.CopyHugoBuildFiles(v.BuildDistPath)
		if err != nil {
			println(err.Error())
			return err
		}
	}

	return nil
}
