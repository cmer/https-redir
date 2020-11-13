package main
import (
    "net/http"
    "os"
    "flag"
    "fmt"
)

var port string
var help bool

func redirect(w http.ResponseWriter, req *http.Request) {
    println("👍  Redirecting to: " + "https://" + req.Host + req.URL.String())
    http.Redirect(w, req,
        "https://" + req.Host + req.URL.String(),
        http.StatusMovedPermanently)
}

func main() {
    flag.StringVar(&port, "port", "80", "Port to listen on")
    flag.BoolVar(&help, "help", false, "Display Help")
    flag.Parse()

    if help {
        fmt.Fprintf(os.Stderr, "Usage of %s:\n", os.Args[0])
        flag.PrintDefaults()
        os.Exit(1)
    }

    println("👂 Listening on port " + port + "...")
    http.ListenAndServe(":" + port, http.HandlerFunc(redirect))
}