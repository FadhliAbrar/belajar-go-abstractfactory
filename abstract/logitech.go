package abstract

import (
	"github.com/FadhliAbrar/belajar-go-abstractfactory/concrete"
)

type Logitech struct{}

type LogitechHeadset struct {
	Hs concrete.Headset
}

func (l *Logitech) CreateMouse() concrete.ConcreteMouse {
	return &concrete.Mouse{
		MouseSensor: "40k CAM",
		MouseWeight: 100,
	}
}

func (l *Logitech) CreateHeadset() concrete.ConcreteHeadset {
	return &concrete.Headset{
		Berat:         100,
		FrequencyInDb: 20,
	}
}
