# Consulta-Candidatos-2024-BD

### Estrutura da Tabela
- Ferramenta utilizada [BR-MODELO](https://www.brmodeloweb.com/lang/pt-br/index.html)
# Modelo Entidade Relacionamento
![image](https://github.com/user-attachments/assets/37106aeb-acff-4f5a-90f1-bc5db4ee6f45)
# Modelo Lógico
![image](https://github.com/user-attachments/assets/534ae217-7789-482b-8c3e-26ede86a3a70)







## Tabelas 
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



# Como populamos os Dados
- Primeiro devemos criar uma TABLE contendo todo o conteúdo do nosso csv:
  
``` sql

CREATE TABLE IF NOT EXISTS dados_eleitorais (
DT_GERACAO VARCHAR(12),                  -- Data de geração
HH_GERACAO varchar(12),                  -- Hora de geração
ANO_ELEICAO VARCHAR(5),                  -- Ano da eleição
CD_TIPO_ELEICAO VARCHAR(5),              -- Código do tipo de eleição
NM_TIPO_ELEICAO VARCHAR(50),             -- Nome do tipo de eleição
NR_TURNO VARCHAR(2),                     -- Número do turno
CD_ELEICAO VARCHAR(5),                   -- Código da eleição
DS_ELEICAO VARCHAR(100),                 -- Descrição da eleição
DT_ELEICAO VARCHAR(12),                  -- Data da eleição
TP_ABRANGENCIA_ELEICAO VARCHAR(50),      -- Tipo de abrangência da eleição
SG_UF VARCHAR(2),                        -- Sigla da Unidade Federativa
SQ_UE VARCHAR(7),                        -- Sigla da Unidade Eleitoral
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
SG_FEDERACAO VARCHAR(20),                -- Sigla da federação
DS_COMPOSICAO_FEDERACAO TEXT,            -- Descrição da composição da federação
SQ_COLIGACAO VARCHAR(20),                -- Sequência da coligação
NM_COLIGACAO VARCHAR(100),               -- Nome da coligação
DS_COMPOSICAO_COLIGACAO TEXT,            -- Descrição da composição da coligação
SG_UF_NASCIMENTO VARCHAR(2),             -- Sigla da UF de nascimento
DT_NASCIMENTO VARCHAR(12),               -- Data de nascimento
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
DS_SIT_TOT_TURNO VARCHAR(100),           -- Descrição da situação do turno
email TEXT,                              -- Exemplos de Emails
telefone TEXT                            -- Exemplos de Telefones
);
```
### Devemos ir no diretorio do PostgreSQL/bin e entrar com psql -U user -d database -h localhost pelo terminal
### Use esse comando para carregar o csv na nossa tabela "dados_eleitorais"
``` bash
\copy dados_eleitorais FROM 'C:\PATH\candidatos2024SP.csv' WITH (FORMAT csv, HEADER, DELIMITER ';', ENCODING 'UTF-8');
```
### Após isso devemos criar as nossas TABLES separadamente

``` sql
CREATE TABLE IF NOT EXISTS eleicao(
	cd_eleicao VARCHAR(5) PRIMARY KEY, 
	ds_eleicao VARCHAR(100),
        ds_cargo VARCHAR(100),
	dt_eleicao VARCHAR(12),
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
	dt_nascimento VARCHAR(12),
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
	telefones TEXT,
	FOREIGN KEY (sq_candidato) REFERENCES candidato(sq_candidato)
);
CREATE TABLE IF NOT EXISTS email(
	sq_candidato VARCHAR(20),
	emails TEXT,
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
### Criamos as nossas TABLES separadas e temos uma TABLE geral contendo as informações do csv. Agora vamos transferir o conteudo da nossa TABLE "Geral" para as Separadas
``` sql
INSERT INTO eleicao (cd_eleicao, ds_eleicao, ds_cargo, dt_eleicao, tp_abrangencia_eleicao, ano_eleicao, cd_tipo_eleicao, nm_tipo_eleicao, nr_turno)
SELECT cd_eleicao, ds_eleicao, ds_cargo, dt_eleicao, tp_abrangencia_eleicao, ano_eleicao, cd_tipo_eleicao, nm_tipo_eleicao, nr_turno
FROM dados_eleitorais
ON CONFLICT DO NOTHING;


INSERT INTO unidade_eleitoral (
    cd_eleicao, sq_ue, nm_ue, sg_uf
)
SELECT 
    cd_eleicao, sq_ue, nm_ue, sg_uf
FROM dados_eleitorais
ON CONFLICT DO NOTHING;


INSERT INTO federacao (
    nr_federacao, nm_federacao, sg_federacao, ds_composicao_federacao
)
SELECT 
    nr_federacao, nm_federacao, sg_federacao, ds_composicao_federacao
FROM dados_eleitorais
ON CONFLICT DO NOTHING;


INSERT INTO coligacao (
    sq_coligacao, nm_coligacao, ds_composicao_coligacao
)
SELECT 
    sq_coligacao, nm_coligacao, ds_composicao_coligacao
FROM dados_eleitorais
ON CONFLICT DO NOTHING;


INSERT INTO partido (
    nr_federacao, sq_coligacao, nr_partido, sg_partido, nm_partido
)
SELECT 
    nr_federacao, sq_coligacao, nr_partido, sg_partido, nm_partido
FROM dados_eleitorais
ON CONFLICT DO NOTHING;


INSERT INTO candidato (
    cd_eleicao, nr_partido, sq_candidato, nm_candidato, ds_genero, ds_cor_raca, 
    dt_nascimento, nr_titulo_eleitoral_candidato, nm_urna_candidato, 
    nr_candidato, ds_estado_civil, cd_situacao_candidatura, ds_grau_instrucao
)
SELECT 
    cd_eleicao, nr_partido, sq_candidato, nm_candidato, ds_genero, ds_cor_raca, 
    dt_nascimento, nr_titulo_eleitoral_candidato, nm_urna_candidato, 
    nr_candidato, ds_estado_civil, cd_situacao_candidatura, ds_grau_instrucao
FROM dados_eleitorais
ON CONFLICT DO NOTHING;


INSERT INTO telefone (
    sq_candidato, telefones
)
SELECT 
    sq_candidato, telefone
FROM dados_eleitorais
ON CONFLICT DO NOTHING;


INSERT INTO email (
    sq_candidato, emails
)
SELECT 
    sq_candidato, email
FROM dados_eleitorais
ON CONFLICT DO NOTHING;


INSERT INTO ocupacao (
    cd_ocupacao, ds_ocupacao
)
SELECT 
    cd_ocupacao, ds_ocupacao
FROM dados_eleitorais
ON CONFLICT DO NOTHING;


INSERT INTO ocupacao_candidato (
    sq_candidato, cd_ocupacao
)
SELECT 
    sq_candidato, cd_ocupacao
FROM dados_eleitorais
ON CONFLICT DO NOTHING;

```

## Não precisamos mais da nossa TABLE dados_eleitorais, então utilize esse comando para excluí-la
```sql
DROP TABLE dados_eleitorais
```

# Formas Normais das Tabelas 
### Todas estão na (3FN), mas explicarei melhor nessas 5 tabelas:
## eleicao
```sql
CREATE TABLE IF NOT EXISTS eleicao(
	cd_eleicao VARCHAR(5) PRIMARY KEY, 
	ds_eleicao VARCHAR(100),
        ds_cargo VARCHAR(100),
	dt_eleicao VARCHAR(12),
	tp_abrangencia_eleicao VARCHAR(50),
	ano_eleicao VARCHAR(5),
	cd_tipo_eleicao VARCHAR(5),
	nm_tipo_eleicao VARCHAR(50),
	nr_turno VARCHAR(2)
);
```
- 1FN: Está na 1FN porque todos os valores são atômicos (não há listas ou conjuntos de valores em uma única célula).
- 2FN: Está na 2FN porque a chave primária (cd_eleicao) é um atributo único e todos os outros atributos dependem completamente dela.
- 3FN: Está na 3FN porque não há dependências transitivas entre os atributos não chave. Todos os atributos dependem diretamente da chave primária cd_eleicao.
  
## unidade eleitoral
```sql
CREATE TABLE IF NOT EXISTS unidade_eleitoral (
	cd_eleicao VARCHAR(5),
	sq_ue VARCHAR(7) PRIMARY KEY,
	nm_ue VARCHAR(40),
	sg_uf VARCHAR(2),
	FOREIGN KEY (cd_eleicao) REFERENCES eleicao(cd_eleicao)
);
```
- 1FN: Está na 1FN porque os valores são atômicos.
- 2FN: Está na 2FN porque (sq_ue) é a chave primária, e todos os atributos dependem dessa chave. 
- 3FN: Está na 3FN, pois não há dependência transitiva entre os atributos não chave. Todos os atributos (como nm_ue e sg_uf) dependem diretamente da chave primária (sq_ue).
  
## federacao
```sql
CREATE TABLE IF NOT EXISTS federacao (
	nr_federacao VARCHAR(7) PRIMARY KEY,  
	nm_federacao VARCHAR(100),      
	sg_federacao VARCHAR(20),        
	ds_composicao_federacao TEXT
);
```
- 1FN: Está na 1FN porque os valores são atômicos.
- 2FN: Está na 2FN, pois todos os atributos dependem completamente da chave primária (nr_federacao).
- 3FN: Está na 3FN, já que não há dependências transitivas entre os atributos não chave.
  
## candidato
```sql
CREATE TABLE IF NOT EXISTS candidato (
	cd_eleicao VARCHAR(5),
	nr_partido VARCHAR(3),
	sq_candidato VARCHAR(20) PRIMARY KEY,
	nm_candidato VARCHAR(100),
	ds_genero VARCHAR(25),
	ds_cor_raca VARCHAR(20),
	dt_nascimento VARCHAR(12),
	nr_titulo_eleitoral_candidato VARCHAR(15),
	nm_urna_candidato VARCHAR(100),
	nr_candidato VARCHAR(10),
	ds_estado_civil VARCHAR(25),
	cd_situacao_candidatura VARCHAR(4),
	ds_grau_instrucao VARCHAR(50),
	FOREIGN KEY (cd_eleicao) REFERENCES eleicao(cd_eleicao),
	FOREIGN KEY (nr_partido) REFERENCES partido(nr_partido)
);
```
- 1FN: Está na 1FN porque os valores são atômicos.
- 2FN: Está na 2FN porque a chave primária é (sq_candidato), e todos os atributos dependem completamente dela.
- 3FN: Está na 3FN, pois não há dependências transitivas entre os atributos não chave. Todos os atributos (como nm_candidato, ds_genero, ds_cor_raca) dependem diretamente da chave primária (sq_candidato).
  
## telefone
```sql
CREATE TABLE IF NOT EXISTS telefone(
	sq_candidato VARCHAR(20),
	telefones TEXT,
	FOREIGN KEY (sq_candidato) REFERENCES candidato(sq_candidato)
);
```
- 1FN: Está na 1FN, pois os valores são atômicos.
- 2FN: Está na 2FN, já que todos os atributos dependem diretamente da chave estrangeira (sq_candidato). A tabela é uma normalização da relação multivalorada de telefone.
- 3FN: Está na 3FN, pois não há dependências transitivas. O único atributo é telefones, que depende diretamente de (sq_candidato).

