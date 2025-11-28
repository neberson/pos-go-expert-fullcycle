# StressTest CLI - Go

Ferramenta CLI para realizar testes de carga (stress test) em serviços web HTTP.

## Funcionalidades

- Envio de múltiplos requests HTTP concorrentes para uma URL informada
- Parâmetros configuráveis via linha de comando (Cobra CLI)
- Relatório detalhado ao final do teste:
  - Tempo total de execução
  - Total de requests realizados
  - Quantidade de respostas HTTP 200
  - Distribuição de outros códigos de status HTTP
- Pronto para execução via Docker ou Docker Compose

## Como usar

### 1. Build local

```sh
go build -o stresstest ./cmd/stresstest
./stresstest --url=http://google.com --requests=1000 --concurrency=10
```

### 2. Docker

Build da imagem:

```sh
docker build -t stresstest .
```

Execução:

```sh
docker run --rm stresstest --url=http://google.com --requests=1000 --concurrency=10
```

### 3. Docker Compose

Edite o arquivo `docker-compose.yml` para ajustar os parâmetros desejados:

```yaml
command: ["--url=http://google.com", "--requests=1000", "--concurrency=10"]
```

Execute:

```sh
docker compose up --build
```

## Parâmetros CLI

- `--url` (obrigatório): URL do serviço a ser testado
- `--requests`: Número total de requests (padrão: 1)
- `--concurrency`: Número de chamadas simultâneas (padrão: 1)

## Estrutura do Projeto

- `cmd/stresstest/` - Entrada principal da aplicação (Cobra CLI)
- `internal/stresstest/` - Domínio e lógica de stress test
- `pkg/report/` - Geração e exibição do relatório

## Exemplo de Saída

```
Iniciando teste de carga em http://google.com com 1000 requests e 10 concorrentes...

--- Relatório ---
Tempo total: 3.2s
Total de requests: 1000
HTTP 200: 995
Distribuição de status:
HTTP 404: 3
Erros de request: 2
```

## Requisitos

- Go 1.24+
- Docker (opcional)
- Docker Compose (opcional)

## Licença

MIT

---

Desafio FullCycle - Expert Go
