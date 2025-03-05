# API Go com Integração com Stripe

Este projeto é uma API baseada em Go que se integra ao Stripe para gerenciar produtos, preços e sessões de checkout. Ela pode ser executada localmente usando `go run main.go` ou via Docker Compose para uma configuração em contêiner.

---

## Funcionalidades
- Integração com o Stripe para gerenciamento de produtos e checkout.
- API RESTful usando Gin para tratamento de requisições.
- Configuração com Docker Compose.

---

## Pré-requisitos

### Ferramentas Necessárias
- [Go](https://golang.org/doc/install) (versão 1.19 ou superior)
- [Docker](https://www.docker.com/get-started)
- [Docker Compose](https://docs.docker.com/compose/install/)

---

## Configuração

### 1. Clone o repositório
```bash
git clone https://github.com/your-repo/back-end-mirumuh
```

```bash
cd back-end-mirumuh
```

### 2. Instalar dependências

```bash
go mod tidy
```

### 3. Rode o projeto

```bash
go run main.go
```
- Ou também pode ser utilizado o Docker via Docker Compose

```bash
docker compose up
```