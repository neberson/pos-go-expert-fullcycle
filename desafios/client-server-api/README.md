# Go Expert: Desafio Client Server API

Sistema para armazenamento do histórico de cotação do dólar e fornecimento da cotação do dólar em tempo real.

## Componentes do Sistema

O sistema é composto por dois componentes principais: Cliente e Servidor.

### Servidor

O Servidor é responsável por:
- Monitorar as variações na cotação do dólar através da integração com uma API externa (`https://economia.awesomeapi.com.br/json/last/USD-BRL`).
- Persistir estas alterações de forma histórica em um banco de dados SQLite.
- Fornecer uma API (`/cotacao` na porta `8080`) para retornar a cotação atual do dólar.
- Gerenciar timeouts para a chamada da API externa (200ms) e para a persistência no banco de dados (10ms).

### Cliente

O Cliente é responsável por:
- Solicitar ao Servidor a cotação atual do dólar.
- Receber apenas o valor "bid" da cotação.
- Manter um arquivo local `./cotacao.txt` com a cotação atual no formato `Dólar: {valor}`.
- Gerenciar timeout para a requisição ao servidor (300ms).

## Requisitos

Este repositório requer que o Go esteja instalado previamente para a execução ou compilação do sistema.

A versão correta do Go que precisa estar instalada pode ser encontrada em `./go.mod`.

## Como Executar o Sistema

Devido às dependências entre os componentes, é necessário iniciar primeiro o Servidor e, apenas após o servidor estar iniciado, iniciar o Cliente.

No diretório raiz deste repositório, execute os seguintes comandos:

```bash
go run cmd/client/main.go
go run cmd/server/main.go
```

### Sugestão para Verificação dos Requisitos

Todos os parâmetros configuráveis são encontrados nos arquivos de configuração (se aplicável, no seu caso, os timeouts são no código). As lógicas de timeout podem ser facilmente testadas alterando os valores configurados para os limites de leitura e escrita. Uma sugestão de valor para testes é "1ns" para forçar o timeout.

Os logs de erro relacionados aos timeouts serão exibidos no console de cada aplicação (servidor e cliente) caso ocorram.

