package views

const TopicAddHTML = `
{{template "header"}}
        <title>添加文章——Bug has no bug!</title>
    </head>

    <body>
      <div class="navbar navbar-default navbar-fixed-top">                
          <div class="container">
          {{template "navbar" .}}
          </div>
      </div>
      <div class="container">
           <h1>添加文章</h1> 
           <form method="post" action="/topic">
               <div class="form-group">
                   <label>文章标题</label>
                   <input type="text" class="form-control" name="title">
               </div>
               <div>
                 <label>分类</label>
                 <input type="text" class="form-control" name="category">
               </div>
               <div>
                   <label>文章内容</label>
                   <textarea name="content" cols="30" rows="10" class="form-control"></textarea>
               </div>
               <button type="submit" class="btn btn-default">添加文章</button>
           </form>
      </div>
        <script type="text/javascript" src="http://cdn.staticfile.org/jquery/2.0.3/jquery.min.js"></script>
        <script type="text/javascript" src="/static/js/bootstrap.min.js"></script>
    </body>

</html>
`
