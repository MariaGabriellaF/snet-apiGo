package services

import (
	"errors"
	"log"
	"net/http"
	"snet-apiGo/src/models"
	"snet-apiGo/src/repositories"
	"snet-apiGo/src/repositories/db"

)

var (
	ErrNomeExcedeuLimite                 = errors.New("Excedeu o limite de caracteres")
)


func CriarEstabelecimento(w http.ResponseWriter, estabelecimento models.Estabelecimento) (*models.Estabelecimento, error) {

	if len(estabelecimento.Nome) > 50 {
		return nil, ErrNomeExcedeuLimite
	}
	if len(estabelecimento.RazaoSocial) > 100 {
		return nil, ErrNomeExcedeuLimite
	}
	if len(estabelecimento.Endereco) > 100 {
		return nil, ErrNomeExcedeuLimite
	}
	if len(estabelecimento.Estado) > 20 {
		return nil, ErrNomeExcedeuLimite
	}
	if len(estabelecimento.Cidade) > 50 {
		return nil, ErrNomeExcedeuLimite
	}

	db, err := db.AbrirConexao()
	if err != nil {
		log.Println("Erro ao tentar conexão db")
		return &models.Estabelecimento{}, err
	}
	defer db.Close()

	log.Println("Estabelecimento add número dentro service: ", estabelecimento.NumeroEstabelecimento)

	repositorio := repositories.NewRepositories(db)
	estabelecimento.Id, err = repositorio.NewEstabelecimento(estabelecimento)
	if err != nil {
		return &models.Estabelecimento{}, err
	}

	return &estabelecimento, nil
}

func GetEstabelecimentos(w http.ResponseWriter, r *http.Request) []models.Estabelecimento {
	log.Println("Iniciando a obtenção de estabelecimentos")

	db, err := db.AbrirConexao()
	if err != nil {
		log.Println("Erro ao abrir a conexão com o banco de dados:", err)
		return nil
	}
	defer db.Close()

	repositorio := repositories.NewRepositories(db)
	estabelecimentos, err := repositorio.ListarEstabelecimentos()
	if err != nil {
		log.Println("Erro ao obter estabelecimentos:", err)
		return []models.Estabelecimento{}
	}

	log.Println("Obtenção de estabelecimentos concluída com sucesso")

	return estabelecimentos
}

func GetEstabelecimentoPorID(id int) (*models.Estabelecimento, error) {
	log.Printf("Iniciando a obtenção do estabelecimento com ID %d\n", id)

	db, err := db.AbrirConexao()
	if err != nil {
		log.Println("Erro ao abrir a conexão com o banco de dados:", err)
		return &models.Estabelecimento{}, err
	}
	defer db.Close()

	repositorio := repositories.NewRepositories(db)
	estabelecimento, err := repositorio.ListarEstabelecimentoPorID(id)
	if err != nil {
		log.Printf("Erro ao obter o estabelecimento com ID %d: %s\n", id, err)
		return &models.Estabelecimento{}, err
	}

	log.Printf("Obtenção do estabelecimento com ID %d concluída com sucesso\n", id)

	return &estabelecimento, nil
}

func DeletarEstabelecimentoPorID(id int) error {
	log.Printf("Iniciando a exclusão do estabelecimento com ID %d\n", id)

	db, err := db.AbrirConexao()
	if err != nil {
		log.Println("Erro ao abrir a conexão com o banco de dados:", err)
		return err
	}
	defer db.Close()

	repositorio := repositories.NewRepositories(db)
	log.Println("erro ao pegar o repositorio")
	err = repositorio.DeletarEstabelecimentoPorID(id)
	if err != nil {
		log.Printf("Erro ao excluir o estabelecimento com ID %d: %s\n", id, err)
		return err
	}
	
	log.Printf("Exclusão do estabelecimento com ID %d concluída com sucesso\n", id)

	return nil
}

func AtualizarEstabelecimentoPorID(estabelecimento models.Estabelecimento) error {
	log.Printf("Iniciando a atualização do estabelecimento com ID %d\n", estabelecimento.Id)

	db, err := db.AbrirConexao()
	if err != nil {
		log.Println("Erro ao abrir a conexão com o banco de dados:", err)
		return err
	}
	defer db.Close()

	repositorio := repositories.NewRepositories(db)
	err = repositorio.AtualizarEstabelecimento(estabelecimento)
	if err != nil {
		log.Printf("Erro ao atualizar o estabelecimento com ID %d: %s\n", estabelecimento.Id, err)
		return err
	}

	log.Printf("Atualização do estabelecimento com ID %d concluída com sucesso\n", estabelecimento.Id)

	return nil
}

func ListarLojasDoEstabelecimentoPorID(id int) (models.Estabelecimento, error) {
	log.Printf("Iniciando a obtenção de estabelecimento com lojas para o ID %d\n", id)

	db, err := db.AbrirConexao()
	if err != nil {
		log.Println("Erro ao abrir a conexão com o banco de dados:", err)
		return models.Estabelecimento{}, err
	}
	defer db.Close()

	repositorio := repositories.NewRepositories(db)
	estabelecimentos, err := repositorio.GetEstabelecimentoComLojas(id)
	if err != nil {
		log.Printf("Erro ao obter estabelecimento com lojas para o ID %d: %s\n", id, err)
		return models.Estabelecimento{}, err
	}

	log.Printf("Estabelecimento com lojas (ID %d) obtido com sucesso!\n", id)


	return estabelecimentos, nil
}
