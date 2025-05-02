# API Go e framework Chi - Encurtador de Links

## 👨‍💻 Tecnologias

<div align="left">
  <img src="https://cdn.simpleicons.org/go/00ADD8" height="40" alt="go logo"  />
  <img width="12" />
  <img src="https://cdn.jsdelivr.net/gh/devicons/devicon/icons/redis/redis-original.svg" height="40" alt="redis logo"  />
  <img width="12" />
  <img src="https://cdn.jsdelivr.net/gh/devicons/devicon/icons/docker/docker-original.svg" height="40" alt="docker logo"  />
</div>

|Tecnologia | Descrição |
| --------  | --------- |
| Go        | Linguagem de programação estaticamente e fortemente tipada|
| Chi       | Framework Go que facilita a criação de servidores HTTP    |
| Redis     | Banco de dados em cache                                   |
| Docker    | Plataforma de software para implantar aplicativos em contâiners|

## 📝 Descrição do Projeto

O Encurtador de Links é uma API que recebe uma URL http padrão como *https://google.com* e a encurta em um código de poucos caracteres para otimizar o seu compartilhamento.

Toda a API foi desenvolvida em **Go** e com auxílio do framework **Chi** para acelar o desenvolvimento do projeto e adicionar recursos avançados como *middlewares*. Além disso, para a persistência dos dados estou usando o banco de dados em cache **Redis** e para ele rodar na aplicação usei sua imagem oficinal no **Docker**.

A opção de uso do **Redis** nessa aplicação se dá porque iremos salvar os dados nos pares: código e fullURL, isto é, a url encurtada e sua respectiva url completa. E também por se tratar de uma API simples com poucas funcionalidades.

## ⚡ Funcionalidades do projeto

- Encurtar a url passada e devolver a url encurtada (código)
- Devolver a url completa a partir do código

## ⚙ Endpoints

- Encurtar url: ``/api/shorten``
- Devolver url completa: ``/api/{code}``

Para a rota ``/api/shorten`` use o HTTP Method ``POST`` e para a rota ``/api/{code}`` use o HTTP Method ``GET``.

## 📂 Estrutura de pastas

```shell
├───cmd        # Pontos de entrada da aplicação
│   └───api    # API aplicada
└───internal   # Lógica de negócio
    ├───api    # Rotas e funcionalidades das rotas
    └───store  # Banco de dados da aplicação 
```

## Rodando localmente

Você precisa ter instalado [Go](https://go.dev/) em sua  máquina. Versão utilizada nesse projeto: ``1.22.4``.

1. Clone este repositório
2. Instale todas as dependências com os comandos: ``go mod tidy`` e também ``go get -u ./...``
3. Inicie o servidor com o comando ``go run .``
4. Acesse os endpoint fornecidos

**Obs.:** Foi utilizado o Docker para subir o banco de dados locamente. Se você tem o Docker instalado em sua máquina lembre-se de subir o banco com o comando ``docker compose up -d``.




