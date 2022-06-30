package concrete

import "fmt"

type Headset struct {
	Berat         int
	FrequencyInDb int
}

func (h *Headset) GetBerat() {
	fmt.Println(h.Berat)
}

func (h *Headset) GetSensitivitasSensor() {
	fmt.Println(h.FrequencyInDb)
}
