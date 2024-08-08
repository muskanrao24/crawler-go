package main

import (
	"distributed-web-crawler/front-end/controller"
	"net/http"
)

func main() {
	http.Handle("/", http.FileServer(
		http.Dir("front-end/view/")))
	http.Handle("/search",
		controller.CreateSearchResultHandler("front-end/view/template.html"))
	err := http.ListenAndServe(":8888", nil)
	if err != nil {
		panic(err)
	}
}
