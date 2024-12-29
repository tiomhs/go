package controller

import (
	"database/sql"
	"net/http"
	"path/filepath"
	"text/template"
)

func NewUpdateEmployeeController(db *sql.DB) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {

		if r.Method == "POST" {
			id := r.URL.Query().Get("id")
			r.ParseForm()

			name := r.FormValue("name")
			address := r.FormValue("address")
			npwp := r.FormValue("npwp")

			_, err := db.Exec("UPDATE employee SET name = ?, address = ?, npwp = ? WHERE id = ?", name, address, npwp, id)
			if err != nil {
				w.Write([]byte(err.Error()))
				w.WriteHeader(http.StatusInternalServerError)
				return
			}

			http.Redirect(w, r, "/employee", http.StatusMovedPermanently)

		} else if r.Method == "GET" {

			id := r.URL.Query().Get("id")

			row := db.QueryRow("SELECT name, npwp, address FROM employee WHERE id = ?", id)

			var employee Employee
			err := row.Scan(
				&employee.Name,
				&employee.NPWP,
				&employee.Address,
			)

			employee.Id = id
			if err != nil {
				w.Write([]byte(err.Error()))
				w.WriteHeader(http.StatusInternalServerError)
				return
				
			}

			fp := filepath.Join("views", "update.html")
			tmpl, err := template.ParseFiles(fp)

			if err != nil {
				w.Write([]byte(err.Error()))
				w.WriteHeader(http.StatusInternalServerError)
				return
			}

			data := make(map[string]any)
			data["employee"] = employee

			err = tmpl.Execute(w, data)

			if err != nil {
				w.Write([]byte(err.Error()))
				w.WriteHeader(http.StatusInternalServerError)
				return
			}
		}

	}
}
