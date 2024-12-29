package controller

import (
	"database/sql"
	"net/http"
	"path/filepath"
	"text/template"
)

func NewCreateEmployeeController(db *sql.DB) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {

		if r.Method == "POST" {
			r.ParseForm()

			name := r.FormValue("name")
			address := r.FormValue("address")
			npwp := r.FormValue("npwp")

			_, err := db.Exec("INSERT INTO employee (name, address, npwp) VALUES (?, ?, ?)", name, address, npwp)
			if err != nil {
				w.Write([]byte(err.Error()))
				w.WriteHeader(http.StatusInternalServerError)
				return
			}

			http.Redirect(w, r, "/employee", http.StatusMovedPermanently)

		} else if r.Method == "GET" {
			fp := filepath.Join("views", "create.html")
			tmpl, err := template.ParseFiles(fp)

			if err != nil {
				w.Write([]byte(err.Error()))
				w.WriteHeader(http.StatusInternalServerError)
				return
			}

			err = tmpl.Execute(w, nil)

			if err != nil {
				w.Write([]byte(err.Error()))
				w.WriteHeader(http.StatusInternalServerError)
				return
			}
		}

	}
}
