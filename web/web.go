package web

import (
	db "cf_help_bot/database"
	user "cf_help_bot/user"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"
)

type Page struct {
	Title string
	Body  []byte
}

func loadPage(title string) (*Page, error) {
	filename := title + ".txt"
	body, err := os.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	return &Page{Title: title, Body: body}, nil
}

func Start() {
	http.HandleFunc("/", homeHandler)
	http.HandleFunc("/login", loginHandler)
	http.HandleFunc("/register", registerHandler)
	fmt.Println("Server started on http://localhost:8000")
	http.ListenAndServe(":8000", nil)
}

func homeHandler(w http.ResponseWriter, r *http.Request) {
	title := r.URL.Path[len("/"):]
	p, err := loadPage(title)
	if err != nil {
		p = &Page{Title: title}
	}
	t, _ := template.ParseFiles("web/template/home.html")
	t.Execute(w, p)
}

func loginHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		handle := r.FormValue("handle")
		err, isPersonExists := db.Does_person_exist_in_database_by_handle(handle)
		log.Println(isPersonExists)
		if err != nil {
			log.Fatal(err)
		}
		if isPersonExists {
			data, err := db.Get_user_data_by_handle(handle)
			if err != nil {
				log.Fatal(err)
			}
			log.Println(data)
			http.Redirect(w, r, "/", http.StatusSeeOther)
		} else {
			log.Println("Person doesn't exist in database")
			http.Redirect(w, r, "/register", http.StatusSeeOther)
		}
	} else {
		title := r.URL.Path[len("/login"):]
		p, err := loadPage(title)
		if err != nil {
			p = &Page{Title: title}
		}
		t, _ := template.ParseFiles("web/template/login.html")
		t.Execute(w, p)
	}
}

func registerHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		handle := r.FormValue("handle")
		lang := r.FormValue("lang")
		isLangSelected := true
		isLangSelection := false
		err, isPersonExists := db.Does_person_exist_in_database_by_handle(handle)
		if err != nil {
			log.Fatal(err)
		}
		if isPersonExists {
			log.Println("Person exists")
			http.Redirect(w, r, "/login", http.StatusSeeOther)
		} else {
			// userId := new id TODO db
			u := user.User{}
			userId := 0
			u.Initialize(userId, handle, lang, isLangSelected, isLangSelection)
			db.Set_user_data(userId, u, isLangSelected, isLangSelection, lang)
			http.Redirect(w, r, "/", http.StatusSeeOther)
		}
	} else {
		title := r.URL.Path[len("/register"):]
		p, err := loadPage(title)
		if err != nil {
			p = &Page{Title: title}
		}
		t, _ := template.ParseFiles("web/template/register.html")
		t.Execute(w, p)
	}
}
