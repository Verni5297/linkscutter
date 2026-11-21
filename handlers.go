package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"html/template"
	"net/http"
)

type shortenRequest struct {
	URL string `json:"url"`
}

func shortenHandler(db *sql.DB) http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {
		var req shortenRequest
		json.NewDecoder(r.Body).Decode(&req)
		if req.URL == "" {
			http.Error(w, "empty url", http.StatusBadRequest)
			return
		}
		mod_url := generateUrl(6)

		_, err := db.Exec(`
			INSERT INTO links (original_url, mod_url) VALUES (?, ?) 

		`, req.URL, mod_url)

		if err != nil {
			fmt.Println(err)
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(map[string]string{
			"modurl": mod_url,
		})

	}

}

func findOriginalHandler(db *sql.DB) http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {
		mod_url := r.PathValue("mod_url")
		fmt.Println(mod_url)

		var orig_url string

		err := db.QueryRow(`SELECT original_url FROM links WHERE mod_url=?`, mod_url).Scan(&orig_url)
		if err != nil {
			fmt.Println(err)

			return
		}
		fmt.Println(orig_url)
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(map[string]string{
			"orig_url": orig_url,
		})
	}

}

func testhandler(db *sql.DB) http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(map[string]string{})

	}

}

type Link struct {
	OriginalUrl string
	ModUrl      string
	FullUrl     string
}

type PageData struct {
	Links []Link
}

func linksHandler(db *sql.DB) http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {

		rows, err := db.Query("SELECT original_url, mod_url FROM links")

		if err != nil {
			fmt.Println(err)

			return
		}

		defer rows.Close()

		var links []Link

		for rows.Next() {
			var l Link
			rows.Scan(&l.OriginalUrl, &l.ModUrl)
			l.FullUrl = "http://" + r.Host + "/" + l.ModUrl
			fmt.Println(l.FullUrl)
			links = append(links, l)
		}

		tmpl, err := template.ParseFiles("index.html")

		if err != nil {

			fmt.Println(err)

			return

		}

		err = tmpl.Execute(w, PageData{Links: links})

		if err != nil {

			fmt.Println(err)
			return

		}

		fmt.Println(links)
	}

}

func testhandler2(db *sql.DB) http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {

		url := r.PathValue("url")

		fmt.Println(url)

	}

}

func welcomeHandler(w http.ResponseWriter, r *http.Request) {

	http.ServeFile(w, r, "index.html")

}

func redirectHandler(db *sql.DB) http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {
		mod_url := r.PathValue("mod_url")

		var orig_url string

		err := db.QueryRow(`SELECT original_url FROM links WHERE mod_url=?`, mod_url).Scan(&orig_url)
		if err != nil {
			fmt.Println(err)

			return
		}

		http.Redirect(w, r, orig_url, http.StatusFound)
	}

}
