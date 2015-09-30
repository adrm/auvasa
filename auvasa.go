package auvasa

import (
	"errors"
	"net/http"
	"strconv"
	"time"

	"github.com/yhat/scrape"

	"golang.org/x/net/html"
	"golang.org/x/net/html/atom"

	"golang.org/x/text/encoding/charmap"
	"golang.org/x/text/transform"
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

// Get devuelve el conjunto de tiempos de llegada para los buses de la parada
// dada. Hay que comprobar que no se devuelve error.
func Get(parada int) (TiemposParada, error) {
	resp, err := http.Get("http://www.auvasa.es/paradamb.asp?codigo=" +
		strconv.Itoa(parada))
	if err != nil {
		return TiemposParada{}, errors.New("Error al conectar con el servidor de AUVASA.")
	}

	rInUTF8 := transform.NewReader(resp.Body, charmap.Windows1252.NewDecoder())
	root, err := html.Parse(rInUTF8)
	if err != nil {
		return TiemposParada{}, errors.New("Error en la respuesta de AUVASA.")
	}

	headers := scrape.FindAll(root, scrape.ByTag(atom.H1))
	if len(headers) < 2 {
		return TiemposParada{}, errors.New("La parada indicada parece errónea.")
	}

	lineasTiempos := scrape.FindAll(root, scrape.ByClass("style36"))
	resultados := make([]ProximoBus, len(lineasTiempos))
	for i, item := range lineasTiempos {
		valores := scrape.FindAll(item, scrape.ByClass("style38"))
		resultados[i] = ProximoBus{
			Linea:   scrape.Text(valores[0]),
			Destino: scrape.Text(valores[2]),
			Minutos: scrape.Text(valores[3]),
		}
	}

	if len(resultados) == 0 {
		return TiemposParada{}, errors.New("No hay tiempos para la parada especificada. Puede que sea errónea o que ya no haya buses.")
	}

	return TiemposParada{
		Nombre:  scrape.Text(headers[1]),
		Tiempos: resultados,
		Momento: time.Now(),
		Codigo:  parada,
	}, nil

}
