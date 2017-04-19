package auvasa

import (
	"time"
)

// TiemposParada agrupa los tiempos de llegada de los buses para una parada.
type TiemposParada struct {
	Nombre  string
	Codigo  int
	Momento time.Time
	Tiempos []ProximoBus
}

// ProximoBus describe un tiempo de llegada para un bus concreto.
type ProximoBus struct {
	Linea   string
	Destino string
	Minutos string
}
