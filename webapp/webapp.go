package main

import (
	"fmt"
	"net/http"
)

func indexHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, `<h1>Whoa, Go!</h1>
	<p>Whoa, Go!</p>
	<p>Whoa, Go!</p>
	`)
	fmt.Fprintf(w, "<h1>Whoa, Go!</h1>")
	fmt.Fprintf(w, "<p>Whoa, Go!</p>")
	fmt.Fprintf(w, "<p>Whoa, Go!</p>")
	fmt.Fprintf(w, "<p>You %s even add %s</p>", "can", "<strong>variables</strong>")

}

func aboutHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Expert web design")
}

func main() {
	http.HandleFunc("/", indexHandler)
	http.HandleFunc("/about", aboutHandler)
	http.ListenAndServe(":8000", nil)
}
