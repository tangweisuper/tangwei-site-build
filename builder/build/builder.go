package build

import (
	"tangwei-site-build/builder/consts"
	"tangwei-site-build/builder/core"
)

func Build() error {
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
