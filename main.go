package main

import (
	"html/template"
	// "io"
	"log"
	"net/http"
)

type player struct {
	Name string
    Hand deck
}

type players []player

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseGlob("templates/*"))
}

func main() {
	http.HandleFunc("/", foo)
	http.Handle("/res/", http.StripPrefix("/res", http.FileServer(http.Dir("./assets"))))
	http.ListenAndServe(":8080", nil)
}

func foo(w http.ResponseWriter, req *http.Request) {
	// w.Header().Set("Content-Type", "text/html; charset=utf-8")
	// io.WriteString(w, hand())
	d := newDeck()
	d = d.shuffle()
	d, h1 := d.deal(5)
	_, h2 := d.deal(5)
	p := players{
		player{
			Name: "bob", 
			Hand: h1,
			},
		player{
			Name: "eric", 
			Hand: h2,
		},
	} 
	err := tpl.ExecuteTemplate(w, "something.gohtml", p)
	if err != nil {
		log.Fatalln("error executing template", err)
	}
}


// func hand() deck {

// 	// var hand string
// 	// for _, img := range h {
// 	// 	hand = hand + `<img src="/res/` + img + `.svg">`
// 	// }
// 	return h
// }
