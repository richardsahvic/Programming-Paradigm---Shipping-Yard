package ship

type kapal interface {
	Sail()
}

type commonData struct {
	name   string
	size   string
	date   string
	price  int
	sailed bool
}

func (data commonData) Sail() {
	data.sailed = true
}

// PerahuLayar class untuk tipe kapal perahu layar
type PerahuLayar struct {
	commonData
	sails int
	seat  int
}

// PerahuMotor class untuk tipe kapal perahu motor
type PerahuMotor struct {
	commonData
	fuelCapacity int
	seat         int
}

// KapalPesiar class untuk tipe kapal kapal pesiar
type KapalPesiar struct {
	commonData
	fuelCapacity int
	room         int
}
