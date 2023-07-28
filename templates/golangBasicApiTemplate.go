package templates

func GetGolangApiTemplate() string {
	return golangApiTemplate
}

const golangApiTemplate = `
package main

import (
	"fmt"
	"log"
	"net/http"
)

func getUserData (w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "{\n  \"name\": \"John\",\n  \"age\": 30,\n  \"city\": \"New York\"\n}")
}

func main() {
	fmt.Println("API running on port :{{.Port}}")
	http.HandleFunc("/users", getUserData)
	log.Fatal(http.ListenAndServe(":{{.Port}}", nil))
}

`
