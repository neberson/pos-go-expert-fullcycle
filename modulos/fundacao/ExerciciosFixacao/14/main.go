/*
Enunciado: Crie um map onde as chaves são palavras (string) e os valores são seus sinônimos (string). Adicione os seguintes pares:
"casa": "lar"
"feliz": "alegre"
"rápido": "veloz"
Em seguida, faça o seguinte:
Imprima o sinônimo da palavra "feliz".
Tente acessar o sinônimo de uma palavra que não existe no map (ex: "triste") e imprima uma mensagem indicando se a palavra foi encontrada ou não.
Atualize o sinônimo de "casa" para "residência".
Delete a palavra "rápido" do map.
Imprima o map completo após todas as operações.
Objetivos de Aprendizado:
Criar e inicializar maps.
Adicionar, acessar, atualizar e deletar elementos em um map.
Verificar a existência de uma chave em um map.
*/
package main

import "fmt"

func main() {
	sinonimos := map[string]string{"casa": "lar", "feliz": "alegre", "rápido": "veloz"}

	fmt.Println("Dicionário de Sinônimos Inicial:", sinonimos)

	fmt.Printf(`O sinônimo da palavra "feliz" é: %s`, sinonimos["feliz"])
	fmt.Println("")

	sinonimoTriste, encontrado := sinonimos["triste"]
	if encontrado {
		fmt.Println("Sinônimo de 'triste':", sinonimoTriste)
	} else {
		fmt.Println("Palavra 'triste' não encontrada no dicionário.")
	}

	sinonimos["casa"] = "residência"
	fmt.Println("Dicionário após atualização de 'casa':", sinonimos)

	delete(sinonimos, "rápido")
	fmt.Println("Dicionário após remoção de 'rápido':", sinonimos)
}
