package main

import (
	"github.com/google/uuid"
	"github.com/neberson/pos-go-expert-fullcycle/tree/main/modulos/Packaging/6-UsandoWorkSpaces/math"
)

func main() {
	m := math.NewMath(1, 2)
	println(m.Add())
	println(uuid.New().String())
}
