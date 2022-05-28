package core

import (
	"path"
	"tangwei-site-build/builder/consts"
)

var jieqiList = []string{
	"立春",
	"雨水",
	"惊蛰",
	"春分",
	"清明",
	"谷雨",
	"立夏",
	"小满",
	"芒种",
	"夏至",
	"小暑",
	"大暑",
	"立秋",
	"处暑",
	"白露",
	"秋分",
	"寒露",
	"霜降",
	"立冬",
	"小雪",
	"大雪",
	"冬至",
	"小寒",
	"大寒"}

var pingyinMap = make(map[string]string)

func init() {
	pingyinMap["立春"] = "lichun"
	pingyinMap["雨水"] = "yushui"
	pingyinMap["惊蛰"] = "jingzhe"
	pingyinMap["春分"] = "chunfen"
	pingyinMap["清明"] = "qingming"
	pingyinMap["谷雨"] = "guyu"
	pingyinMap["立夏"] = "lixia"
	pingyinMap["小满"] = "xiaoman"
	pingyinMap["芒种"] = "mangzhong"
	pingyinMap["夏至"] = "xiazhi"
	pingyinMap["小暑"] = "xiaoshu"
	pingyinMap["大暑"] = "dashu"
	pingyinMap["立秋"] = "liqiu"
	pingyinMap["处暑"] = "chushu"
	pingyinMap["白露"] = "bailu"
	pingyinMap["秋分"] = "qiufen"
	pingyinMap["寒露"] = "hanlu"
	pingyinMap["霜降"] = "shuangjiang"
	pingyinMap["立冬"] = "lidong"
	pingyinMap["小雪"] = "xiaoxue"
	pingyinMap["大雪"] = "daxue"
	pingyinMap["冬至"] = "dongzhi"
	pingyinMap["小寒"] = "xiaohan"
	pingyinMap["大寒"] = "dahan"
}

type JieQi struct {
	Name          string
	PicPath       string
	IntroTextPath string
	BuildDistPath string
}

func Get24JieQi() []JieQi {
	list := make([]JieQi, 24)
	for i := 0; i < 24; i++ {
		name := jieqiList[i]
		list[i] = JieQi{
			Name:          name,
			PicPath:       getPicPath(name),
			IntroTextPath: getIntroTextPath(name),
			BuildDistPath: getBuildDistPath(name),
		}
	}

	return list
}

func getPicPath(jieqiName string) string {
	return path.Join(consts.PicFoldPath, jieqiName+".png")
}

func getIntroTextPath(jieqiName string) string {
	return path.Join(consts.TextFoldPath, jieqiName+".txt")
}

func getBuildDistPath(jieqiName string) string {
	return path.Join(consts.BuildDistPath, pingyinMap[jieqiName])
}
