package main

import (
	"fmt"
	"net/http"
)

type testHandler struct {
	http.Handler
}

func (t *testHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Test")
}

func barHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "barTest")
}

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {

		for i := 0; i < 10; i++ {

			fmt.Fprint(w, "hello")
		}
	})

	http.HandleFunc("/hel", func(w http.ResponseWriter, r *http.Request) {

		fmt.Fprint(w, "hel")
	})
	http.HandleFunc("/testbar", barHandler)
	http.Handle("/test", &testHandler{})

	http.ListenAndServe(":3000", nil)

}
