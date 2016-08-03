package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"html/template"
	"io/ioutil"
	"net/http"
	"os"
	"time"
)

var indexTemplate *template.Template
var db *sql.DB

func init() {
	f, err := os.Open("templates/index.html")
	if err != nil {
		panic(err)
	}
	b, _ := ioutil.ReadAll(f)
	s := string(b)
	indexTemplate = template.Must(template.New("index").Parse(s))

	if db, err = sql.Open("mysql", "wdeqin:wdeqin@tcp(localhost:3306)/devdb"); err != nil {
		panic(err)
	}
}

type Path struct {
	Id       int
	Name     string
	Time     time.Time
	Location string
}

func main() {
	http.HandleFunc("/hello", func(w http.ResponseWriter, r *http.Request) {
		rows, err := db.Query("SELECT id, name, time, location FROM devdb.path")
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			panic(err)
		}
		defer rows.Close()

		p := Path{}
		if rows.Next() {
			rows.Scan(&p.Id, &p.Name, &p.Time, &p.Location)
			fmt.Printf("%#v\n", p)
		}
		indexTemplate.Execute(w, p)

	})
	http.ListenAndServe(":8080", nil)
}
