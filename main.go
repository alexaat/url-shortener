package main

import (
	"flag"
	"fmt"
	"net/http"
	"os"
	"url-shortener/urlshort"
)

func main() {

	mux := defaultMux()

	//----------- MAP -------------------
	// pathsToUrl := map[string]string{
	// 	"/urlshort-godoc": "https://godoc.org/github.com/gophercises/urlshort",
	// 	"/yaml-godoc":     "https://godoc.org/gopkg.in/yaml.v2",
	// }
	// mapHandler := urlshort.MapHandler(pathsToUrl, mux)
	// http.ListenAndServe(":8080", mapHandler)
	//-------------------------------------

	//----------- YAML -------------------
	fileName := flag.String("path", "urls.yaml", "The path to input file")
	flag.Parse()
	b, _ := os.ReadFile(*fileName)
	yamlHandler, error := urlshort.YAMLHandler(b, mux)
	if error != nil {
		fmt.Println(error)
	}
	http.ListenAndServe(":8080", yamlHandler)
	//-------------------------------------

	//----------------------WITHOUT handler.go------------------
	// http.ListenAndServe(":8080", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	// 	if path, found := pathsToUrl[r.URL.String()]; found {
	// 		http.Redirect(w, r, path, http.StatusFound)
	// 	} else {
	// 		w.Write([]byte("invalid url"))
	// 	}
	// }))
	//-----------------------------------------------------------
}

func defaultMux() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("/", hello)
	return mux
}

func hello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Unvalid url")
}
