package auvasa

import (
	"errors"
	"net/http"
	"strconv"

	"gitlab.com/adrm/hipsterbot/Godeps/_workspace/src/github.com/yhat/scrape"

	"gitlab.com/adrm/hipsterbot/Godeps/_workspace/src/golang.org/x/net/html"
	"gitlab.com/adrm/hipsterbot/Godeps/_workspace/src/golang.org/x/net/html/atom"

	"gitlab.com/adrm/hipsterbot/Godeps/_workspace/src/golang.org/x/text/encoding/charmap"
	"gitlab.com/adrm/hipsterbot/Godeps/_workspace/src/golang.org/x/text/transform"
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
		return "", nil, errors.New("Error al conectar con el servidor de AUVASA.")
	}

	rInUTF8 := transform.NewReader(resp.Body, charmap.Windows1252.NewDecoder())
	root, err := html.Parse(rInUTF8)
	if err != nil {
		return "", nil, errors.New("Error en la respuesta de AUVASA.")
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

	if len(headers) >= 2 {
		return scrape.Text(headers[1]), resultados, nil
	}
	return "", nil, errors.New("La respuesta de AUVASA no se corresponde con lo esperado.")

}
