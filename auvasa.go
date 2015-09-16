package auvasa

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/yhat/scrape"

	"golang.org/x/net/html"
	"golang.org/x/net/html/atom"

	"golang.org/x/text/encoding/charmap"
	"golang.org/x/text/transform"
)

// ProximoBus describe un tiempo de llegada para un bus concreto.
type ProximoBus struct {
	Linea   string
	Destino string
	Minutos string
}

// Get devuelve el conjunto de tiempos de llegada para los buses de la parada
// dada. Hay que comprobar que no se devuelve error.
func Get(parada int) (string, []ProximoBus, error) {
	resp, err := http.Get("http://www.auvasa.es/paradamb.asp?codigo=" +
		strconv.Itoa(parada))
	if err != nil {
		return fail("Error al conectar con el servidor de AUVASA.")
	}

	rInUTF8 := transform.NewReader(resp.Body, charmap.Windows1252.NewDecoder())
	root, err := html.Parse(rInUTF8)
	if err != nil {
		return fail("Error en la respuesta de AUVASA.")
	}

	headers := scrape.FindAll(root, scrape.ByTag(atom.H1))
	lineasTiempos := scrape.FindAll(root, scrape.ByClass("style36"))
	var resultadosArray [100]ProximoBus
	resultados := resultadosArray[0:0]
	for _, item := range lineasTiempos {
		valores := scrape.FindAll(item, scrape.ByClass("style38"))
		resultados = append(resultados, ProximoBus{
			Linea:   scrape.Text(valores[0]),
			Destino: scrape.Text(valores[2]),
			Minutos: scrape.Text(valores[3]),
		})
	}

	if len(resultados) == 0 {
		return fail("No hay tiempos para la parada especificada.")
	}

	if len(headers) >= 2 {
		return scrape.Text(headers[1]), resultados, nil
	}
	return fail("La respuesta de AUVASA no se corresponde con lo esperado.")
}

func fail(msg string) (string, []ProximoBus, error) {
	return "", nil, errors.New(msg)
}
