package manipulador

import "html/template"

//ModeloOla armazena o modelo de pagina ola que serão executados pelos manipuladores
var ModeloOla = template.Must(template.ParseFiles("html/ola.html"))

//ModeloLocal armazena o modelo de pagina local que serão executados pelos manipuladores
var ModeloLocal = template.Must(template.ParseFiles("html/local.html"))
