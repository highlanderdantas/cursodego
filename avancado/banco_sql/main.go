package main

import (
	"fmt"
	"net/http"

	"github.com/highlanderdantas/cursodego/avancado/banco_sql/manipulador"
	"github.com/highlanderdantas/cursodego/avancado/banco_sql/repo"
)

func init() {
	fmt.Println("Vamos começar a subir o servidor....")
}

func main() {

	err := repo.AbreConexaoComBancoDeDados()
	if err != nil {
		fmt.Println("Parando a carga do servidor. Erro ao abrir o banco de dados: ", err.Error())
		return
	}

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "Olá mundo")
	})
	http.HandleFunc("/funcao", manipulador.Funcao)
	http.HandleFunc("/ola", manipulador.Ola)
	http.HandleFunc("/local/", manipulador.Local)

	fmt.Println("Estou escutando....")
	http.ListenAndServe(":8181", nil)

}
