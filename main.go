package main

import (
	"log"
	"net/http"
	"os"

	"github.com/circle-makotom/hello-uname/handlers"
)

var (
	BuildVersion   = "dev"
	BuildTimeStamp = ""

	printer = log.New(os.Stdout, "", 0)
)

func main() {
	printer.Println(BuildVersion)
	printer.Println(BuildTimeStamp)
	printer.Println()

	http.Handle("/", new(handlers.HelloUnameHandler))
	log.Fatal(http.ListenAndServe(":8080", nil))
}
