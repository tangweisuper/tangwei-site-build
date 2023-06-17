package consts

const (
	PicFoldPath  = "resource/pics"
	TextFoldPath = "resource/text"

	HugoBGPicFolder = "web/static/images"
	HugoBGPicPath   = HugoBGPicFolder + "/bg.webp"
	HugoIntroPath   = "web/data/en/intro.yaml"

	HugoBuildPath = "web/public"

	BuildDistPath = "dist"

	FlaskTemplateIntroPath     = TextFoldPath + "/flask_template.yaml"
	FlaskTemplatePicUrl        = "{{bg_url}}"
	FlaskTemplateBuildDistPath = BuildDistPath + "/flask_template"
)
