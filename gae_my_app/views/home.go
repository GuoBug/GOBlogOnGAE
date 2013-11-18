package views

const HomePage = `
{{template "header" .}}
        <title>Bug has no bug!</title>
    </head>
	<body>

        <div class="navbar navbar-default navbar-fixed-top" role="navigation">                
            <div class="container">
            {{template "navbar" .}}
            </div>
        </div>	
        <div class="container">
        	{{range . }}
            <div class="page-header">
                <h1>{{.Topic.Title}}</h1>
                <h6 class="text-muted">文章发表于 {{.Topic.Created}}</h6>
                <p>
                	{{.Topic.Content}}
                </p>
            </div>
            {{end}}
        </div>

        <script type="text/javascript" src="http://cdn.staticfile.org/jquery/2.0.3/jquery.min.js"></script>
        <script type="text/javascript" src="/static/js/bootstrap.min.js"></script>
    </body>
</html>
`
const TopicTemplate = `{{define "Topic"}}Id:{{.Topic.Id}},Uid:{{.Topic.Uid}},Title:{{.Topic.Title}},Content:{{.Topic.Content}},Attachment:{{.Topic.Attachment}},Created:{{.Topic.Created}},Updated:{{.Topic.Updated}},Views:{{.Topic.Views}},Author:{{.Topic.Author}},ReplyTime:{{.Topic.ReplyTime}},ReplyCount:{{.Topic.ReplyCount}},ReplyLastUId:{{.Topic.ReplyLastUId}}{{end}}`
