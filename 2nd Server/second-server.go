package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

type Page struct {
	Title string
	Body  []byte
}

func main() {

	pag1 := &Page{Title: "Example", Body: []byte(
		"Hello little person! This is the body of your page.")}

	pag1.save()

	http.HandleFunc("/portfolio/", viewpageHandler)
	fmt.Println("[+] Your server is running...")
	http.ListenAndServe(":8080", nil)
}

func viewpageHandler(w http.ResponseWriter, r *http.Request) {
	title := r.URL.Path
	page, _ := chargePage(title)
	fmt.Fprintf(w, "<h1>%s</h1><div>%s</div>",
		page.Title[len("/portfolio/"):], page.Body)

	fmt.Println("here it's your page!")
}

func (p *Page) save() error {
	name := p.Title + ".txt"
	return ioutil.WriteFile("./portfolio/"+name, p.Body, 0600)
}

func chargePage(title string) (*Page, error) {
	file_name := title + ".txt"
	fmt.Println("The client asked for:" + file_name)
	body, err := ioutil.ReadFile("." + file_name)

	if err != nil {
		return nil, err
	}
	return &Page{Title: title, Body: body}, nil
}
