# Neoway Backend Test

![Go](https://img.shields.io/badge/Go-1.20-blue)
![Gin](https://img.shields.io/badge/Gin-Framework-brightgreen)
![PostgreSQL](https://img.shields.io/badge/PostgreSQL-Database-blue)
![Docker](https://img.shields.io/badge/Docker-Containerization-blue)
![GORM](https://img.shields.io/badge/GORM-ORM-orange)

Este projeto é uma implementação de uma API RESTful para cadastro, listagem e consulta de clientes (PF e PJ) de uma empresa fictícia. Também é possível verificar a existência de um determinado CPF/CNPJ na base e obter informações de uptime do servidor.

## Sumário

- [Descrição](#descrição)
- [Requisitos Funcionais](#requisitos-funcionais)
- [Arquitetura](#arquitetura)
- [Tecnologias Utilizadas](#tecnologias-utilizadas)
- [Instalação e Configuração](#instalação-e-configuração)
- [Importação do CSV](#importação-do-csv)
- [Execução](#execução)
- [Rotas da API](#rotas-da-api)
- [Validação de Documentos (CPF/CNPJ)](#validação-de-documentos-cpfcnpj)
- [Testes e Cobertura](#testes-e-cobertura)
- [Docker e Docker Compose](#docker-e-docker-compose)
- [Melhorias Futuras](#melhorias-futuras)

---

## Descrição

O projeto visa importar dados de clientes a partir de um arquivo CSV, cadastrar clientes (pessoas físicas ou jurídicas) em um banco de dados PostgreSQL, listar todos os clientes, consultar clientes por documento (CPF/CNPJ), filtrar clientes por nome e verificar se um documento já existe. Além disso, há uma rota que fornece informações sobre o uptime do servidor e a contagem de requisições.

---

## Requisitos Funcionais

- **Cadastro de clientes** utilizando CPF/CNPJ como chave, armazenando nome/razão social e se o cliente está bloqueado ou não.
- **Consulta de todos os clientes** cadastrados, com opção de busca por nome e ordenação alfabética.
- **Consulta de um cliente pelo CPF/CNPJ**.
- **Verificação de existência de um determinado CPF/CNPJ** na base de clientes.
- **Validação de CPF/CNPJ** na consulta e inclusão.
- **Rota `/status`** que retorna o tempo de uptime do servidor e a quantidade de requisições realizadas.
- **Importação dos dados do CSV** de forma eficiente (desejável execução em menos de 1 minuto).
- **Armazenamento dos dados no PostgreSQL**.

---

## Arquitetura

A solução adota uma arquitetura inspirada em Clean Architecture e DDD, separando bem as camadas:

- **domain**: Entidades e regras de negócio puras.
- **usecase**: Casos de uso que orquestram a lógica de negócio.
- **infrastructure**: Implementações concretas de banco de dados, importação e validações.
- **interface**: Handlers HTTP (controllers) e roteamento.
- **cmd/server**: Ponto de entrada da aplicação.

Essa abordagem facilita manutenção, testes e evolução do código.

---

## Tecnologias Utilizadas

- **Linguagem**: Go (golang) 1.20+
- **Framework HTTP**: Gin
- **Banco de Dados**: PostgreSQL
- **ORM**: GORM
- **Testes**: `testing` nativo do Go
- **Docker**: Para containerização do banco de dados

---

## Instalação e Configuração

1. **Pré-requisitos**:
   - Go 1.20+
   - Docker e Docker Compose (opcional, caso queira rodar o banco em containers)
   - PostgreSQL (caso não utilize Docker, instale localmente)

2. **Clonar o repositório**:
   ```bash
   git clone https://github.com/gabrielmoura33/neoway-backend-test.git
   cd neoway-backend-test
   ```

3. **Variáveis de Ambiente**:
   Crie um arquivo `.env` na raiz do projeto com os seguintes valores (ajuste conforme necessidade):
   ```env
   DB_HOST=postgres
   DB_PORT=5432
   DB_USER=postgres
   DB_PASSWORD=postgres
   DB_NAME=neowaydb
   PORT=8080
   CSV_PATH=Base_Dados_Teste.csv
   ```

4. **Dependências**:
   ```bash
   go mod tidy
   ```

---

## Importação do CSV

O CSV deve seguir o formato:

```
DOCUMENTO,NOME/RAZAO_SOCIAL
30.607.887/0001-41,Alphadale Curtis
78.355.355/0001-90,Alphadale Hess
51.458.168/0001-86,Alphadale Kramer
43.472.780/0001-85,Alphadale Lawrence
```

- A importação ocorre automaticamente no start da aplicação, conforme configurado em `cmd/server/main.go`, utilizando a variável de ambiente `CSV_PATH` (por padrão, `Base_Dados_Teste.csv`).
- Caso deseje importar após o servidor estar rodando, você pode criar uma rota para importação manual (não implementada no exemplo final, mas descrita anteriormente).

---

## Execução

### Sem Docker

- Certifique-se que o PostgreSQL está rodando e as variáveis de ambiente apontam para o BD correto.
- Rode o servidor:
  ```bash
  go run cmd/server/main.go
  ```

A aplicação estará disponível em `http://localhost:8080`.

### Com Docker

1. Suba o container do Postgres:
   ```bash
   docker-compose up -d
   ```
   
2. Rode o servidor Go localmente (após configurar `.env`):
   ```bash
   go run cmd/server/main.go
   ```

---

## Rotas da API

- **GET `/status`**  
  Retorna informações de uptime do servidor e o número de requisições recebidas.

  **Exemplo de resposta**:
  ```json
  {
    "uptime": "5m30s",
    "request_count": 42
  }
  ```

- **POST `/clients`**  
  Cria um novo cliente.  
  **Exemplo de request**:
  ```json
  {
    "document": "30.607.887/0001-41",
    "name": "Alphadale Curtis",
    "is_blocked": false
  }
  ```

- **GET `/clients`**  
  Lista todos os clientes. Pode filtrar por nome usando `?name=`.  
  **Exemplo**: `/clients?name=Alphadale`

  **Exemplo de resposta**:
  ```json
  [
    {
      "id": 1,
      "document": "30.607.887/0001-41",
      "name": "Alphadale Curtis",
      "type": "PJ",
      "is_blocked": false
    }
  ]
  ```

- **GET `/clients/:document`**  
  Retorna o cliente pelo documento (CPF ou CNPJ).

- **GET `/exists?document=`**  
  Verifica se um documento já existe na base.  
  **Exemplo**: `/exists?document=30.607.887/0001-41`

  **Exemplo de resposta**:
  ```json
  {
    "exists": true
  }
  ```

---

## Validação de Documentos (CPF/CNPJ)

Por padrão, o código valida a estrutura do CPF/CNPJ. Caso o documento não seja válido, a criação falhará. Porém, a critério do negócio, a validação de dígitos verificadores pode ser removida ou relaxada, aceitando qualquer 11 dígitos como CPF e qualquer 14 dígitos como CNPJ.

---
## Testes e Cobertura

Para rodar os testes:

```bash
go test ./... -v
```

Para gerar o relatório de cobertura:

```bash
go test ./... -coverprofile=coverage.out
go tool cover -html=coverage.out
```

Isso abrirá o relatório de cobertura no navegador, mostrando quais linhas foram cobertas pelos testes.

---

## Docker e Docker Compose

O projeto inclui um `docker-compose.yml` e um `Dockerfile` para subir apenas o serviço do PostgreSQL:

- Para iniciar o PostgreSQL via Docker:
  ```bash
  docker-compose up -d
  ```
  
- O banco ficará disponível em `localhost:5432`, com usuário `postgres`, senha `postgres` e db `neowaydb`.

A aplicação Go pode ser executada localmente conectando-se ao container do PostgreSQL conforme as variáveis de ambiente definidas no `.env`.

---

## Melhorias Futuras

- Implementar testes de integração para rotas HTTP.
- Adicionar autenticação e autorização (caso necessário).
- Ampliar a lógica de negócio, por exemplo, controle de clientes bloqueados.
- Otimizar a importação, caso o CSV seja muito grande (ex: usando bulk insert).
- Criar uma rota para importar o CSV sob demanda.
- Adicionar uma containerização completa do projeto via docker-compose.