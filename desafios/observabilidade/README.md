# Sistema de Temperatura por CEP — Observabilidade com OTEL e Zipkin

Este projeto implementa dois serviços em Go:

- **Serviço A**: Recebe um CEP via POST, valida o input e encaminha para o Serviço B.
- **Serviço B**: Recebe o CEP, consulta a cidade via ViaCEP, obtém a temperatura atual via WeatherAPI e retorna as temperaturas em Celsius, Fahrenheit, Kelvin, juntamente com o nome da cidade.

A solução implementa tracing distribuído com OpenTelemetry (OTEL) e Zipkin, permitindo rastrear as requisições ponta a ponta entre os serviços.

---

## Arquitetura

```
[Usuário] → [Serviço A] → [Serviço B] → [ViaCEP/WeatherAPI]
                        ↘
                        [OTEL Collector + Zipkin]
```

- **Serviço A**: expõe `/weather` (POST), valida o CEP e repassa para o Serviço B.
- **Serviço B**: expõe `/weather/{cep}` (GET), orquestra as chamadas externas e retorna o resultado.
- **OTEL Collector + Zipkin**: coletam e exibem os traces das requisições.

---

## Endpoints

### Serviço A

- **POST /weather**
  - Body: `{ "cep": "29902555" }`
  - Valida se o CEP é uma string de 8 dígitos.
  - Encaminha para o Serviço B.
  - Respostas:
    - **422**: `{ "message": "invalid zipcode" }` (CEP inválido)
    - **200**: `{ "city": "São Paulo", "temp_C": 28.5, "temp_F": 83.3, "temp_K": 301.5 }` (sucesso)
    - **404**: `{ "message": "can not find zipcode" }` (CEP não encontrado)

### Serviço B

- **GET /weather/{cep}**
  - Parâmetro: `cep` (string de 8 dígitos)
  - Busca cidade via ViaCEP e temperatura via WeatherAPI.
  - Respostas:
    - **422**: `{ "message": "invalid zipcode" }`
    - **200**: `{ "city": "São Paulo", "temp_C": 28.5, "temp_F": 83.3, "temp_K": 301.5 }`
    - **404**: `{ "message": "can not find zipcode" }`

---

## Observabilidade

- **Tracing distribuído**: Implementado com OpenTelemetry (OTEL) e Zipkin.
- **Spans**: Medem o tempo de resposta das chamadas ao ViaCEP e WeatherAPI.
- **Collector**: O OTEL Collector recebe os traces dos dois serviços e exporta para o Zipkin.

Acesse o Zipkin em [http://localhost:9411](http://localhost:9411) para visualizar os traces.

---

## Como rodar o projeto (ambiente de desenvolvimento)

1. **Pré-requisitos**:

   - Docker e Docker Compose instalados.
   - Chave da WeatherAPI (cadastre-se em https://www.weatherapi.com/).

2. **Configuração**:

   - Crie um arquivo `.env` na raiz do projeto:
     ```
     WEATHER_API_KEY=sua_chave_weatherapi
     ```
   - (Opcional) Defina a variável de ambiente manualmente:
     - PowerShell: `$env:WEATHER_API_KEY="sua_chave_weatherapi"`

3. **Suba os serviços**:

   ```sh
   docker compose up --build
   ```

4. **Testes rápidos**:

   - **Serviço A (POST):**

     ```http
     POST http://localhost:8080/weather
     Content-Type: application/json

     {
       "cep": "75915000"
     }
     ```

   - **Serviço B (GET):**

     ```http
     GET http://localhost:8181/weather/75915000
     ```

   - Exemplos de respostas e testes estão em [api/weather-a.http](api/weather-a.http) e [api/weather-b.http](api/weather-b.http).

5. **Acompanhe os traces**:
   - Acesse [http://localhost:9411](http://localhost:9411) para visualizar os traces no Zipkin.

---

## Respostas esperadas

- **Sucesso**:
  ```json
  {
    "city": "São Paulo",
    "temp_C": 28.5,
    "temp_F": 83.3,
    "temp_K": 301.5
  }
  ```
- **CEP inválido**:
  ```json
  { "message": "invalid zipcode" }
  ```
- **CEP não encontrado**:
  ```json
  { "message": "can not find zipcode" }
  ```

---

## Observações

- O tracing distribuído está implementado entre os serviços, com spans para as chamadas externas.
- O projeto utiliza multi-stage Dockerfiles para imagens leves.
- O OTEL Collector e Zipkin são orquestrados via Docker Compose.

---

## Referências

- [ViaCEP](https://viacep.com.br/)
- [WeatherAPI](https://www.weatherapi.com/)
- [OpenTelemetry](https://opentelemetry.io/)
- [Zipkin](https://zipkin.io/)
