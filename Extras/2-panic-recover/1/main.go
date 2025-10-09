package main

func myPanic1() {
	panic("panic1")
}

func myPanic2() {
	panic("panic2")
}
func main() {
	defer func() {
		if r := recover(); r != nil {
			if r == "panic1" {
				println("panic1 recovered", r.(string))
			}

			if r == "panic2" {
				println("panic2 recovered", r.(string))
			}
		}
	}()

	myPanic1()
}
