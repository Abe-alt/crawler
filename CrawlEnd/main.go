package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

func main() {
	m := http.NewServeMux()
	m.HandleFunc("/", handlerPage)
	m.Handle("/main.js", http.FileServer(http.Dir(".")))

	const addr = ":8080"

	srv := http.Server{
		Handler:      m,
		Addr:         addr,
		WriteTimeout: 30 * time.Second,
		ReadTimeout:  30 * time.Second,
	}

	fmt.Println("Server started on port:", addr)
	err := srv.ListenAndServe()
	log.Fatal(err)

}

func handlerPage(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	w.Header().Set("Access-Control-Allow-Headers", "GET")
	const page = `
<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <title>Document</title>
    <script type="module" src="main.js"></script>
</head>
<body>
<input type="text" id="crawl">
<button id="actionBtn">Crawl this page</button>
<div id="result"></div>
</body>
</html>`
	w.WriteHeader(200)
	w.Write([]byte(page))
}
