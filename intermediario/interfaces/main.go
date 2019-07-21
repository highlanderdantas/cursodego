package main

import (
	"fmt"

	"cursodego/intermediario/interfaces/model"
)

func queroAcordarComUmCacarejo(g model.Galinha) {
	fmt.Println(g.Cacareja())
}

func queroOuvirUmaPataNoLago(p model.Pata) {
	fmt.Println(p.Grasna())
}

func main() {
	jojo := model.Ave{}
	jojo.Nome = "Jojo da Silva"

	queroAcordarComUmCacarejo(jojo)
	queroOuvirUmaPataNoLago(jojo)
}
