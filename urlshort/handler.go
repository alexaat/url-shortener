package urlshort

import (
	"net/http"
	"gopkg.in/yaml.v3"
)

// MapHandler will return an http.HandlerFunc (which also
// implements http.Handler) that will attempt to map any
// paths (keys in the map) to their corresponding URL (values
// that each key in the map points to, in string format).
// If the path is not provided in the map, then the fallback
// http.Handler will be called instead.
func MapHandler(pathsToUrls map[string]string, fallback http.Handler) http.HandlerFunc {

	//-----------------My Code-------------------
	// return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	// 	if path, found := pathsToUrls[r.URL.String()]; found {
	// 		http.Redirect(w, r, path, http.StatusFound)
	// 	} else {
	// 		fallback.ServeHTTP(w, r)
	// 	}
	// })

	return func(w http.ResponseWriter, r *http.Request){
		url:= r.URL.Path
		if path, ok := pathsToUrls[url]; ok {
			http.Redirect(w, r, path, http.StatusFound)
			return
		}
		fallback.ServeHTTP(w, r)
	}
}

// YAMLHandler will parse the provided YAML and then return
// an http.HandlerFunc (which also implements http.Handler)
// that will attempt to map any paths to their corresponding
// URL. If the path is not provided in the YAML, then the
// fallback http.Handler will be called instead.
//
// YAML is expected to be in the format:
//
//     - path: /some-path
//       url: https://www.some-url.com/demo
//
// The only errors that can be returned all related to having
// invalid YAML data.
//
// See MapHandler to create a similar http.HandlerFunc via
// a mapping of paths to urls.
func YAMLHandler(yml []byte, fallback http.Handler) (http.HandlerFunc, error) {
	pathsToUrls := make(map[string]string)

	var data []UrlData
	err := yaml.Unmarshal([]byte(yml), &data)
	if err != nil {
		return nil, err
	}
	for _, value := range data {
		pathsToUrls[value.Path] = value.Url
	}
	return MapHandler(pathsToUrls, fallback), nil
}

type UrlData struct {
	Path string `yaml:"path"`
	Url  string `yaml:"url"`
}
