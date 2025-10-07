# Desafio Go Expert - Clean Architecture

## Proposta

Olá devs!

Agora é a hora de botar a mão na massa. Para este desafio, você precisará criar a listagem das orders.

Esta listagem precisa ser feita com:

- **Endpoint REST** (GET /order)
- **Service ListOrders com GRPC**
- **Query ListOrders GraphQL**

## Setup inicial

Clone o repositório com o comando abaixo:

```bash
git clone https://github.com/neberson/pos-go-expert-fullcycle.git
```

Entre no diretório do projeto:

```bash
cd pos-go-expert-fullcycle/desafios/CleanArchitecture
```

Execute o docker para subir a imagem do MySQL e do RabbitMQ:

```bash
docker-compose up -d
```

Após subir as imagens, vamos verificar se o banco de dados `orders` está criado:

```bash
docker exec -it mysql bash -c "mysql -u root -proot -D orders"
```

Caso não esteja criado, será exibido no terminal a seguinte mensagem: `Unknown database 'orders'`, portanto, execute o seguinte comando:

```bash
docker exec -it mysql bash -c "mysql -u root -proot"
CREATE DATABASE orders;
```

Com o banco de dados criado, vamos verificar se existe a tabela orders:

```bash
docker exec -it mysql bash -c "mysql -u root -proot -D orders"
```

Rode o comando abaixo para verificar se existe:

```sql
select * from orders.orders;
```

Caso não, execute o comando abaixo:

```sql
CREATE TABLE orders (id varchar(255) NOT NULL, price float NOT NULL, tax float NOT NULL, final_price float NOT NULL, PRIMARY KEY (id));
```

## Como rodar

No diretório do projeto, execute o comando abaixo para baixar as dependências:

```bash
go mod tidy
```

Agora iremos executar nossa aplicação:

```bash
cd cmd/ordersystem && go run main.go wire_gen.go
```

Caso ocorra tudo bem, os serviços estarão rodando nos endereços:

**Rest** em http://localhost:8000:

- Use os arquivos na pasta `/api` para interagir;
- Será necessário instalar a extensão: https://marketplace.visualstudio.com/items?itemName=humao.rest-client.

**GraphQL** em http://localhost:8080:

Use o template abaixo:

```graphql
mutation createOrder {
  createOrder(input: { id: "change-id", Price: 10.0, Tax: 0.5 }) {
    id
    Price
    Tax
    FinalPrice
  }
}

query orders {
  listOrders {
    id
    Price
    Tax
    FinalPrice
  }
}
```

**gRPC** na porta 50051:

- Será necessário uma aplicação externa para interagir, sugiro a ferramenta `evans`.
