package views

var HeadTemplateHtml = `
{{define "header"}}
<!DOCTYPE html>

<html>
  	<head>
	    	<link rel="shortcut icon" href="/static/img/favicon.jpg" />	
    	<meta http-equiv="Content-Type" content="text/html; charset=utf-8">
             <!-- Stylesheets -->
              <link href="/static/css/bootstrap.min.css" rel="stylesheet" />
{{end}}
`