package trash

import (
	"html/template"
	"math/rand"
	"net/http"
	"time"
)

var (
	repoo PostRepository = NewPostRepository()
)

func Home(w http.ResponseWriter, r *http.Request) {
	tmpl, _ := template.ParseFiles("templates/home.html")
	Posts, _ := repoo.FindAll()
	tmpl.Execute(w, Posts)
}

func Person(w http.ResponseWriter, r *http.Request) {
	tmpl, _ := template.ParseFiles("templates/person.html")
	Posts, _ := repoo.FindAll()
	lent := len(Posts)
	// Seed the random number generator
	rand.Seed(time.Now().UnixNano())

	// Generate a random number between 1 and x
	randomNumber := rand.Intn(lent)
	post := Posts[randomNumber]

	tmpl.Execute(w, post)
}
