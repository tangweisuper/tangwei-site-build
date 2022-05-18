# 静态页面生成
不同二十四节气生成时需要修改的部分：

1. config.toml

以下配置地址对应的图片，改为节气背景图

    [params]
        backgroundPath = "images/bg.png"
    
    
2. data/en/intro.yaml

一下配置改成对应节气的文本

    title:
      - line: '立夏'
      - line: '天地始交，万物并秀'

    description:
      - line: '《遵生八笺》'
      - line: '孟夏之日，天地始交，万物并秀'


