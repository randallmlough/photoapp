package main

import (
	"fmt"
	"net/http"
)

func handlerfunc(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "<h1>DANNNNNNG BRUHHHHHH</h1>")
}
func main() {
	http.HandleFunc("/", handlerfunc)
	http.ListenAndServe(":3000", nil)
}
