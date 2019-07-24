package repo

import (
	"github.com/jmoiron/sqlx"
	/*
		github.com/go-sql-driver/mysql não e usado diretamente usado pela apliação
	*/
	_ "github.com/go-sql-driver/mysql"
)

//Db armazena a conexão com banco de dados
var Db *sqlx.DB

//AbreConexaoComBancoDeDados função que abre a conexão com banco MYSQL
func AbreConexaoComBancoDeDados() (err error) {
	Db, err = sqlx.Open("mysql", "root:root@tcp(127.0.0.1:3306)/cursodego")
	if err != nil {
		return
	}
	err = Db.Ping()
	if err != nil {
		return
	}
	return
}
