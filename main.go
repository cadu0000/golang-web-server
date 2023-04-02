package main

import (
	"fmt"
	"log"
	"net/http"
)

func formHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/form" { //confere a url
		http.Error(w, "404 not found", http.StatusNotFound)
		return
	}
	nome := r.FormValue("nome")
	idade := r.FormValue("idade")
	if nome == "" || idade == "" {
		fmt.Fprintf(w, "invalid values")
		return
	}
	fmt.Fprintf(w, "nome = %s\n", nome)
	fmt.Fprintf(w, "idade = %s\n", idade)
}

func calculatorHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/calculator" {
		http.Error(w, "404 not found", http.StatusNotFound)
		return
	}

}
func main() {
	fileServer := http.FileServer(http.Dir("./pages"))

	http.Handle("/", fileServer)
	http.HandleFunc("/form", formHandler) //coanecta o html com a função
	http.HandleFunc("/calculator", calculatorHandler)

	fmt.Printf("Starting server at port 3030\n")
	err := http.ListenAndServe(":3030", nil)
	if err != nil {
		log.Fatal(err)
	}
}
