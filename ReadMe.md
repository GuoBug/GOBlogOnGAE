我想在[GAE](https://appengine.google.com/)上重现[golang](http://golang.org/)的blog


目前结构:
|gae_my_app
	|controller		组视图 调方法
	|myServer		一些公共函数,算是基础库
	|static			静态文件,css,picture
	|views			用go 存的视图,应为貌似用模板文件总报错

所以建立了这个项目，
现在还很菜，但是不会一直这么菜（我是要成为大牛的男人！）

2013年11月11日
今日竟然会一直研究这个这么久。。。。 废寝没忘食……上学都没这么认真……
第一步，显示出网页
##[gowalker](http://gowalker.org/)是个好网站

2013年11月12日
gae不支持模板文件。今天查了半天，改了半天，还是不能成功，今天是个失败的版本

2013年11月13日
将模板文件改为模板变量，成功。变量进不去…… 还是失败

2013年11月14日
导入变量成功

[测试站](http://goguobug.appspot.com/)