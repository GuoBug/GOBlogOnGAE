package views

const TopicHTML = `
{{template "header"  .}}
        <title>文章——Bug has no bug!</title>
    </head>

    <body>
    	<div class="navbar navbar-default navbar-fixed-top" role="navigation">                
    	    <div class="container">
    	    {{template "navbar" .}}
    	    </div>
    	</div>
    	<div class="container">
            <h1>文章列表</h1>
            <a type="submit" class="btn btn-default" href="/topic/add">添加文章</a>
            <table class="table table-striped">
                <thead>
                    <tr>
                        <th>#</th>
                        <th>标题</th>
                        <th>更新时间</th>
                        <th>浏览数</th>
                        <th>回复数</th>
                        <th>操作</th>
                    </tr>
                </thead>
                <tbody>
                    {{range  .}}
                    <tr>
                        <th>{{.Topic.Id}}</th>
                        <th><a href="/topic/view/{{.Topic.Id}}">{{.Topic.Title}}</a></th>
                        <th>{{.Topic.Updated}}</th>
                        <th>{{.Topic.Views}}</th>
                        <th>{{.Topic.ReplyCount}}</th>
                        <th>
                            <a href="/topic/modify?tid={{.Topic.Id}}">修改</a>
                            <a href="/topic/del?{{.Topic.Id}}">删除</a>
                        </th>
                    </tr>
                    {{end}}
                </tbody>
            </table>
    	</div>
    </body>

</html>
`
