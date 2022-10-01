package main

import (
	"net/http"

	"github.com/gorilla/mux"
	"twitter.com/views"
)

// var homeView *views.View
// var twitView *views.View
var signView *views.View

//var contactTemplate *template.Template

//	func home(w http.ResponseWriter, r *http.Request) {
//		w.Header().Set("Content-Type", "text/html")
//		err := homeView.Template.ExecuteTemplate(w,
//			homeView.Layout, nil)
//		if err != nil {
//			panic(err)
//		}
//	}
func signup(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	err := signView.Template.ExecuteTemplate(w,
		signView.Layout, nil)
	if err != nil {
		panic(err)
	}
}

// func twit(w http.ResponseWriter, r *http.Request) {
// 	w.Header().Set("Content-Type", "text/html")
// 	err := twitView.Template.ExecuteTemplate(w,
// 		twitView.Layout, nil)
// 	if err != nil {
// 		panic(err)
// 	}
// }

// func contact(w http.ResponseWriter, r *http.Request) {
// 	w.Header().Set("Content-Type", "text/html")
// 	if err := contactTemplate.Execute(w, nil); err != nil {
// 		panic(err)
// 	}
// }

func main() {
	// var err error
	// homeTemplate, err = template.ParseFiles(
	// 	"views/home.gohtml",
	// 	"views/twitter.gohtml")
	// if err != nil {
	// 	panic(err)
	// }
	// contactTemplate, err = template.ParseFiles(
	// 	"views/contact.gohtml",
	// 	"views/layouts/footer.gohtml")
	// if err != nil {
	// 	panic(err)
	// }
	// homeView = views.NewView("bootstrap",
	// 	"views/home.gohtml")
	// twitView = views.NewView("bootstrap",
	// 	"views/navbar.gohtml")
	signView = views.NewView("bootstrap",
		"views/users/newth.gohtml")

	r := mux.NewRouter()
	// r.HandleFunc("/", home)
	// r.HandleFunc("/twit", twit)
	r.HandleFunc("/", signup)
	//r.HandleFunc("/contact", contact)

	// Assets
	assetHandler := http.FileServer(http.Dir("./assets/"))
	assetHandler = http.StripPrefix("/assets/", assetHandler)
	r.PathPrefix("/assets/").Handler(assetHandler)

	http.ListenAndServe(":3000", r)
}
