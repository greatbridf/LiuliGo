package main

import (
	"net/http"

	"github.com/greatbridf/LiuliGo/liuli/handlers"
)

func main() {
	//http.HandleFunc("/", handlers.HandleRoot)
	http.HandleFunc("/articles", handlers.HandleArticles)
	http.HandleFunc("/content", handlers.HandleContent)
	//http.HandleFunc("/magnet", handlers.HandleMagnet)
	//http.HandleFunc("/delete", handlers.HandleDelete)
	err := http.ListenAndServe(":14250", nil)
	if err != nil {
		panic(err)
	}
}
