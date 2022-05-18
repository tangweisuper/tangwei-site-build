# tangwei-site-build
用来生成 tangwei.cool 主站点

[点此访问](https://tangwei.cool)

##  说明
网站使用hugo生成静态页面

此builder的作用是可以根据24个节气，每次build出24个拥有不同背景图和文本的页面。

需要在网站部署后根据当前所属节气路由到对应地址，可以达到动态的根据节气实时切换的效果

## 使用
前提条件：

- 安装hugo

- themes文件夹中安装hugo-theme-massively模板

 
     https://github.com/curtiscde/hugo-theme-massively


运行

    /builder/main.go


## 过程
1. 将对应的背景图写入

2. 将对应文本写入

3. 执行命令


        hugo

4. 在dist文件夹中创建对应节气文件夹
5. 将web/public文件夹中的所有内容拷贝到节气文件夹中