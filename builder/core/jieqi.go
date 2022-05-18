package core

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

type JieQi struct {
	Name             string
	PicPath          string
	Subtitle         string
	DescriptionTitle string
	Description      string
}

func Get24JieQi() []JieQi {
	list := make([]JieQi, 24)
	for i := 0; i < 24; i++ {
		name := jieqiList[i]
		list[i] = JieQi{
			Name: name,
		}
	}

	return list
}
