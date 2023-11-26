package urlshort

import (
	"encoding/json"
	"fmt"
	"net/http"

	"gopkg.in/yaml.v2"
)

// MapHandler returns an http.HandlerFunc that will attempt to map any
// paths to their corresponding URL. If the path is not provided in the map,
// then the fallback http.Handler will be called instead.
func MapHandler(pathsToUrls map[string]string, fallback http.Handler) http.HandlerFunc {
	fmt.Println("h4")

	return func(w http.ResponseWriter, r *http.Request) {
		path := r.URL.Path
		if dest, ok := pathsToUrls[path]; ok {
			fmt.Println("h5")
			http.Redirect(w, r, dest, http.StatusFound)
			return
		}
		fmt.Println("h6")
		fallback.ServeHTTP(w, r)
		fmt.Println("h7")
	}
}

// YAMLHandler parses the provided YAML and then returns an http.HandlerFunc
// that will attempt to map any paths to their corresponding URL.
// If the path is not provided in the YAML, then the fallback http.Handler
// will be called instead.
func YAMLHandler(yml []byte, fallback http.Handler) (http.HandlerFunc, error) {
	fmt.Println("h8")
	pathURLs, err := parseYAML(yml)
	if err != nil {
		fmt.Println("h9")
		return nil, err
	}
	fmt.Println("h10")
	pathsToUrls := buildMap(pathURLs)
	return MapHandler(pathsToUrls, fallback), nil
}

func parseYAML(yml []byte) ([]pathURL, error) {
	fmt.Println("h11")
	var pathURLs []pathURL
	err := yaml.Unmarshal(yml, &pathURLs)
	if err != nil {
		fmt.Println("h12")
		return nil, err
	}
	fmt.Println("h13")
	return pathURLs, nil
}

// JSONHandler parses the provided JSON and then returns an http.HandlerFunc
// that will attempt to map any paths to their corresponding URL.
// If the path is not provided in the JSON, then the fallback http.Handler
// will be called instead.
func JSONHandler(jsonData []byte, fallback http.Handler) (http.HandlerFunc, error) {
	fmt.Println("h14")
	pathURLs, err := parseJSON(jsonData)
	fmt.Println(pathURLs)
	if err != nil {
		fmt.Println("h15")
		return nil, err
	}
	fmt.Println("h16")
	pathsToUrls := buildMap(pathURLs)
	return MapHandler(pathsToUrls, fallback), nil
}

func parseJSON(jsonData []byte) ([]pathURL, error) {
	fmt.Println("h17")
	var pathURLs []pathURL
	err := json.Unmarshal(jsonData, &pathURLs)
	if err != nil {
		fmt.Println("h18")
		return nil, err
	}
	fmt.Println("h19")
	return pathURLs, nil
}

func buildMap(pathURLs []pathURL) map[string]string {
	pathsToUrls := make(map[string]string)
	for _, pu := range pathURLs {
		pathsToUrls[pu.Path] = pu.URL
	}
	return pathsToUrls
}

type pathURL struct {
	Path string `yaml:"path"`
	URL  string `yaml:"url"`
}
