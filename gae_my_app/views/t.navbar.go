package views

var NavbarTemplateHtml = `
{{define "navbar"}}
<a class="navbar-brand" href="/">锅巴的博客</a>
<div>
    <ul class="nav navbar-nav">
        <li ><a href="/">首页</a></li>
        <li ><a href="/topic">文章</a></li>
    </ul>
</div>

<div class="pull-right">
	<wl class="nav navbar-nav">
		<li><a href="/login">登录</a></li>
	</wl>
</div>
{{end}}
`
