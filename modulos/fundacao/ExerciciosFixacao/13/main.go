/*
Enunciado: Crie um slice de strings para representar uma lista de tarefas. Adicione as tarefas "Estudar Go", "Fazer exercícios" e "Preparar jantar". Em seguida, adicione mais uma tarefa: "Lavar louça". Imprima a lista de tarefas completa após todas as adições.
Objetivos de Aprendizado:
Criar e inicializar slices.
Adicionar elementos a um slice usando append().
Compreender a natureza dinâmica dos slices.
*/
package main

import "fmt"

func main() {
	tarefas := []string{"Estudar Go", "Fazer exercícios", "Preparar jantar"}

	fmt.Println("Lista de Tarefas Inicial:", tarefas)

	tarefas = append(tarefas, "Lavar louça")

	fmt.Println("Lista de Tarefas Atualizada: ", tarefas)
}
