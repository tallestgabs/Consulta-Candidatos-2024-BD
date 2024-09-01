# Consulta-Candidatos-2024-BD

### Estrutura da Tabela
- Alterar para as tabelas do MER

![MER](https://github.com/user-attachments/assets/555ced09-f2e8-4c90-81b6-d0fe8d0b3f51)

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
    DT_NASCIMENTO VARCHAR(12),                  -- Data de nascimento
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

### Entrar no diretorio do PostgreSQL/bin e entrar com psql -U user -d database -h localhost
``` bash
\copy dados_eleitorais FROM 'C:\Users\Gabriel\Downloads\consulta_cand_2024_SP.csv' WITH (FORMAT csv, HEADER, DELIMITER ';');
```



