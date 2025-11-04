package main

func add(a, b int) int {
	sum := a + b // sum é alocado na stack
	return sum   // liberada quando a função termina
}

type User struct {
	Name string
}

// // A variável user é um ponteiro para um objeto User.
// // O ponteiro user é retornado da função NewUser
// // O objeto User precisa existindo após a função terminar,
// // pois o ponteiro é usado fora da função.
// // Portanto, o objeto User é alocado no heap.

func NewUser(name string) *User {
	user := &User{Name: name} // user é alocado no heap
	return user               // retornando ponteiro para o heap
}

func main() {
	//println(add(1, 2))

	user := NewUser("Alice")
	println(user.Name)
}
