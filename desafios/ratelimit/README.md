# Rate Limiter Go (Desafio FullCycle)

## Objetivo

Middleware de rate limiting para Go, configurável por IP ou token, com persistência em Redis e fácil extensão para outros storages.

## Como funciona

- Limita requisições por IP ou por token de acesso (header `API_KEY`).
- Configuração de limites e expiração via `.env` ou variáveis de ambiente.
- Persistência dos contadores e bloqueios no Redis.
- Middleware HTTP pronto para uso em servidores Go.
- Estratégia de storage desacoplada (fácil trocar Redis por outro mecanismo).

## Configuração

Crie um arquivo `.env` na raiz do projeto:

```
IP_LIMIT=10                # Limite de requisições por segundo por IP
IP_EXPIRE_SECONDS=1        # Janela de expiração do IP (segundos)
BLOCK_SECONDS=300          # Tempo de bloqueio após exceder limite (segundos)
REDIS_ADDR=localhost:6379  # Endereço do Redis
REDIS_PASSWORD=            # Senha do Redis (opcional)
REDIS_DB=0                 # DB do Redis
TOKEN_LIMITS=abc123:100    # Limites por token (token:limite, separados por vírgula)
TOKEN_EXPIRES=abc123:1     # Expiração por token (token:expira, separados por vírgula)
```

- O limite do token (se presente) sempre sobrepõe o do IP.
- Para múltiplos tokens: `TOKEN_LIMITS=abc123:100,xyz789:50` e `TOKEN_EXPIRES=abc123:1,xyz789:2`

## Uso

1. Suba o Redis (veja docker-compose abaixo).
2. Execute o servidor:

   ```sh
   go run ./cmd/ratelimit/main.go
   ```

3. Faça requisições para `http://localhost:8080/` com ou sem header `API_KEY`.

- **Se o limite for excedido:** resposta HTTP 429:

  ```
  you have reached the maximum number of requests or actions allowed within a certain time frame
  ```

- **Se o Redis estiver offline ou ocorrer erro interno:** resposta HTTP 500:
  ```
  internal server error: rate limiter unavailable
  ```

## Docker Compose (Execução completa)

O projeto já possui um `docker-compose.yml` pronto para subir tanto o Redis quanto a aplicação Go.

### Subindo tudo com Docker Compose

```sh
docker compose up --build
```

Isso irá:

- Buildar a aplicação Go usando o Dockerfile
- Subir o Redis e a aplicação já configurada com as variáveis do rate limiter

Acesse a aplicação em `http://localhost:8080/`.

---

#### Exemplo do arquivo `docker-compose.yml`:

```yaml
version: "3.8"
services:
  redis:
    image: redis:7
    ports:
      - "6379:6379"
    restart: always

  ratelimit:
    build: .
    depends_on:
      - redis
    environment:
      - REDIS_ADDR=redis:6379
      - REDIS_PASSWORD=
      - REDIS_DB=0
      - IP_LIMIT=10
      - IP_EXPIRE_SECONDS=1
      - BLOCK_SECONDS=300
      - TOKEN_LIMITS=abc123:100
      - TOKEN_EXPIRES=abc123:1
    ports:
      - "8080:8080"
```

## Testes

- Implemente testes automatizados em Go para garantir a robustez do rate limiter.
- Teste limites por IP, por token, bloqueios e desbloqueios.

## Extensão

- Para trocar o storage, implemente a interface `LimiterStorage` em `internal/storage`.

## Estrutura

- `cmd/ratelimit/main.go`: servidor e configuração
- `internal/limiter/`: lógica do rate limiter
- `internal/middleware/`: middleware HTTP
- `internal/storage/`: interface e implementação Redis

## Referências

- [Go Redis](https://github.com/redis/go-redis)
- [godotenv](https://github.com/joho/godotenv)

---

Desafio FullCycle - Expert Go
