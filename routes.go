//simple example
package main

import (
	"fmt"
	"net/http"
)

func main(){
	http.HandleFunc("/hello", func(w http.ResponseWriter, r *http.Request){
		fmt.Fprintf(w, "Hello world!")
	})

	http.ListenAndServe(":8080", nil)
}