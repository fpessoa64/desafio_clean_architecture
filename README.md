# Order Clean Architecture (Go + SQLite)

## Visão Geral

Aplicação Go demonstrando Clean Architecture, SOLID, e múltiplas interfaces (REST, gRPC, GraphQL) usando SQLite.

- **REST:** GET/POST `/order`
- **gRPC:** ListOrders
- **GraphQL:** query listOrders
- **Banco:** SQLite (arquivo local)

## Estrutura

```
order-clean-arch/
├─ cmd/server/main.go
├─ internal/
│  ├─ entities/
│  ├─ usecase/
│  ├─ repository/
│  │   └─ sqlite/
│  └─ delivery/
│      ├─ rest/
│      ├─ grpc/
│      └─ graphql/
├─ migrations/
├─ proto/
├─ api.http
├─ go.mod
├─ Dockerfile
├─ Dockerfile.dev
├─ docker-compose.yaml
├─ docker-compose.dev.yaml
└─ README.md
```

## Rodando em Produção

```sh
docker compose up --build
```
- A aplicação será exposta em:
  - REST/GraphQL: http://localhost:8080
  - gRPC: localhost:50051
- O banco SQLite será criado como `orders.db` no container.

## Rodando em Desenvolvimento (VS Code)

```sh
docker compose -f docker-compose.dev.yaml up --build
```
- Código é montado via volume.
- Alterações exigem reinício do container.

## Variáveis de Ambiente
- `DATABASE_URL=file:orders.db?cache=shared&_foreign_keys=on`



## Documentação Swagger (REST)

O projeto já inclui geração automática de documentação Swagger para a API REST.

**Como gerar/atualizar a documentação:**

```sh
swag init -g cmd/server/main.go --parseDependency --parseInternal
```

**Como acessar a interface Swagger UI:**

- Inicie a aplicação normalmente
- Acesse: [http://localhost:8080/swagger/index.html](http://localhost:8080/swagger/index.html)

Você pode testar e visualizar todos os endpoints REST diretamente pela interface web.

---

## Endpoints REST (v1)

```
### Criar order
POST http://localhost:8080/v1/order
Content-Type: application/json

{
  "customer_name": "Fernando",
  "amount": 123.45,
  "status": "pending"
}

### Listar orders
GET http://localhost:8080/v1/order
```

Você pode testar facilmente os endpoints REST usando o arquivo `tests/orders.http` no VS Code (com a extensão REST Client).

## GraphQL

```
POST http://localhost:8080/graphql
Content-Type: application/json

{"query":"{ listOrders { id customerName amount status createdAt } }"}
```


## gRPC

### Testando com Evans (CLI gRPC)

Você pode testar o serviço gRPC facilmente usando o Evans:

1. Certifique-se de que o servidor está rodando e a porta gRPC está exposta (padrão: 50051).
2. No terminal, execute:

```sh
evans -r repl --proto proto/order.proto --host localhost --port 50051
```

3. No prompt do Evans, selecione o package e o serviço:

```
package order
service OrderService
```

4. Para listar pedidos:

```
call ListOrders
```

5. Para criar um pedido:

```
call CreateOrder
```
Preencha os campos conforme solicitado, por exemplo:
```
name (TYPE_STRING) => Fernando
amount (TYPE_DOUBLE) => 123.45
status (TYPE_STRING) => pending
```

Veja o arquivo `proto/order.proto` para detalhes dos campos.

## Migrações
- Migrações SQL em `migrations/`
- Aplicadas automaticamente ao iniciar a aplicação.

## Dependências principais
- Go 1.21+
- SQLite
- chi, gqlgen, gRPC, migrate

## Testes
- (Opcional) Testes unitários podem ser adicionados em `internal/usecase/`

---

Qualquer dúvida, consulte o código ou abra uma issue!
