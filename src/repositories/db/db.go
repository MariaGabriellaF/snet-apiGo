package db

import (
	"snet-apiGo/src/repositories/config"
	"database/sql"
	"errors"
	"fmt"

	_ "github.com/lib/pq"
)

//para conectar ao banco de dados

func AbrirConexao() (*sql.DB, error) {
	conf, err := configs.Load()
	fmt.Println(conf.Host, conf.Port, conf.Pass, conf.User)
	if err != nil {
		return nil, err
	}

	if conf == (configs.DBConfig{}) {
		return nil, errors.New("Configurações DB não inicializadas ou vazias")
	}

	stringCon := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", conf.Host, conf.Port, conf.User, conf.Pass, conf.Database)

	conexao, err := sql.Open("postgres", stringCon)

	if err != nil {
		return nil, err
	}

	if pingErr := conexao.Ping(); pingErr != nil {
		conexao.Close()
		return nil, pingErr
	}

	return conexao, nil
}
