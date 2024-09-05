# Consulta-Candidatos-2024-BD

### Estrutura da Tabela
![image](https://github.com/user-attachments/assets/37106aeb-acff-4f5a-90f1-bc5db4ee6f45)


``` sql
CREATE TABLE IF NOT EXISTS dados_eleitorais (
    DT_GERACAO VARCHAR(12),                  -- Data de geração
    HH_GERACAO TIME,                         -- Hora de geração
    ANO_ELEICAO VARCHAR(5),                  -- Ano da eleição
    CD_TIPO_ELEICAO VARCHAR(5),              -- Código do tipo de eleição
    NM_TIPO_ELEICAO VARCHAR(50),             -- Nome do tipo de eleição
    NR_TURNO VARCHAR(2),                     -- Número do turno
    CD_ELEICAO VARCHAR(5),                   -- Código da eleição
    DS_ELEICAO VARCHAR(100),                 -- Descrição da eleição
    DT_ELEICAO DATE,                         -- Data da eleição
    TP_ABRANGENCIA_ELEICAO VARCHAR(50),      -- Tipo de abrangência da eleição
    SG_UF VARCHAR(2),                        -- Sigla da Unidade Federativa
    SG_UE VARCHAR(7),                        -- Sigla da Unidade Eleitoral
    NM_UE VARCHAR(40),                       -- Nome da Unidade Eleitoral
    CD_CARGO VARCHAR(5),                     -- Código do cargo
    DS_CARGO VARCHAR(100),                   -- Descrição do cargo
    SQ_CANDIDATO VARCHAR(20),                -- Sequência do candidato
    NR_CANDIDATO VARCHAR(10),                -- Número do candidato
    NM_CANDIDATO VARCHAR(100),               -- Nome do candidato
    NM_URNA_CANDIDATO VARCHAR(100),          -- Nome do candidato na urna
    NM_SOCIAL_CANDIDATO VARCHAR(100),        -- Nome social do candidato
    NR_CPF_CANDIDATO VARCHAR(4),             -- CPF do candidato (Não Divulgado)
    DS_EMAIL VARCHAR(70),                    -- E-mail do candidato (Não Divulgado)
    CD_SITUACAO_CANDIDATURA VARCHAR(4),      -- Código da situação da candidatura
    DS_SITUACAO_CANDIDATURA VARCHAR(5),      -- Descrição da situação da candidatura
    TP_AGREMIACAO VARCHAR(50),               -- Tipo de agremiação
    NR_PARTIDO VARCHAR(3),                   -- Número do partido
    SG_PARTIDO VARCHAR(30),                  -- Sigla do partido
    NM_PARTIDO VARCHAR(100),                 -- Nome do partido
    NR_FEDERACAO VARCHAR(7),                 -- Número da federação
    NM_FEDERACAO VARCHAR(100),               -- Nome da federação
    SG_FEDERACAO VARCHAR(20),                 -- Sigla da federação
    DS_COMPOSICAO_FEDERACAO TEXT,            -- Descrição da composição da federação
    SQ_COLIGACAO VARCHAR(20),                -- Sequência da coligação
    NM_COLIGACAO VARCHAR(100),               -- Nome da coligação
    DS_COMPOSICAO_COLIGACAO TEXT,            -- Descrição da composição da coligação
    SG_UF_NASCIMENTO VARCHAR(2),             -- Sigla da UF de nascimento
    DT_NASCIMENTO DATE,                      -- Data de nascimento
    NR_TITULO_ELEITORAL_CANDIDATO VARCHAR(15),  -- Número do título de eleitor do candidato
    CD_GENERO VARCHAR(2),                       -- Código do gênero
    DS_GENERO VARCHAR(25),                   -- Descrição do gênero
    CD_GRAU_INSTRUCAO VARCHAR(1),            -- Código do grau de instrução
    DS_GRAU_INSTRUCAO VARCHAR(50),           -- Descrição do grau de instrução
    CD_ESTADO_CIVIL VARCHAR(2),              -- Código do estado civil
    DS_ESTADO_CIVIL VARCHAR(25),             -- Descrição do estado civil
    CD_COR_RACA VARCHAR(3),                  -- Código da cor/raça
    DS_COR_RACA VARCHAR(25),                 -- Descrição da cor/raça
    CD_OCUPACAO VARCHAR(5),                  -- Código da ocupação
    DS_OCUPACAO VARCHAR(100),                -- Descrição da ocupação
    CD_SIT_TOT_TURNO VARCHAR(8),             -- Código da situação do turno
    DS_SIT_TOT_TURNO VARCHAR(100)            -- Descrição da situação do turno

);
```
### Tabelas Separadas
1. Eleição
2. Unidade Eleitoral
3. Candidato
4. Email
5. Telefone
6. Ocupação
7. Ocupação_Candidato
8. Partido
9. Federação
10. Coligação

``` sql
CREATE TABLE IF NOT EXISTS eleicao(
	cd_eleicao VARCHAR(5) PRIMARY KEY, 
	ds_eleicao VARCHAR(100),
	dt_eleicao DATE,
	tp_abrangencia_eleicao VARCHAR(50),
	ano_eleicao VARCHAR(5),
	cd_tipo_eleicao VARCHAR(5),
	nm_tipo_eleicao VARCHAR(50),
	nr_turno VARCHAR(2)
);
CREATE TABLE IF NOT EXISTS unidade_eleitoral (
	cd_eleicao VARCHAR(5),
	sq_ue VARCHAR(7) PRIMARY KEY,
	nm_ue VARCHAR(40),
	sg_uf VARCHAR(2),
	FOREIGN KEY (cd_eleicao) REFERENCES eleicao(cd_eleicao)
);
CREATE TABLE IF NOT EXISTS federacao (
	nr_federacao VARCHAR(7) PRIMARY KEY,  
	nm_federacao VARCHAR(100),      
	sg_federacao VARCHAR(20),        
	ds_composicao_federacao TEXT
);

CREATE TABLE IF NOT EXISTS coligacao (
	sq_coligacao VARCHAR(20) PRIMARY KEY,          
	nm_coligacao VARCHAR(100),              
	ds_composicao_coligacao TEXT
);
CREATE TABLE IF NOT EXISTS partido (
	nr_federacao VARCHAR(7),
	sq_coligacao VARCHAR(20),
	nr_partido VARCHAR(3) PRIMARY KEY,
	sg_partido VARCHAR(30),
	nm_partido VARCHAR(100),
	FOREIGN KEY (nr_federacao) REFERENCES federacao(nr_federacao),
	FOREIGN KEY (sq_coligacao) REFERENCES coligacao(sq_coligacao)
);

CREATE TABLE IF NOT EXISTS candidato (
	cd_eleicao VARCHAR(5),
	nr_partido VARCHAR(3),
	sq_candidato VARCHAR(20) PRIMARY KEY,
	nm_candidato VARCHAR(100),
	ds_genero VARCHAR(25),
	ds_cor_raca VARCHAR(20),
	dt_nascimento DATE,
	nr_titulo_eleitoral_candidato VARCHAR(15),
	nm_urna_candidato VARCHAR(100),
	nr_candidato VARCHAR(10),
	ds_estado_civil VARCHAR(25),
	cd_situacao_candidatura VARCHAR(4),
	ds_grau_instrucao VARCHAR(50),
	FOREIGN KEY (cd_eleicao) REFERENCES eleicao(cd_eleicao),
	FOREIGN KEY (nr_partido) REFERENCES partido(nr_partido)
);

CREATE TABLE IF NOT EXISTS telefone(
	sq_candidato VARCHAR(20),
	numero VARCHAR(15),
	FOREIGN KEY (sq_candidato) REFERENCES candidato(sq_candidato)
);
CREATE TABLE IF NOT EXISTS email(
	sq_candidato VARCHAR(20),
	ds_email VARCHAR(50),
	FOREIGN KEY (sq_candidato) REFERENCES candidato(sq_candidato)
);
CREATE TABLE IF NOT EXISTS ocupacao(
	cd_ocupacao VARCHAR(10) PRIMARY KEY,
	ds_ocupacao VARCHAR(255)
);
CREATE TABLE IF NOT EXISTS ocupacao_candidato(
	sq_candidato VARCHAR(20),
	cd_ocupacao VARCHAR(255),
	PRIMARY KEY (sq_candidato, cd_ocupacao),
	FOREIGN KEY (sq_candidato) REFERENCES candidato(sq_candidato),
	FOREIGN KEY (cd_ocupacao) REFERENCES ocupacao(cd_ocupacao)
);
```
### Entrar no diretorio do PostgreSQL/bin e entrar com psql -U user -d database -h localhost
``` bash
\copy dados_eleitorais FROM 'C:\Users\Gabriel\Downloads\consulta_cand_2024_SP.csv' WITH (FORMAT csv, HEADER, DELIMITER ';');
```



