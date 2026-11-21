package main

import "net/http"

func main() {

	db := initDB()

	defer db.Close()

	mux := http.NewServeMux()
	mux.HandleFunc("POST /short", shortenHandler(db))
	mux.HandleFunc("GET /find/{mod_url}", findOriginalHandler(db))
	mux.HandleFunc("GET /test", testhandler(db))
	mux.HandleFunc("GET /{mod_url}", redirectHandler(db))
	mux.HandleFunc("GET /test/{url}", testhandler2(db))
	mux.HandleFunc("GET /", linksHandler(db))
	//mux.HandleFunc("GET /links", linksHandler(db))
	http.ListenAndServe(":8080", mux)
}
