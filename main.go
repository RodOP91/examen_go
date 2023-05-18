package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
)

type Fruta struct {
	ID       int
	Nombre   string
	Cantidad int
}

var frutas []Fruta

func main() {
	http.HandleFunc("/frutas", handleFrutas)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func handleFrutas(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		getFrutas(w, r)
	case http.MethodPost:
		anadirFruta(w, r)
	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
		fmt.Fprintf(w, "Método HTTP no válido")
	}
}

func getFrutas(w http.ResponseWriter, r *http.Request) {
	for _, fruta := range frutas {
		fmt.Fprintf(w, "ID: %d, Nombre: %s, Cantidad: %d\n", fruta.ID, fruta.Nombre, fruta.Cantidad)
	}
}

func anadirFruta(w http.ResponseWriter, r *http.Request) {
	id := len(frutas) + 1
	nombre := r.FormValue("nombre")
	cantidadstr := r.FormValue("cantidad")

	cantidad, err := strconv.Atoi(cantidadstr)
	if err != nil {
		fmt.Println(err)
		w.WriteHeader(http.StatusBadRequest)
		fmt.Println("Tipo de dato no aceptado para cantidad de frutas")
	}

	// Validar si el registro ya existe
	for _, fruit := range frutas {
		if fruit.Nombre == nombre {
			w.WriteHeader(http.StatusBadRequest)
			fmt.Fprintf(w, "La fruta '%s' ya existe en el arreglo", nombre)
			return
		}
	}

	frutas = append(frutas, Fruta{
		ID:       id,
		Nombre:   nombre,
		Cantidad: cantidad,
	})

	fmt.Fprintln(w, "Fruta agregada con éxito")
}
