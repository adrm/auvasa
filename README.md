# auvasa
Parser sencillo del servico de autobuses urbanos de Valladolid, programado en Go (golang).

# Instalación

```sh
go get github.com/adrm/auvasa
```

Y después, en el fichero donde lo vayas a usar, `import "github.com/adrm/auvasa"` y listo.


# Funcionalidades

## Obtener los tiempos de llegada de una parada en tiempo real

```go
auvasa.Get(numeroParada int) (TiemposParada, error)
```

Para usar esta función, indica el número que identifica a la parada y comprueba si devuelve un error. Si no, consulta el primer
objeto que se devuelve. Es un struct definido así:

```go
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
```

Y ahí está toda la información.


# Colaboraciones

Si tienes alguna idea para mejorar esta sencilla librería, no dudes en hacer un pull request. Algunas tareas sugeridas
pueden ser:

- [ ] Eliminar la dependencia con github.com/yhat/scrape
- [ ] ¿Se te ocurre alguna nueva funcionalidad?


# Licencia

El código está bajo una licencia GPLv3, convirtiéndolo en software libre.
