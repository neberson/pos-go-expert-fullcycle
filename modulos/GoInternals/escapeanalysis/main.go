package main

//func add(a, b int) int {
//	sum := a + b // sum é alocado na stack
//	return sum   // liberada quando a função termina
//}

//type User struct {
//	Name string
//}

// // A variável user é um ponteiro para um objeto User.
// // O ponteiro user é retornado da função NewUser
// // O objeto User precisa existindo após a função terminar,
// // pois o ponteiro é usado fora da função.
// // Portanto, o objeto User é alocado no heap.

//func NewUser(name string) *User {
//	user := &User{Name: name} // user é alocado no heap
//	return user               // retornando ponteiro para o heap
//}

// Armzenamento em Estrutura de Dados
// Se uma variável é armazenada em uma estrutura de dados que que sobrevive ao escopo da função,
// ela deve ser alocada no heap.
//func storeInMap() map[string]*int {
//	m := make(map[string]*int) // m é mapa que pode sobreviver ao escopo da função
//	i := 42                    // i é váriavel local
//	m["key"] = &i              // o ponteiro para i é armazenado no mapa
//	return m                   // retornando o mapa
//	// IMPORTANTE: se m não fosse retornado, ele não seria alocado para o heap
//}

// // // Se uma váriavel local é usada dentro de uma goroutine, ela deve ser alocada no heap.
// // // pois a goroutine pode continuar executando após a função retornar.

func startGoroutine() {
	go func() {
		i := 42 // i deve ser alocado no heap
		println(i)
	}()
}

func main() {
	//println(add(1, 2))

	//	user := NewUser("Alice")
	//	println(user.Name)
	// m := storeInMap()
	// println(*m["key"])

	startGoroutine() // espera a goroutine terminar
}
