package main

import (
    "fmt"
    "net/http"
    "log"
    "os"
	"bytes"
	"text/template"
)

const (
    DEFAULT_PORT = "9000"
)

var indexPage string =  `<!DOCTYPE html>
        <html lang="en">
        <head>
        <meta charset="UTF-8">
        <title>Go Web App</title>
        </head>
        <body>
        <center>
	<h1> {{.Name}}  </h1>

	<form>

	<table cellspacing="10px" cellpadding="10px">
	<tr>
	    <td>User Name:</td>
	    <td><input type="text" name="userName" /></td>
	</tr>
	<tr>
	    <td>Password:</td>
            <td><input type="password" name="password" /></td>
	</tr>
	<tr>
	    <td></td>
    	    <td><input type="button" name="login" value="Login" /></td>
	</tr>	
	</table>
	
	</form>	
	</center>
        </body>
        </html>`

type Sample struct {
	Name string
}


func HelloServer(w http.ResponseWriter, req *http.Request) {
    var t *template.Template
    t = template.Must(template.New("sample").Parse(indexPage))
    var page bytes.Buffer
    sample := Sample{
	Name: "Go Web App",
    }
    t.Execute(&page,sample)
    fmt.Fprintln(w, page.String())
 }

func main() {
    var port string
    if port = os.Getenv("PORT"); len(port) == 0 {
        log.Printf("Warning, PORT not set. Defaulting to %+vn", DEFAULT_PORT)
        port = DEFAULT_PORT
    }
  
    http.HandleFunc("/", HelloServer)
    err := http.ListenAndServe(":" + port, nil)
    if err != nil {
        log.Printf("ListenAndServe: ", err)
    }
}

