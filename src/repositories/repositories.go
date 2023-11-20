package repositories

import (
	"database/sql"
	"errors"
	"log"
	"snet-apiGo/src/models"
)

type repositoriesDB struct {
	db *sql.DB
}

func NewRepositories(db *sql.DB) *repositoriesDB {
	return &repositoriesDB{db}
}

var (
	ErrAoApagarEstabelecimento = errors.New("O estabelecimento não pode ser apagado, pois existe lojas assosciadas.")
	ErrVerificarLojasAssociadas = errors.New("Erro ao verificar lojas.")
)

// Estabelecimento
func (r repositoriesDB) NewEstabelecimento(estabelecimento models.Estabelecimento) (int, error) {
	stmt, err := r.db.Prepare("insert into estabelecimento (nome, razao_social, endereco, estado, cidade, cep, numero_estabelecimento) " +
		"values ($1, $2, $3, $4, $5, $6, $7 ) returning id")

	if err != nil {
		log.Println("Erro ao preparar a declaração SQL:", err)
		return 0, err
	}
	defer stmt.Close()

	var id int
	if err = stmt.QueryRow(estabelecimento.Nome, estabelecimento.RazaoSocial, estabelecimento.Endereco, estabelecimento.Estado, estabelecimento.Cidade, estabelecimento.Cep, estabelecimento.NumeroEstabelecimento).Scan(&id); err != nil {
		log.Println("Erro ao executar a consulta SQL:", err)
		return 0, err
	}

	log.Printf("Novo estabelecimento criado com sucesso. ID: %d\n", id)
	return id, nil
}

func (r repositoriesDB) ListarEstabelecimentos() ([]models.Estabelecimento, error) {
	rows, err := r.db.Query("select * from estabelecimento")
	if err != nil {
		log.Println("Erro ao executar a consulta SQL:", err)
		return nil, err
	}
	defer rows.Close()

	var estabelecimentos []models.Estabelecimento
	for rows.Next() {
		var estabelecimento models.Estabelecimento

		if err = rows.Scan(&estabelecimento.Id, &estabelecimento.Nome, &estabelecimento.RazaoSocial, &estabelecimento.Endereco, &estabelecimento.Estado, &estabelecimento.Cidade, &estabelecimento.Cep, &estabelecimento.NumeroEstabelecimento); err != nil {
			log.Println("Erro ao ler os resultados da consulta SQL:", err)
			return nil, err
		}
		estabelecimentos = append(estabelecimentos, estabelecimento)
	}

	log.Printf("Listagem de estabelecimentos concluída. Total de estabelecimentos: %d\n", len(estabelecimentos))
	return estabelecimentos, nil
}

func (r repositoriesDB) ListarEstabelecimentoPorID(id int) (models.Estabelecimento, error) {
	var estabelecimento models.Estabelecimento

	row := r.db.QueryRow("Select * from estabelecimento where id = $1", id)

	// Pega os dados para o objeto Estabelecimento
	err := row.Scan(&estabelecimento.Id, &estabelecimento.Nome, &estabelecimento.RazaoSocial, &estabelecimento.Endereco, &estabelecimento.Estado, &estabelecimento.Cidade, &estabelecimento.Cep, &estabelecimento.NumeroEstabelecimento)

	if err != nil {
		log.Printf("Erro ao consultar o estabelecimento com ID %d: %s\n", id, err)
		return models.Estabelecimento{}, err
	}

	log.Printf("Consulta do estabelecimento com ID %d concluída. Dados do estabelecimento: %+v\n", id, estabelecimento)

	return estabelecimento, nil
}


func (r repositoriesDB) DeletarEstabelecimentoPorID(id int) error {
	// Verificar se existem lojas associadas ao estabelecimento
	var count int
	err := r.db.QueryRow("SELECT COUNT(*) FROM loja WHERE id_estabelecimento = $1", id).Scan(&count)
	if err != nil {
		log.Printf("Erro ao verificar lojas associadas ao estabelecimento com ID %d: %s\n", id, err)
		return ErrVerificarLojasAssociadas
	}

	if count > 0 {
		return ErrAoApagarEstabelecimento
	}

	// Se não houver lojas associadas, proceder com a exclusão do estabelecimento
	_, err = r.db.Exec("DELETE FROM estabelecimento WHERE id = $1", id)
	if err != nil {
		log.Printf("Erro ao excluir o estabelecimento com ID %d: %s\n", id, err)
		return ErrAoApagarEstabelecimento
	}

	log.Printf("Exclusão do estabelecimento com ID %d concluída com sucesso\n", id)
	return nil
}

func (r repositoriesDB) AtualizarEstabelecimento(estabelecimento models.Estabelecimento) error {
	_, err := r.db.Exec("UPDATE estabelecimento SET nome = $1, razao_social = $2, endereco = $3, estado = $4, cidade = $5, cep = $6, numero_estabelecimento = $7 WHERE id = $8",
		estabelecimento.Nome, estabelecimento.RazaoSocial, estabelecimento.Endereco, estabelecimento.Estado, estabelecimento.Cidade, estabelecimento.Cep, estabelecimento.NumeroEstabelecimento, estabelecimento.Id)

	if err != nil {
		return err
	}

	return nil
}

//Loja

