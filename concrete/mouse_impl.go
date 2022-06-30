package concrete

type Mouse struct {
	MouseSensor string
	MouseWeight int
	CableLen    int
}

func (m *Mouse) GetSensor() string {
	return m.MouseSensor
}

func (m *Mouse) GetBerat() int {
	return m.MouseWeight
}

func (m *Mouse) GetCableLen() int {
	return m.CableLen
}
func (m *Mouse) SetCableLen(cm int) {
	m.CableLen = cm
}
