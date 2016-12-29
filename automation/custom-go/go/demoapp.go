package main
  
import (
    "fmt"
    "net/http"
    "log"
    "os"
)
  
const (
    DEFAULT_PORT = "9000"
)
  
func HelloServer(w http.ResponseWriter, req *http.Request) {
    fmt.Fprintln(w, "Lets Start working on Automation")
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
