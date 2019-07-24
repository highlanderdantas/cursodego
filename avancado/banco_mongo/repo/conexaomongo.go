package repo

import (
	"gopkg.in/mgo.v2"
)

//SessaoMongo armazena a sessão de conexão com o mongo
var SessaoMongo *mgo.Session

//AbreSessaoComMongo faz a conexão com o mongo
func AbreSessaoComMongo() (err error) {

	err = nil

	SessaoMongo, err = mgo.Dial("mongodb://localhost:27017/cursodego")
	return
}
