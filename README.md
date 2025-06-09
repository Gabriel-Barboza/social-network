# Social Network Project

Este repositório contém o código-fonte de um projeto de rede social básica, implementado para demonstrar funcionalidades CRUD (Create, Read, Update, Delete) e exposição de endpoints. O backend é desenvolvido em Go, enquanto o frontend utiliza tecnologias web padrão como HTML, JavaScript e CSS.

## Visão Geral do Projeto

O objetivo principal deste projeto é fornecer uma base para uma rede social simples, permitindo a gestão de entidades (como usuários e posts, mas pode ser expandido) através de operações CRUD. Ele é estruturado em duas partes principais:

* **API (Backend em Go):** Responsável pela lógica de negócio, manipulação de dados e exposição de endpoints RESTful para interação.
* **WebApp (Frontend em HTML, JavaScript, CSS):** Uma interface de usuário que consome a API para exibir e interagir com os dados da rede social.

## Tecnologias Utilizadas

* **Backend:**
    * Go (Golang)
* **Frontend:**
    * HTML
    * JavaScript
    * CSS

## Estrutura do Repositório

* `api/`: Contém o código-source do backend em Go.
* `webapp/`: Contém os arquivos do frontend (HTML, CSS, JavaScript).
* `.gitignore`: Arquivo de configuração para o controle de versão do Git.

## Como Rodar o Projeto (Configuração e Execução)

Para colocar o projeto em funcionamento na sua máquina local, siga os passos abaixo:

### Pré-requisitos

Certifique-se que você tem o Go instalado na sua máquina. Você pode baixar e instalar a versão mais recente em [golang.org/dl/](https://golang.org/dl/).

### Backend (API)

1.  **Navegue até o diretório da API:**
    ```bash
    cd social-network/api
    ```
2.  **Baixe as dependências (se houver):**
    ```bash
    go mod tidy
    ```
3.  **Execute o servidor da API:**
    ```bash
    go run main.go # ou o arquivo principal do seu backend, se for diferente
    ```
    O servidor da API geralmente rodará em uma porta como `8080`.

### Frontend (WebApp)

1.  **Navegue até o diretório da WebApp:**
    ```bash
    cd social-network/webapp
    ```
2.  **Abra o arquivo `index.html` (ou o arquivo principal) em seu navegador.**
    * **Observação:** Dependendo da sua implementação, você pode precisar de um servidor HTTP local simples para evitar problemas de CORS (Cross-Origin Resource Sharing) ao fazer requisições para a API. Uma opção comum é usar o `http-server` do Node.js:
        ```bash
        npm install -g http-server
        http-server .
        ```
        E então acesse `http://localhost:8080` (ou a porta indicada pelo `http-server`).

### Interação

* Com o backend e o frontend rodando, a interface da `webapp` deverá interagir com os endpoints expostos pela `api` para realizar as operações de rede social (criação de usuários, posts, etc.).


**Autor:** Gabriel Barboza

---
