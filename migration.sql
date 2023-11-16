CREATE TABLE public.estabelecimento (
	id bigserial NOT NULL,
	nome varchar(50) NOT NULL,
	razao_social varchar(100) NOT NULL,
	endereco varchar(50) NOT NULL,
	estado bpchar(2) NOT NULL,
	cidade varchar(50) NOT NULL,
	cep varchar(10) NOT NULL,
	numero_estabelecimento int8 NOT NULL,
	CONSTRAINT estabelecimento_pkey PRIMARY KEY (id)
);

CREATE TABLE public.loja (
	id bigserial NOT NULL,
	id_estabelecimento int8 NOT NULL,
	nome varchar(50) NOT NULL,
	razao_social varchar(100) NOT NULL,
	endereco varchar(50) NOT NULL,
	estado bpchar(2) NOT NULL,
	cidade varchar(50) NOT NULL,
	cep varchar(10) NOT NULL,
	CONSTRAINT loja_pkey PRIMARY KEY (id),
	CONSTRAINT fk_estabelecimento FOREIGN KEY (id_estabelecimento) REFERENCES public.estabelecimento(id)
);