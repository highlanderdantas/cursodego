package main

import (
	"fmt"
	"net/http"

	"github.com/highlanderdantas/cursodego/avancado/servidor_web/manipulador"
)

func main() {

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "Olá mundo")
	})
	http.HandleFunc("/funcao", manipulador.Funcao)
	http.HandleFunc("/ola", manipulador.Ola)

	fmt.Println("Estou escutando....")
	http.ListenAndServe(":8181", nil)

}
