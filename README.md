![logounb fw_](https://github.com/user-attachments/assets/035506cb-3c85-4ce6-8981-dde07042a663)
> Departamento de Ciência da Computação
> 
> Disciplina: Banco de Dados
---
# Consulta Candidatos 2024 BD

### Sumário
1. [Introdução](#Introdução)
2. [Estrutura da Tabela](#estrutura-da-tabela)
3. [Modelo Lógico](#modelo-logico)
4. [Tabelas](#tabelas)
5. [Como Populamos os Dados](#como-populamos-os-dados)
6. [Views](#views)
7. [Formas Normais](#formas-normais)
8. [Interface de Usuário](#interface-de-usuario)
---


## Introdução
Consiste em uma aplicação web com o uso de um banco de dados relacional para a visualização de Vereadores das Eleições Municipais de São Paulo capital. Para o desenvolvimento dessa ferramenta foi utilizado as tecnologias:
No Front-end:
- Linguagens HTML, CSS e JavaScript;
No servidor e no backend:
- Linguagem Go;
- SQL, com as soluçõs do SGBD PosgresSQL.
Assim criado 
 	
	

### Estrutura da Tabela
- Ferramenta utilizada [BR-MODELO](https://www.brmodeloweb.com/lang/pt-br/index.html)
# Modelo Entidade Relacionamento
![image](https://github.com/user-attachments/assets/95464d0a-0207-4523-9940-e856df8e9154)

# Modelo Lógico
![image](https://github.com/user-attachments/assets/994aeac6-ef9b-42e7-8d4a-3456d046c47f)

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

---


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
* Devemos ir no diretorio do PostgreSQL/bin e entrar com psql -U user -d database -h localhost pelo terminal
* Use esse comando para carregar o csv na nossa tabela "dados_eleitorais"
``` bash
\copy dados_eleitorais FROM 'C:\PATH\candidatos2024SP.csv' WITH (FORMAT csv, HEADER, DELIMITER ';', ENCODING 'UTF-8');
```
* Após isso devemos criar as nossas TABLES separadamente

``` sql
CREATE TABLE IF NOT EXISTS eleicao(
	foto BYTEA,
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
* Criamos as nossas TABLES separadas e temos uma TABLE geral contendo as informações do csv. Agora vamos transferir o conteúdo da nossa TABLE "Geral" para as Separadas
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

* Não precisamos mais da nossa TABLE dados_eleitorais, então utilize esse comando para excluí-la
```sql
DROP TABLE dados_eleitorais
```

# Views
```sql
CREATE VIEW view_genero_cargo_partido AS
SELECT 
    c.foto,
    c.nm_candidato AS nome,
    c.sq_candidato AS sequencia,
    c.ds_genero AS genero,
    e.ds_cargo AS cargo,
    p.sg_partido AS partido
FROM candidato c
JOIN eleicao e ON c.cd_eleicao = e.cd_eleicao
JOIN partido p ON c.nr_partido = p.nr_partido;



CREATE VIEW view_cor_cargo_partido AS
SELECT 
    c.foto,
    c.nm_candidato AS nome,
    c.sq_candidato AS sequencia,
    c.ds_cor_raca AS cor,
    e.ds_cargo AS cargo,
    p.sg_partido AS partido
FROM candidato c
JOIN eleicao e ON c.cd_eleicao = e.cd_eleicao
JOIN partido p ON c.nr_partido = p.nr_partido;



CREATE VIEW view_instrucao_cargo_partido AS
SELECT 
    c.foto,
    c.nm_candidato AS nome,
    c.sq_candidato AS sequencia,
    c.ds_grau_instrucao AS grau_instrucao,
    e.ds_cargo AS cargo,
    p.sg_partido AS partido
FROM candidato c
JOIN eleicao e ON c.cd_eleicao = e.cd_eleicao
JOIN partido p ON c.nr_partido = p.nr_partido;



CREATE VIEW view_federacao_turno_partido AS
SELECT 
    c.foto,
    c.nm_candidato AS nome,
    c.sq_candidato AS sequencia,
    f.sg_federacao AS federacao,
    e.nr_turno AS turno,
    p.sg_partido AS partido
FROM candidato c
JOIN partido p ON c.nr_partido = p.nr_partido
JOIN federacao f ON p.nr_federacao = f.nr_federacao
JOIN eleicao e ON c.cd_eleicao = e.cd_eleicao;



CREATE VIEW view_instrucao_ocupacao_data AS
SELECT
    c.foto,
    c.nm_candidato AS nome,
    c.sq_candidato AS sequencia,
    c.ds_grau_instrucao AS grau_instrucao,
    o.ds_ocupacao AS ocupacao,
    e.dt_eleicao AS data_eleicao
FROM candidato c
JOIN eleicao e ON c.cd_eleicao = e.cd_eleicao
JOIN ocupacao_candidato oc ON c.sq_candidato = oc.sq_candidato
JOIN ocupacao o ON oc.cd_ocupacao = o.cd_ocupacao;
```


---
## Caso você queira acessar de uma forma fácil estamos disponibilizando um arquivo chamado "Candidatos2024.sql", baixe ele [AQUI](https://drive.google.com/file/d/1buiqizU8WMNA5Q_rxh9dlg5I47-V7JLn/view?usp=sharing) e utilize este comando para carregar o arquivo no Postgresql:
```bash
psql -U usuario -d nome_do_banco -f /caminho/para/Candidatos2024.sql
```

---

# Formas Normais das Tabelas 
* Todas estão na (3FN), mas explicarei melhor nessas 5 tabelas:
#### eleicao
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
  
#### unidade eleitoral
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
  
#### federacao
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
  
#### candidato
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
  
#### telefone
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


# Fotos dos candidatos no banco de dados
- Estávamos na dúvida entre duas opções, usar um serviço de cloud como o aws da Amazon, ou inserir os dados diretamente no banco de dados com o Blob (Binary Large Object), cada um tem suas vantagens e desvantagens, se tratando de um projeto de candidatos onde existem dezenas de milhares de fotos, a cloud claramente era a melhor opção, ainda mais quando o plano gratuito da cloud ja seria o suficiente para armazenar nossos dados e depois era so salvar a url no banco, porém mesmo sabendo disso optamos por fazer com BYTEA (Equivalente ao Blob do PostgreSQL), já que se trata de um projeto acadêmico e não haverá utilização, decidimos que seria interessante mostrar aos outros alunos se um banco de dados com milhares de fotos armazenadas impactaria tanto assim no tempo de consulta.
### Na pasta scripts há o código em python desenvolvido para armazenar as fotos em binário no banco de dados
### Resultados
- Anteriormente a nossa consulta no banco de dados durava em torno de 2 segundos, após inserir o binário das fotos em cada candidato não só o tamanho do sql subiu exponencialmente (para 4GB), como o nosso tempo de consulta no banco de dados mudou para 25 segundos em média
- Consumo exagerado de memória RAM: Se a pesquisa for alguma coisa mais geral que vai resultar em muitos candidados, um notebook com 8GB de RAM não é capaz de suportar e quando a memoria chega em 100% a aplicação é interrompida pelo OS automáticamente, havendo casos onde o notebook simplesmente para de responder
- 8GB não é nada comparado ao poder de um servidor, porém é nítido que com a escolha de armazenamento errada, há um desperdício enorme de recursos, além do risco de instabilidade constante
  
---
# Controlador
- O controlador é constituído por um web server escrito em GO.
Basicamente, ele recebe uma requisição da interface (camada view) com dados de formulário. Esses dados são utilizados para que uma consulta seja realizada no banco de dados (camada model). Ao fim da consulta, o controlador realiza o tratamento dos dados para serem enviados de volta à interface (camada view).
```go
/Puts candidates info cards on the page
func executeTemplate(w http.ResponseWriter, page_struct pageCards) {
	templ, err := template.ParseFiles("pages/desktop.html")

	if err != nil {
		fmt.Println(err)
		return
	}
	
	templ.Execute(w, page_struct)
}

//Request handler for "/"
func indexHandler(w http.ResponseWriter, r* http.Request) {
	no_cards := pageCards{nil}
	executeTemplate(w, no_cards)
}

//Request handler for "/search"
func searchHandler(w http.ResponseWriter, r* http.Request) {
	if r.Method == "GET" {
		r.ParseForm()
	}
	
	var query_result [][]perfilCandidato
	var queryStr string
	
	//Query 1
	queryStr = createQueryStr(strings.ToUpper(r.Form["cor-raca"][0]), strings.ToUpper(r.Form["cargo"][0]), strings.ToUpper(r.Form["partido"][0]), 1)
	query_result = append(query_result, getFromView(queryStr))

	//Query 2
	queryStr = createQueryStr(strings.ToUpper(r.Form["federacao"][0]), strings.ToUpper(r.Form["turno"][0]), strings.ToUpper(r.Form["partido"][0]), 2)
	query_result = append(query_result, getFromView(queryStr))
	
	//Query 3
	queryStr = createQueryStr(strings.ToUpper(r.Form["genero"][0]), strings.ToUpper(r.Form["cargo"][0]), strings.ToUpper(r.Form["partido"][0]), 3)
	query_result = append(query_result, getFromView(queryStr))
	
	//Query 4
	queryStr = createQueryStr(strings.ToUpper(r.Form["instrucao"][0]), strings.ToUpper(r.Form["cargo"][0]), strings.ToUpper(r.Form["partido"][0]), 4)
	query_result = append(query_result, getFromView(queryStr))
	
	//Query 5
	queryStr = createQueryStr(strings.ToUpper(r.Form["instrucao"][0]), strings.ToUpper(r.Form["ocupacao"][0]), "", 5)
	query_result = append(query_result, getFromView(queryStr))
	
	cards := generateCards(query_result) 
	page_struct := pageCards{cards}	

	executeTemplate(w, page_struct)
}

func main() {
	//Creating http request handlers
	http.HandleFunc("/", indexHandler)
	http.HandleFunc("/search", searchHandler)
	
	//Creating image and audio files handler	
	http.Handle("/assets/", http.StripPrefix("/assets/", http.FileServer(http.Dir("assets"))))

	//Run http server on port 8080
	fmt.Println("Starting server on port 8080...")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
```
# Diagrama da camada de mapeamento
![image](https://github.com/user-attachments/assets/34de3447-30fc-40f8-b7c9-3c21f9ac885a)
```go
type perfilCandidato struct { 
Foto []byte 
Nome string 
Id string 
Atr1 string 
Atr2 string 
Atr3 string 
}
 type cardInfo struct { 
Imagem string 
Nome string 
Id string 
Partido string 
}
 type pageCards struct { 
Cards []cardInfo 
}
```

# Interface de Usuário
Dando foco aos elementos essenciais para uma interface de usuário que interagem com o banco de dados, temos essencialmente duas <div>s, uma conta um formulário (id= forms )com menus dropdowns que se asssemlham as [Tabelas](#tabelas), sendo o total de 12 menus dropdowns para a seleção de quantos vereadores devem ser exibidos para o usuário. Ao confirma as seleções e especificações da busca que o usuário deseja fazer, uma requisição é feita ao controler de Go, onde será feita atualização com as buscas desejadas especificadas nos slugs da URL. A outra div importante que conta é a de cartões (id=cards), onde um script na linguagem Go irá inteirar com cartões com as informações dos candidatos e inserindo na tela de usuário.

Desta forma, a interface do usuário busca oferecer uma maneira intuitiva ao público comum uma aplicação que permite a pesquisa dos candidatos a vereadores, com inclusive utilizanod de inspiração o design de uma urna eletrêonica e tendo interações sonoras com o famoso som de uma urna eletrônica. No final é possível deixar o email para que o usuário possa receber novas notícias sobre os candidatos da sua cidade.
 
