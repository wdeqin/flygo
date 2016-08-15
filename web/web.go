package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"html/template"
	"io/ioutil"
	"net/http"
	"os"
)

var indexTemplate *template.Template
var db *sql.DB

func init() {
	f, err := os.Open("templates/index.html")
	checkErr(err)
	b, _ := ioutil.ReadAll(f)
	s := string(b)
	indexTemplate = template.Must(template.New("index").Parse(s))

	db, err = sql.Open("mysql", "wdeqin:wdeqin@/devdb?charset=utf8")
	checkErr(err)
}

type Path struct {
	Id       uint
	Name     string
	Time     string
	Location string
}

func main() {
	http.HandleFunc("/query", func(w http.ResponseWriter, r *http.Request) {
		id := r.FormValue("id")
		if id == "" {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		var ps []Path

		rows, err := db.Query("SELECT id, name, time, location FROM path")
		checkErr(err)
		defer rows.Close()

		for rows.Next() {
			var id uint
			var name string
			var location string
			var time string
			rows.Scan(&id, &name, &time, &location)
			p := Path{id, name, time, location}
			ps = append(ps, p)
		}

		fmt.Println(ps)
		indexTemplate.Execute(w, ps)

	})
	http.ListenAndServe(":8080", nil)
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}
