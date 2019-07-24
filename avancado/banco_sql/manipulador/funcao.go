package manipulador

import (
	"fmt"
	"net/http"
)

//Funcao e manipulador de requisição
func Funcao(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Aqui e um manipulador usando função num pacote")
}
