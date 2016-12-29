package main
  
import (
    "fmt"
    "net/http"
    "log"
    "os"
	 "html/template"
	"strings"
)
  
const (
    DEFAULT_PORT = "9000"
)
  
func sayhelloName(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintln(w, "Start working with automation")
r.ParseForm()  //Parse url parameters passed, then parse the response packet for the POST body (request body)
    // attention: If you do not call ParseForm method, the following data can not be obtained form
    fmt.Println(r.Form) // print information on server side.
    fmt.Println("path", r.URL.Path)
    fmt.Println("scheme", r.URL.Scheme)
    fmt.Println(r.Form["url_long"])
    for k, v := range r.Form {
        fmt.Println("key:", k)
        fmt.Println("val:", strings.Join(v, ""))
    }
    fmt.Fprintf(w, "Hello astaxie!") // write data to response
}
  
func login(w http.ResponseWriter, r *http.Request) {
    fmt.Println("method:", r.Method) //get request method
    if r.Method == "GET" {
        t, _ := template.ParseFiles("login.gtpl")
        t.Execute(w, nil)
    } else {
        r.ParseForm()
        // logic part of log in
        fmt.Println("username:", r.Form["username"])
        fmt.Println("password:", r.Form["password"])
    }
}
func main() {
    var port string
    if port = os.Getenv("PORT"); len(port) == 0 {
        log.Printf("Warning, PORT not set. Defaulting to %+vn", DEFAULT_PORT)
        port = DEFAULT_PORT
    }
  
    http.HandleFunc("/", sayhelloName)
    http.HandleFunc("/login", login)
    err := http.ListenAndServe(":" + port, nil)
    if err != nil {
        log.Printf("ListenAndServe: ", err)
    }
}
