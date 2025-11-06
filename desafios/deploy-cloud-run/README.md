# Weather by CEP (Go + Cloud Run)

Sistema em Go que recebe um CEP (8 dígitos), identifica a cidade via ViaCEP e retorna o clima atual (temperaturas em Celsius, Fahrenheit e Kelvin) consultando a WeatherAPI. O serviço está publicado no Google Cloud Run e também pode ser executado localmente via Docker Compose.

- Demo (Cloud Run, exemplo de sucesso):  
  https://weather-api-on-d5bzjlhlkq-uc.a.run.app/weather/75915000

## Objetivo

- Receber CEP válido de 8 dígitos.
- Buscar a localização (cidade) via ViaCEP.
- Consultar temperaturas atuais na WeatherAPI.
- Retornar as temperaturas em:
  - Celsius (C)
  - Fahrenheit (F = C \* 1.8 + 32)
  - Kelvin (K = C + 273)

## Requisitos de resposta

- Sucesso
  - HTTP 200
  - Body: { "temp_C": 28.5, "temp_F": 83.3, "temp_K": 301.5 }
- Falha (CEP inválido — formato diferente de 8 dígitos)
  - HTTP 422
  - Body: { "message": "invalid zipcode" }
- Falha (CEP não encontrado)
  - HTTP 404
  - Body: { "message": "can not find zipcode" }

## Endpoint

- GET /weather/{cep}
  - Parâmetros:
    - cep: string com 8 dígitos (somente números)
  - Exemplo:
    - GET /weather/75915000

## Execução local

Você pode rodar a aplicação localmente com Docker Compose. A aplicação expõe a porta 8080 por padrão.

1. Crie um arquivo .env na raiz do projeto com sua chave da WeatherAPI (https://www.weatherapi.com/):

```
WEATHER_API_KEY=coloque_sua_chave_aqui
```

2. Suba a aplicação:

- Windows PowerShell:

  - docker compose up --build

- Alternativa sem .env:
  - PowerShell
    - $env:WEATHER_API_KEY="coloque_sua_chave_aqui"
    - docker compose up --build

O serviço ficará disponível em:

- http://localhost:8080/weather/75915000

### Testes rápidos (HTTP)

Use a coleção em api/weather.http (VS Code + extensão “REST Client”) ou curl:

- Sucesso:
  - curl http://localhost:8080/weather/75915000
- CEP inválido:
  - curl http://localhost:8080/weather/75915-000
  - curl http://localhost:8080/weather/75915
  - curl http://localhost:8080/weather/7591500000
- CEP não encontrado:
  - curl http://localhost:8080/weather/00000000

## Deploy (Cloud Run)

O serviço está publicado e acessível publicamente:

- Base: https://weather-api-on-d5bzjlhlkq-uc.a.run.app
- Exemplo:
  - https://weather-api-on-d5bzjlhlkq-uc.a.run.app/weather/75915000

Notas:

- O serviço foi containerizado (Dockerfile multi-stage) e implantado no Cloud Run.
- A chave da WeatherAPI é lida via variável de ambiente WEATHER_API_KEY (configure-a como variável de ambiente no serviço do Cloud Run).

## Ambiente e configuração

- Variáveis de ambiente:
  - WEATHER_API_KEY (obrigatória): chave da WeatherAPI.

## Desenvolvimento

- Go (binário principal):

  - cmd/server/main.go

- Build local (sem Docker), se desejar:

  - go mod tidy
  - go run ./cmd/server/main.go

## Docker

- Dockerfile:

  - Multi-stage build (golang para build, alpine para runtime).
  - Binário estático para execução em container leve.

- Docker Compose:
  - Sobe a aplicação expondo 8080:8080.
  - Encaminha WEATHER_API_KEY do seu ambiente/.env para o container.

## Troubleshooting

- 422 invalid zipcode: confirme que o CEP tem exatamente 8 dígitos numéricos.
- 404 can not find zipcode: CEP não encontrado na ViaCEP.
- Container não inicia: verifique se a variável WEATHER_API_KEY está definida corretamente.

## Referências

- ViaCEP: https://viacep.com.br/
- WeatherAPI: https://www.weatherapi.com/
