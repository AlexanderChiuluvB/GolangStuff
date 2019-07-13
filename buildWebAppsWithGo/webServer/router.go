package main

import (
	"fmt"
	"net/http"
)

type MyMux struct {
}

func (p *MyMux) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path == "/" {
		sayhelloName(w, r)
		return
	}
	//open http://localhost:9090/alex
	//it will run openALex function
	if r.URL.Path =="/alex"{
		openAlex(w,r)
		return
	}
	http.NotFound(w, r)
	return
}

func openAlex(w http.ResponseWriter,r *http.Request){

	fmt.Fprintf(w,"hello alex!")

}

func sayhelloName(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello myroute!")
}

func main() {
	mux := &MyMux{}
	http.ListenAndServe(":9090", mux)
}