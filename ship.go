package ship

// Abstraction
type kapal interface {
	Sail()
}

// Class
type commonData struct {
	name   string
	size   string
	date   string
	price  int
	sailed bool
}

// 				||
// INHERITENCE	||
// 			 	\/

// PerahuLayar class untuk tipe kapal perahu layar
type PerahuLayar struct {
	commonData
	sails int
	seat  int
}

// Sail ... Overriding for PerahuLayar
func (pl *PerahuLayar) Sail() {
	pl.sailed = true
}

// PerahuMotor class untuk tipe kapal perahu motor
type PerahuMotor struct {
	commonData
	fuelCapacity int
	seat         int
}

// Sail ... Overriding for PerahuMotor
func (pm *PerahuMotor) Sail() {
	pm.sailed = true
}

// KapalPesiar class untuk tipe kapal kapal pesiar
type KapalPesiar struct {
	commonData
	fuelCapacity int
	room         int
}

// Sail ... Overriding for KapalPesiar
func (kp *KapalPesiar) Sail() {
	kp.sailed = true
}
