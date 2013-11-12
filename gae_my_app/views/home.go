package views

import 
var headTemplateHtml string
headTemplateHtml = `
{{define "header"}}
<!DOCTYPE html>
<html>
    <head>
    	<meta charset="UTF-8" />
{{end}}
`
var homePage string
homePage = `
{{template "header" .}}
                <title>Bug</title>
        </head>
        <body>
                <h1>Test</h1>
                <span>Template works successfully.</span>
        </body> 
</html>
`