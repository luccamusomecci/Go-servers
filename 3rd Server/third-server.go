package main

import (
	"fmt"
	"html/template"
	"io/ioutil"
	"net/http"
)

type Page struct {
	Title string
	Body  []byte
}

//Main of the program
func main() {
	pag1 := &Page{Title: "Example", Body: []byte(
		"Hello little person! This is the body of your page.")}
	pag1.save()

	http.HandleFunc("/portfolio/", viewpageHandler)
	http.HandleFunc("/edit/", editHandler)
	//http.HandleFunc("/save", saveHandler)
	fmt.Println("[+] Your server is running...")
	http.ListenAndServe(":3000", nil)
}

//Save the page method
func (p *Page) save() error {
	name := p.Title + ".txt"
	return ioutil.WriteFile("./portfolio/"+name, p.Body, 0600)
}

// Charge the page method
func chargePage(title string) (*Page, error) {
	file_name := title + ".txt"
	fmt.Println("The client asked:" + file_name)
	body, err := ioutil.ReadFile("./portfolio/" + file_name)
	if err != nil {
		return nil, err
	}
	return &Page{Title: title, Body: body}, nil
}

//Charge the HTML Templates
func chargeTemplate(w http.ResponseWriter, templatex string, p *Page) {
	t, _ := template.ParseFiles(templatex + ".html")
	t.Execute(w, p)
}

//Petitions Handler
func viewpageHandler(w http.ResponseWriter, r *http.Request) {
	title := r.URL.Path[len("/portfolio/"):]
	p, _ := chargePage(title)
	chargeTemplate(w, "portfolio", p)
}

//Editer Handler
func editHandler(w http.ResponseWriter, r *http.Request) {
	title := r.URL.Path[len("/edit/"):]
	p, err := chargePage(title)
	if err != nil {
		p = &Page{Title: title}
	}
	chargeTemplate(w, "edit", p)
}
