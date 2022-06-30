package abstract

import "github.com/FadhliAbrar/belajar-go-abstractfactory/concrete"

type Zowie struct{}

func (z *Zowie) CreateMouse() concrete.ConcreteMouse {
	return &concrete.Mouse{
		MouseSensor: "3360",
		MouseWeight: 82,
	}
}

func (z *Zowie) CreateHeadset() concrete.ConcreteHeadset {
	return &concrete.Headset{
		FrequencyInDb: 22,
		Berat:         120,
	}
}
