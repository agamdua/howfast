package main

import (
	"io"
	"log"
	"net/http"
	"os"

	"github.com/julienschmidt/httprouter"
)

func main() {
	router := httprouter.New()
	router.GET("/downloads", Downloads)

	port, set := os.LookupEnv("HOWFAST_PORT")

	if !set {
		port = "8080"
	}

	log.Println("Starting server")
	log.Fatal(http.ListenAndServe(":"+port, router))
}

func readFile() (*os.File, error) {
	filePath, set := os.LookupEnv("HOWFAST_DL_FILEPATH")

	if !set {
		// TODO: We actually need to panic here :)
		filePath = "/dev/null"
	}

	f, err := os.Open(filePath)

	if err != nil {
		log.Println("The file could not be opened")
		return f, err
	}

	return f, nil
}

// Downloads is a request handler for the download test
func Downloads(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	log.Println("Testing download speeds")
	log.Println(r.RemoteAddr)

	f, err := readFile()

	defer f.Close()

	if err != nil {
		// TODO: don't panic
		panic(err)
	}

	log.Println("Sending file over")

	io.Copy(w, f)
}
