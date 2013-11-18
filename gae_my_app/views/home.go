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
        	{{template "Topic" .}}
            <div class="page-header">
                <h1>{{.Title}}</h1>
                <h6 class="text-muted">文章发表于 {{.Created}}</h6>
                <p>
                	{{.Content}}
                </p>
            </div>
            {{end}}
        </div>

        <script type="text/javascript" src="http://cdn.staticfile.org/jquery/2.0.3/jquery.min.js"></script>
        <script type="text/javascript" src="/static/js/bootstrap.min.js"></script>
    </body>
</html>
`
const TopicTemplate = `{{define "Topic"}}Title:{{.topic.Title}},Created:{{.topic.Created}},Content:{{.topic.Created}}{{end}}`
