package main

import (
	"fmt"
	"github.com/FadhliAbrar/belajar-go-abstractfactory/abstract"
)

func main() {
	factory, err := abstract.New("logitech")
	if err != nil {
		panic(err)
	}
	f := factory.CreateMouse()
	f.SetCableLen(200)

	fmt.Println(f.GetCableLen())

}
