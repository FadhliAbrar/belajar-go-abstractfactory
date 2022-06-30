package abstract

import (
	"errors"
	"github.com/FadhliAbrar/belajar-go-abstractfactory/concrete"
)

type GamingFactory interface {
	CreateMouse() concrete.ConcreteMouse
	CreateHeadset() concrete.ConcreteHeadset
}

func New(merk string) (GamingFactory, error) {
	if merk == "logitech" {
		return &Logitech{}, nil
	}
	if merk == "zowie" {
		return &Zowie{}, nil
	}
	return nil, errors.New("Merk tidak ditemukan!")
}
