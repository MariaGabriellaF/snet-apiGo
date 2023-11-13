package services

import (
	
	"log"
	"net/http"
	"snet-apiGo/src/models"
	"snet-apiGo/src/repositories"
	"snet-apiGo/src/repositories/db"
)

func CriarLoja(w http.ResponseWriter, loja models.Loja) (*models.Loja, error) {

	if len(loja.Nome) > 50 {
		return nil, ErrNomeExcedeuLimite
	}
	if len(loja.RazaoSocial) > 100 {
		return nil, ErrNomeExcedeuLimite
	}
	if len(loja.Endereco) > 100 {
		return nil, ErrNomeExcedeuLimite
	}
	if len(loja.Estado) > 20 {
		return nil, ErrNomeExcedeuLimite
	}
	if len(loja.Cidade) > 50 {
		return nil, ErrNomeExcedeuLimite
	}

	db, err := db.AbrirConexao()
	if err != nil {
		log.Println("Erro ao tentar abrir a conexão com o banco de dados")
		return &models.Loja{}, err
	}
	defer db.Close()

	repositorio := repositories.NewRepositories(db)
	loja.Id, err = repositorio.NewLoja(loja)
	if err != nil {
		log.Println("Erro ao criar a loja no banco de dados:", err)
		return &models.Loja{}, err
	}

	log.Println("Criação da loja concluída com sucesso")

	return &loja, nil
}

func GetLojas(w http.ResponseWriter, r *http.Request) []models.Loja {
	log.Println("Iniciando a obtenção da lista de lojas")

	db, err := db.AbrirConexao()
	if err != nil {
		log.Println("Erro ao tentar abrir a conexão com o banco de dados")
		return nil
	}
	defer db.Close()

	repositorio := repositories.NewRepositories(db)
	lojas, err := repositorio.ListarLojas()
	if err != nil {
		log.Println("Erro ao obter a lista de lojas:", err)
		return []models.Loja{}
	}

	log.Println("Obtenção da lista de lojas concluída com sucesso")

	return lojas
}

func DeletarLojaPorID(id int) error {
	log.Printf("Iniciando a exclusão da loja com ID %d\n", id)

	db, err := db.AbrirConexao()
	if err != nil {
		log.Println("Erro ao tentar abrir a conexão com o banco de dados")
		return err
	}
	defer db.Close()

	repositorio := repositories.NewRepositories(db)

	err = repositorio.DeletarLojaPorID(id)
	if err != nil {
		log.Printf("Erro ao excluir a loja com ID %d: %s\n", id, err)
		return err
	}

	log.Printf("Exclusão da loja com ID %d concluída com sucesso\n", id)

	return nil
}

func AtualizarLojaPorID(loja models.Loja) error {
	log.Printf("Iniciando a atualização da loja com ID %d\n", loja.Id)

	db, err := db.AbrirConexao()
	if err != nil {
		log.Println("Erro ao tentar abrir a conexão com o banco de dados")
		return err
	}
	defer db.Close()

	repositorio := repositories.NewRepositories(db)

	err = repositorio.AtualizarLoja(loja)
	if err != nil {
		log.Printf("Erro ao atualizar a loja com ID %d: %s\n", loja.Id, err)
		return err
	}

	log.Printf("Atualização da loja com ID %d concluída com sucesso\n", loja.Id)

	return nil
}

