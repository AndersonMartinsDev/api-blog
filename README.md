Markdown
# Projeto: API de Blog (Go)

## Descrição
Esta API RESTful em Go foi desenvolvida para fornecer os dados necessários para um blog, como publicações e seus respectivos autores. A API utiliza o banco de dados PostgreSQL para persistir as informações.

## Pré-requisitos
* **Go:** Certifique-se de ter o Go instalado e configurado em seu sistema.
* **Docker:** (Opcional) Para facilitar a execução e a configuração do ambiente.
* **PostgreSQL:** Um banco de dados PostgreSQL com as tabelas necessárias criadas.

## Instalação
1. **Clone o repositório:**
   ```bash
    git clone https://github.com/AndersonMartinsDev/api-blog
   ```
2. **Instale as dependências:**
    ```bash
     cd api-blog
     go mod tidy
    ```
2. **Configure o banco de dados:**

    Crie um arquivo .env na raiz do projeto com as seguintes variáveis:

    ```env
     APPLICATION_NAME | string
     DATABASE_NAME | string
     DATABASE_HOST | string
     DATABASE_PORT | string
     DATABASE_USER | string
     DATABASE_PASSWORD | string
    ```

## Execução

1. **Construir o Docker image (opcional)**
    ```bash
     docker build -t sua-imagem .
    ```
2. **Iniciar o container**
    ```bash
     docker run -p 8080:8080 sua-imagem
    ```
3. **Executar diretamente (sem Docker)**
    ```bash
     go run main.go
    ```

## Endpoints API 


- Criar uma publicação:
    - Método: POST
    - URL: /publication/insert
    - Corpo da requisição: JSON com os campos título, corpo, autor

- Listar publicações:
    - Método: GET 
    - URL: /publication
    - Parâmetros de consulta:
         - autor: Filtra por autor
         - order: Ordena por data (asc ou desc)

- Detalhes de uma publicação:
    - Método: GET
    - URL: /publication/{title}/{autor}

- Atualizar uma publicação:
    - Método: PUT
    - URL: /publication
    - Corpo da requisição: JSON com os campos título, corpo, autor

- Deletar uma publicação:
    - Método: DELETE
    - URL: /publication/{title}/{autor}

## Estrutura do Projeto

```
api-blog/
├── src/
│   ├── application/
|   |   ├── dtos
|   |   └── mappers
│   ├── domain/
|   |   ├── entities
|   |   ├── repositories
|   |   |   └── impl
|   |   └── services
|   |   |   └── impl
│   ├── external/
|   |   ├── controllers
|   |   └── routes
│   └── infrastructure/
|   |   ├── commons/
|   |   ├── configuration
|   |   ├── database
|   |   ├── interceptor
|   |   ├── request
|   |   └── response
├── test/
│   ├── domain/
|   |   ├── services
|   │   └── mocks
├── .env
├── go.mod
├── go.sum
├── LICENSE
├── main.go
└── README.md

```