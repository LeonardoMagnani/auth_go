# Sistema de Autenticação JWT em Go com MongoDB e Gin-Gonic

##### Adaptado da versão: _https://github.com/akmamun/go-jwt_

Este projeto foi criado para estudo da linguagem. É um sistema de autenticação simples em Go que utiliza MongoDB para armazenamento de dados, JWT para autenticação e autorização, e oferece rotas para realizar operações como criar usuário, acessar, buscar todos usuários e buscar usuário por ID.

## Pré-requisitos

- Go (versão 1.16 ou superior)
- MongoDB (instância local ou acesso a uma instância remota)
- Pacotes Go necessários (instalados automaticamente pelo go mod, veja instruções abaixo)

## Instalação

#### 1. Clone o repositório
> git clone https://github.com/LeonardoMagnani/auth_go
#
#### 2. Instale as dependências usando `go mod`
> go mod tidy
#
#### 3. Configure as variáveis de ambiente
Crie um arquivo .env na raiz do projeto e configure as variáveis de ambiente necessárias. Um exemplo pode ser encontrado em .env.example.
>   MONGO_URI=mongodb://localhost:27017/db_name
    JWT_SECRET=secret_jwt
#
## Executando o Projeto
> go run main.go

O projeto será executado na porta :8080

## Rotas

#### Criar Usuário
> POST: /api/signup
{
    "username": "novousuario",
    "password": "senhadonovousuario"
}
#
#### Acessar
> POST: /api/login
{
    "username": "usuarioexistente",
    "password": "senhadousuario"
}
A resposta conterá um token JWT que deve ser incluído no cabeçalho Authorization para acessar rotas protegidas.
#
#### Buscar Todos Usuários
> GET: /api/users
> **HTTP Header:** `Authorization: Bearer TOKEN_JWT`
#
#### Buscar Usuário por ID
> GET: /api/users/{id}
> **HTTP Header:** `Authorization: Bearer TOKEN_JWT`
#
#
#### Se preferir, entre em contato. :D
[![LinkedIn](https://img.shields.io/badge/LinkedIn-Profile-blue?logo=linkedin)](https://www.linkedin.com/in/magnani-leonardo)

