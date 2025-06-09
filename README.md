Social Network Project
Este repositório contém o código-fonte de um projeto de rede social básica, implementado para demonstrar funcionalidades CRUD (Create, Read, Update, Delete) e exposição de endpoints. O backend é desenvolvido em Go, enquanto o frontend utiliza tecnologias web padrão como HTML, JavaScript e CSS.

Visão Geral do Projeto
O objetivo principal deste projeto é fornecer uma base para uma rede social simples, permitindo a gestão de entidades (como usuários e posts, mas pode ser expandido) através de operações CRUD. Ele é estruturado em duas partes principais:

API (Backend em Go): Responsável pela lógica de negócio, manipulação de dados e exposição de endpoints RESTful para interação.
WebApp (Frontend em HTML, JavaScript, CSS): Uma interface de usuário que consome a API para exibir e interagir com os dados da rede social.
Tecnologias Utilizadas
Backend:
Go (Golang)
Frontend:
HTML
JavaScript
CSS
Estrutura do Repositório
api/: Contém o código-source do backend em Go.
webapp/: Contém os arquivos do frontend (HTML, CSS, JavaScript).
.gitignore: Arquivo de configuração para o controle de versão do Git.
Como Rodar o Projeto (Configuração e Execução)
Para colocar o projeto em funcionamento na sua máquina local, siga os passos abaixo:

Pré-requisitos
Certifique-se que você tem o Go instalado na sua máquina. Você pode baixar e instalar a versão mais recente em golang.org/dl/.

Backend (API)
Navegue até o diretório da API:
Bash

cd social-network/api
Baixe as dependências (se houver):
Bash

go mod tidy
Execute o servidor da API:
Bash

go run main.go # ou o arquivo principal do seu backend, se for diferente
O servidor da API geralmente rodará em uma porta como 8080.
Frontend (WebApp)
Navegue até o diretório da WebApp:
Bash

cd social-network/webapp
Abra o arquivo index.html (ou o arquivo principal) em seu navegador.
Observação: Dependendo da sua implementação, você pode precisar de um servidor HTTP local simples para evitar problemas de CORS (Cross-Origin Resource Sharing) ao fazer requisições para a API. Uma opção comum é usar o http-server do Node.js:
Bash

npm install -g http-server
http-server .
E então acesse http://localhost:8080 (ou a porta indicada pelo http-server).
Interação
Com o backend e o frontend rodando, a interface da webapp deverá interagir com os endpoints expostos pela api para realizar as operações de rede social (criação de usuários, posts, etc.).
Endpoints da API
A API Go expõe endpoints RESTful para gerenciar as funcionalidades da rede social. Abaixo estão os endpoints comuns que você pode esperar em um projeto desse tipo, com base em operações CRUD para entidades como Usuários e Posts.

Base URL
A URL base para os endpoints da API será, tipicamente, algo como http://localhost:8080/api/v1 (ou a porta e prefixo que seu servidor Go estiver configurado para rodar).

Entidade: Usuários (/users)
POST /users

Descrição: Cria um novo usuário.
Corpo da Requisição (Exemplo):
JSON

{
    "username": "novo_usuario",
    "email": "email@example.com",
    "password": "senha_segura"
}
Resposta: Retorna os detalhes do usuário criado (sem a senha).
GET /users

Descrição: Lista todos os usuários.
Resposta: Um array de objetos de usuário.
GET /users/{id}

Descrição: Busca um usuário específico pelo ID.
Resposta: Os detalhes do usuário correspondente.
PUT /users/{id}

Descrição: Atualiza os dados de um usuário existente.
Corpo da Requisição (Exemplo):
JSON

{
    "email": "novo_email@example.com"
}
Resposta: Os detalhes atualizados do usuário.
DELETE /users/{id}

Descrição: Exclui um usuário específico.
Resposta: Confirmação da exclusão.
Entidade: Posts (/posts)
POST /posts

Descrição: Cria um novo post para um usuário.
Corpo da Requisição (Exemplo):
JSON

{
    "user_id": "ID_do_usuario",
    "content": "Conteúdo do meu novo post!"
}
Resposta: Os detalhes do post criado.
GET /posts

Descrição: Lista todos os posts (opcionalmente pode ter filtros por usuário, etc.).
Resposta: Um array de objetos de post.
GET /posts/{id}

Descrição: Busca um post específico pelo ID.
Resposta: Os detalhes do post correspondente.
PUT /posts/{id}

Descrição: Atualiza o conteúdo de um post.
Corpo da Requisição (Exemplo):
JSON

{
    "content": "Conteúdo do post atualizado."
}
Resposta: Os detalhes atualizados do post.
DELETE /posts/{id}

Descrição: Exclui um post específico.
Resposta: Confirmação da exclusão.
Contribuição
Contribuições são bem-vindas! Se você tiver sugestões ou quiser melhorar o projeto, sinta-se à vontade para abrir uma issue ou enviar um pull request.

Autor: Gabriel Barboza

