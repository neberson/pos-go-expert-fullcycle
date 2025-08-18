/*
Neste desafio você terá que usar o que aprendemos com Multithreading e APIs para buscar o resultado mais rápido entre duas APIs distintas.

As duas requisições serão feitas simultaneamente para as seguintes APIs:

https://brasilapi.com.br/api/cep/v1/ + cep

http://viacep.com.br/ws/" + cep + "/json/

Os requisitos para este desafio são:

- Acatar a API que entregar a resposta mais rápida e descartar a resposta mais lenta.

- O resultado da request deverá ser exibido no command line com os dados do endereço, bem como qual API a enviou.

- Limitar o tempo de resposta em 1 segundo. Caso contrário, o erro de timeout deve ser exibido.


Requisitos:

1 - Buscar CEP nas API's:
	Brasil API: https://brasilapi.com.br/api/cep/v1/01153000 + cep
	Via CEP: http://viacep.com.br/ws/" + cep + "/json/

2 - Acatar a API que entregar  a resposta mais rápida
3 - Descartar a reposta mais lenta
4 - Imprimir no command line os dados do endereço e qual API que deu o retorno.
5 - Limitar o tempo de resposta para 1 segundo.
6 - Apresentar timemout se exceder o tempo de resposta.
*/

package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"
)

type ViaCepAPIResponse struct {
	CEP         string `json:"cep"`
	Logradouro  string `json:"logradouro"`
	Complemento string `json:"complemento"`
	Unidade     string `json:"unidade"`
	Bairro      string `json:"bairro"`
	Localidade  string `json:"localidade"`
	UF          string `json:"uf"`
	Estado      string `json:"estado"`
	Regiao      string `json:"regiao"`
	Ibge        string `json:"ibge"`
	Gia         string `json:"gia"`
	DDD         string `json:"ddd"`
	Siafi       string `json:"siafi"`
}

type BrasilAPIResponse struct {
	CEP          string `json:"cep"`
	State        string `json:"state"`
	City         string `json:"city"`
	Neighborhood string `json:"neighborhood"`
	Street       string `json:"street"`
	Service      string `json:"service"`
}

func BuscaCepBrasilApi(cep string, canalBrasilAPI chan<- BrasilAPIResponse) {
	brasilApi := BrasilAPIResponse{}
	resp, err := http.Get("https://brasilapi.com.br/api/cep/v1/" + cep)
	if err != nil {
		log.Println(err.Error())
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		log.Printf("erro ao buscar cep no Brasil API: %s", resp.Status)
	}

	if err := json.NewDecoder(resp.Body).Decode(&brasilApi); err != nil {
		log.Printf("error decoding Brasil API response: %v", err)
	}

	canalBrasilAPI <- brasilApi
}

func BuscaCepViaCep(cep string, canalViaCep chan<- ViaCepAPIResponse) {
	viaCep := ViaCepAPIResponse{}
	resp, err := http.Get("http://viacep.com.br/ws/" + cep + "/json/")
	if err != nil {
		log.Println(err.Error())
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		log.Printf("erro ao buscar cep no Via Cep : %s", resp.Status)
	}

	if err := json.NewDecoder(resp.Body).Decode(&viaCep); err != nil {
		log.Printf("error decoding Via Cep response: %v", err)
	}

	canalViaCep <- viaCep
}

func main() {
	canalViaCep := make(chan ViaCepAPIResponse)
	canalBrasilAPI := make(chan BrasilAPIResponse)

	go BuscaCepBrasilApi("01153000", canalBrasilAPI)
	go BuscaCepViaCep("01153000", canalViaCep)

	select {
	case viaCep := <-canalViaCep:
		close(canalBrasilAPI)
		fmt.Printf("ViaCep API retornou: %+v\n", viaCep)
	case brasilApi := <-canalBrasilAPI:
		close(canalViaCep)
		fmt.Printf("Brasil API retornou: %+v\n", brasilApi)
	case <-time.After(1 * time.Second):
		fmt.Println("Timeout: Nenhuma resposta recebida em 1 segundo.")
	}
}
