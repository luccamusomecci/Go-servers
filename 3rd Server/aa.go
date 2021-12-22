package main

import (
	"fmt"
	"html/template"
	"io/ioutil"
	"net/http"
)

type Pagina struct {
	Titulo string
	Cuerpo []byte
}

func main() {
	//Creamos y guardamos una página para que el cliente la pida
	pag1 := &Pagina{Titulo: "Ejemplo", Cuerpo: []byte(
		"¡Hola personita! Este es el cuerpo de tu página.")}
	pag1.guardar()
	http.HandleFunc("/view/", manejadorMostrarPagina)
	http.HandleFunc("/edit/", manejadorEditar)
	//http.HandleFunc("/save/", manejadorGuardar)
	fmt.Println("El servidor se encuentra en ejecución")
	http.ListenAndServe(":8080", nil)
}

//Método para guardar página
func (p *Pagina) guardar() error {
	nombre := p.Titulo + ".txt"
	return ioutil.WriteFile("./view/"+nombre, p.Cuerpo, 0600)
}

//Método para cargar página
func cargarPagina(titulo string) (*Pagina, error) {
	nombre_archivo := titulo + ".txt"
	fmt.Println("El cliente ha pedido:" + nombre_archivo)
	cuerpo, err := ioutil.ReadFile("./view/" + nombre_archivo)
	if err != nil {
		return nil, err
	}
	return &Pagina{Titulo: titulo, Cuerpo: cuerpo}, nil
}

//Carga las plantillas HTML
func cargarPlantilla(w http.ResponseWriter, plantilla string, p *Pagina) {
	t, _ := template.ParseFiles(plantilla + ".html")
	t.Execute(w, p)
}

//Manejador de peticiones
func manejadorMostrarPagina(w http.ResponseWriter, r *http.Request) {
	titulo := r.URL.Path[len("/view/"):]
	p, _ := cargarPagina(titulo)
	cargarPlantilla(w, "view", p)
}

//Manejador para editar wikis
func manejadorEditar(w http.ResponseWriter, r *http.Request) {
	titulo := r.URL.Path[len("/edit/"):]
	p, err := cargarPagina(titulo)
	if err != nil {
		p = &Pagina{Titulo: titulo}
	}
	cargarPlantilla(w, "edit", p)
}
