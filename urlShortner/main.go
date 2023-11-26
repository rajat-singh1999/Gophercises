package main

import (
	"flag"
	"fmt"
	"io"
	"net/http"
	"os"
	"urlShortner/urlshort"
)

var yamlFile string
var jsonFile string

func main() {
	// Defining flags for the YAML and JSON files
	flag.StringVar(&yamlFile, "yaml", "", "Path to the YAML file (without extension)")
	flag.StringVar(&jsonFile, "json", "", "Path to the JSON file (without extension)")
	flag.Parse()

	// Checking if both flags are provided, using the yaml in case both are provided
	if yamlFile != "" && jsonFile != "" {
		fmt.Println("Both YAML and JSON flags provided. Using YAML.")
		jsonFile = ""
	}

	// exiting if no flag is provided
	if yamlFile == "" && jsonFile == "" {
		return
	}

	mux := defaultMux()

	// Building the MapHandler using the mux as the fallback, using two starter entries
	pathsToUrls := map[string]string{
		"/urlshort-godoc": "https://godoc.org/github.com/gophercises/urlshort",
		"/yaml-godoc":     "https://godoc.org/gopkg.in/yaml.v2",
	}
	mapHandler := urlshort.MapHandler(pathsToUrls, mux)

	var content []byte
	var err error

	// Reading content from the specified file based on the provided flag
	if yamlFile != "" {
		content, err = readFileContent(yamlFile + ".yaml")
		if err != nil {
			panic(err)
		}
		handler, err := urlshort.YAMLHandler(content, mapHandler)
		if err != nil {
			panic(err)
		}
		fmt.Println("Starting the server on :8080")
		http.ListenAndServe(":8080", handler)
	} else if jsonFile != "" {
		content, err = readFileContent(jsonFile + ".json")
		if err != nil {
			panic(err)
		}
		jsonHandler, err := urlshort.JSONHandler(content, mapHandler)
		if err != nil {
			panic(err)
		}
		fmt.Println("Starting the server on :8080")
		http.ListenAndServe(":8080", jsonHandler)
	} else { // this condition will not be reached
		fmt.Println("No flag was provided")
	}
}

func readFileContent(filename string) ([]byte, error) {
	fileHandle, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer fileHandle.Close()

	content, err := io.ReadAll(fileHandle)
	if err != nil {
		return nil, err
	}

	return content, nil
}

func defaultMux() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("/", hello)
	return mux
}

func hello(w http.ResponseWriter, r *http.Request) {
	var content []byte
	var err error

	// Reading content from the specified file based on the provided flag
	if yamlFile != "" {
		content, err = readFileContent(yamlFile + ".yaml")
		if err != nil {
			panic(err)
		}
	} else if jsonFile != "" {
		content, err = readFileContent(jsonFile + ".json")
		if err != nil {
			panic(err)
		}
	}

	fmt.Fprintf(w, "%s", content)
}
