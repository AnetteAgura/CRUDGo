package main

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
	"html/template"
	"log"
	"net/http"
	"os"
)

type Games struct {
	Id      int
	Name    string
	Year    int
	Company string
}

func init() {
	if err := godotenv.Load(); err != nil {
		log.Print("No .env file found")
	}
}

func dbConn() (db *sql.DB) {
	dbDriver := os.Getenv("DBDRIVER")
	dbUser := os.Getenv("DBUSER")
	dbPass := os.Getenv("DBPASS")
	dbName := os.Getenv("DBNAME")
	db, err := sql.Open(dbDriver, dbUser+":"+dbPass+"@/"+dbName)
	if err != nil {
		panic(err.Error())
	}
	return db
}

var tmpl = template.Must(template.ParseGlob("templates/*"))

func Index(w http.ResponseWriter, r *http.Request) {
	db := dbConn()
	selDB, err := db.Query("SELECT * FROM games ORDER BY id ASC")
	if err != nil {
		panic(err.Error())
	}
	game := Games{}
	var res []Games
	for selDB.Next() {
		var id, year int
		var name, company string
		err = selDB.Scan(&id, &name, &year, &company)
		if err != nil {
			panic(err.Error())
		}
		game.Id = id
		game.Name = name
		game.Year = year
		game.Company = company
		res = append(res, game)
	}
	tmpl.ExecuteTemplate(w, "index.html", res)
	defer db.Close()
}

func Show(w http.ResponseWriter, r *http.Request) {
	db := dbConn()
	nId := r.URL.Query().Get("id")
	selDB, err := db.Query("SELECT * FROM games WHERE id=?", nId)
	if err != nil {
		panic(err.Error())
	}
	game := Games{}
	for selDB.Next() {
		var id, year int
		var name, company string
		err = selDB.Scan(&id, &name, &year, &company)
		if err != nil {
			panic(err.Error())
		}
		game.Id = id
		game.Name = name
		game.Year = year
		game.Company = company
	}
	tmpl.ExecuteTemplate(w, "show.html", game)
	defer db.Close()
}

func Company(w http.ResponseWriter, r *http.Request) {
	db := dbConn()
	nId := r.URL.Query().Get("company")
	selDB, err := db.Query("SELECT * FROM games WHERE company=?", nId)
	if err != nil {
		panic(err.Error())
	}
	game := Games{}
	var res []Games
	for selDB.Next() {
		var id, year int
		var name, company string
		err = selDB.Scan(&id, &name, &year, &company)
		if err != nil {
			panic(err.Error())
		}
		game.Id = id
		game.Name = name
		game.Year = year
		game.Company = company
		res = append(res, game)
	}
	tmpl.ExecuteTemplate(w, "company.html", res)
	defer db.Close()
}

func Year(w http.ResponseWriter, r *http.Request) {
	db := dbConn()
	nId := r.URL.Query().Get("year")
	selDB, err := db.Query("SELECT * FROM games WHERE year=?", nId)
	if err != nil {
		panic(err.Error())
	}
	game := Games{}
	var res []Games
	for selDB.Next() {
		var id, year int
		var name, company string
		err = selDB.Scan(&id, &name, &year, &company)
		if err != nil {
			panic(err.Error())
		}
		game.Id = id
		game.Name = name
		game.Year = year
		game.Company = company
		res = append(res, game)
	}
	tmpl.ExecuteTemplate(w, "company.html", res)
	defer db.Close()
}

func New(w http.ResponseWriter, r *http.Request) {
	tmpl.ExecuteTemplate(w, "new.html", nil)
}

func Edit(w http.ResponseWriter, r *http.Request) {
	db := dbConn()
	nId := r.URL.Query().Get("id")
	selDB, err := db.Query("SELECT * FROM games WHERE id=?", nId)
	if err != nil {
		panic(err.Error())
	}
	game := Games{}
	for selDB.Next() {
		var id, year int
		var name, company string
		err = selDB.Scan(&id, &name, &year, &company)
		if err != nil {
			panic(err.Error())
		}
		game.Id = id
		game.Name = name
		game.Year = year
		game.Company = company
	}
	tmpl.ExecuteTemplate(w, "edit.html", game)
	defer db.Close()
}

func Insert(w http.ResponseWriter, r *http.Request) {
	db := dbConn()
	if r.Method == "POST" {
		name := r.FormValue("name")
		year := r.FormValue("year")
		company := r.FormValue("company")
		insForm, err := db.Prepare("INSERT INTO games(name, year, company) VALUES(?,?,?)")
		if err != nil {
			panic(err.Error())
		}
		insForm.Exec(name, year, company)
		log.Println("INSERT: Name: " + name + " | Year: " + year + " | Company: " + company)
	}
	defer db.Close()
	http.Redirect(w, r, "/", 301)
}

func Update(w http.ResponseWriter, r *http.Request) {
	db := dbConn()
	if r.Method == "POST" {
		name := r.FormValue("name")
		year := r.FormValue("year")
		company := r.FormValue("company")
		id := r.FormValue("uid")
		insForm, err := db.Prepare("UPDATE games SET name=?, year=?, company=? WHERE id=?")
		if err != nil {
			panic(err.Error())
		}
		insForm.Exec(name, year, company, id)
		log.Println("UPDATE: Name: " + name + " | Year: " + year + " | Company: " + company)
	}
	defer db.Close()
	http.Redirect(w, r, "/", 301)
}

func Delete(w http.ResponseWriter, r *http.Request) {
	db := dbConn()
	game := r.URL.Query().Get("id")
	delForm, err := db.Prepare("DELETE FROM games WHERE id=?")
	if err != nil {
		panic(err.Error())
	}
	delForm.Exec(game)
	log.Println("DELETE")
	defer db.Close()
	http.Redirect(w, r, "/", 301)

}

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	log.Println("Server started on: http://localhost:8080")
	http.HandleFunc("/", Index)
	http.HandleFunc("/show", Show)
	http.HandleFunc("/company", Company)
	http.HandleFunc("/year", Year)
	http.HandleFunc("/new", New)
	http.HandleFunc("/edit", Edit)
	http.HandleFunc("/insert", Insert)
	http.HandleFunc("/update", Update)
	http.HandleFunc("/delete", Delete)
	http.ListenAndServe(":8080", nil)

}
