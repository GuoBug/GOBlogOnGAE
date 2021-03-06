package views

const TopicViewHTML = `
{{template "header"}}
        <title>{{.Topic.Title}}——Bug has no bug!</title>
    </head>

    <body>
    	<div class="navbar navbar-default navbar-fixed-top">                
    	    <div class="container">
    	    {{template "navbar" .}}
    	    </div>
    	</div>
    	<div class="container">
           <h1>{{.Topic.Title}}</h1> 
           <h2><a href="/topic/modify?tid={{.Topic.Id}}" class="btn btn-default">修改</a> </h2>
           {{.Topic.Content}}
    	</div>
        <script type="text/javascript" src="http://cdn.staticfile.org/jquery/2.0.3/jquery.min.js"></script>
        <script type="text/javascript" src="/static/js/bootstrap.min.js"></script>
    </body>

</html>
`