func (r repositoriesDB) NewLoja(loja models.Loja) (int, error) {
	stmt, err := r.db.Prepare("insert into loja (id_estabelecimento, nome, razao_social, endereco, estado, cidade, cep) " +
		"values ($1, $2, $3, $4, $5, $6, $7 ) returning id")
	if err != nil {
		log.Println("Erro ao preparar a declaração SQL:", err)
		return 0, err
	}
	defer stmt.Close()

	var id int
	if err = stmt.QueryRow(loja.IdEstabelecimento, loja.Nome, loja.RazaoSocial, loja.Endereco, loja.Estado, loja.Cidade, loja.Cep).Scan(&id); err != nil {
		log.Println("Erro ao executar a consulta SQL:", err)
		return 0, err
	}

	log.Printf("Nova loja criada com sucesso. ID: %d\n", id)

	return id, nil
}

func (r repositoriesDB) ListarLojas() ([]models.Loja, error) {
	rows, err := r.db.Query("select * from loja")
	if err != nil {
		log.Println("Erro ao executar a consulta SQL:", err)
		return nil, err
	}

	defer rows.Close()

	var lojas []models.Loja
	for rows.Next() {
		var loja models.Loja

		if err = rows.Scan(&loja.Id, &loja.IdEstabelecimento, &loja.Nome, &loja.RazaoSocial, &loja.Endereco, &loja.Estado, &loja.Cidade, &loja.Cep); err != nil {
			log.Println("Erro ao escanear os resultados da consulta SQL:", err)
			return nil, err
		}
		lojas = append(lojas, loja)
	}

	log.Printf("Listagem de lojas concluída. Total de lojas: %d\n", len(lojas))

	return lojas, nil
}

func (r repositoriesDB) DeletarLojaPorID(id int) error {
	_, err := r.db.Exec("delete from loja where id = $1", id)

	if err != nil {
		log.Printf("Erro ao excluir a loja com ID %d: %s\n", id, err)
		return err
	}

	log.Printf("Exclusão da loja com ID %d concluída com sucesso\n", id)

	return nil
}

func (r repositoriesDB) AtualizarLoja(loja models.Loja) error {
	_, err := r.db.Exec("update loja set nome = $1, razao_social = $2, endereco = $3, estado = $4, cidade = $5, cep = $6 WHERE id = $7",
		loja.Nome, loja.RazaoSocial, loja.Endereco, loja.Estado, loja.Cidade, loja.Cep, loja.Id)

	if err != nil {
		return err
	}

	return nil
}

func (r repositoriesDB) GetEstabelecimentoComLojas(id int) (models.Estabelecimento, error) {
	query := `
		SELECT
			e.id, e.nome, e.razao_social, e.endereco, e.estado, e.cidade, e.cep, e.numero_estabelecimento,
			l.Id AS id_loja, l.nome AS nome_loja, l.razao_social AS razao_social_loja, l.endereco AS endereco_loja, l.estado AS estado_loja, l.cidade AS cidade_loja, l.cep AS cep_loja
		FROM Estabelecimento e
		INNER JOIN Loja l ON e.Id = l.id_estabelecimento
		WHERE e.Id = $1
	`
	rows, err := r.db.Query(query, id)
	if err != nil {
		log.Println("Erro ao executar a consulta SQL:", err)
		return models.Estabelecimento{}, err
	}
	defer rows.Close()

	log.Println("Consulta SQL concluída com sucesso.")

	var estabelecimento models.Estabelecimento
	lojas := make(map[int]models.Loja)

	for rows.Next() {
		var idEst, idLoja int
		var nomeEst, nomeLoja, razaoSocialEst, razaoSocialLoja, enderecoEst, enderecoLoja, estadoEst, estadoLoja, cidadeEst, cidadeLoja, cepEst, cepLoja string
		var numeroEstabelecimento int

		if err = rows.Scan(&idEst, &nomeEst, &razaoSocialEst, &enderecoEst, &estadoEst, &cidadeEst, &cepEst, &numeroEstabelecimento, &idLoja, &nomeLoja, &razaoSocialLoja, &enderecoLoja, &estadoLoja, &cidadeLoja, &cepLoja); err != nil {
			log.Println("Erro ao escanear os resultados da consulta SQL:", err)
			return models.Estabelecimento{}, err
		}

		if estabelecimento.Id == 0 {
			estabelecimento = models.Estabelecimento{
				Id:                    idEst,
				Nome:                  nomeEst,
				RazaoSocial:           razaoSocialEst,
				Endereco:              enderecoEst,
				Estado:                estadoEst,
				Cidade:                cidadeEst,
				Cep:                   cepEst,
				NumeroEstabelecimento: numeroEstabelecimento,
			}
		}

		if _, ok := lojas[idLoja]; !ok {
			lojas[idLoja] = models.Loja{
				Id:                idLoja,
				IdEstabelecimento: idEst,
				Nome:              nomeLoja,
				RazaoSocial:       razaoSocialLoja,
				Endereco:          enderecoLoja,
				Estado:            estadoLoja,
				Cidade:            cidadeLoja,
				Cep:               cepLoja,
			}
		}
	}

	for _, loja := range lojas {
		estabelecimento.Lojas = append(estabelecimento.Lojas, loja)
	}

	log.Println("Obtenção de estabelecimento com lojas concluída com sucesso")

	return estabelecimento, nil
}
