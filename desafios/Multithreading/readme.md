# Consulta de CEP com Multithreading

## Sobre o Projeto

Este projeto foi desenvolvido como parte do desafio do curso de pós-graduação GO EXPERT. O objetivo é implementar uma aplicação que consulta CEPs em duas APIs diferentes simultaneamente usando multithreading, retornando os dados da API que responder mais rapidamente.

## Desafio Proposto

O desafio consiste em:

1. Realizar requisições simultâneas para duas APIs de consulta de CEP:

   - Brasil API: `https://brasilapi.com.br/api/cep/v1/{cep}`
   - Via CEP: `http://viacep.com.br/ws/{cep}/json/`

2. Acatar a resposta da API que entregar o resultado mais rápido e descartar a mais lenta.

3. Exibir no terminal os dados do endereço retornado, indicando qual API forneceu a resposta.

4. Limitar o tempo de resposta em 1 segundo, exibindo um erro de timeout caso exceda esse limite.

## Implementação

A solução utiliza:

- Goroutines para fazer requisições HTTP paralelas
- Canais para comunicação entre goroutines
- Select statement para tratar a resposta mais rápida
- Timeout para limitar o tempo de espera
