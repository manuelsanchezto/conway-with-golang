package pages

import (
	entities "manu/htmx.conway/pkg/domain"
	db "manu/htmx.conway/pkg/persistence"
	"strconv"

	"github.com/labstack/echo/v4"
)

var size = 20
var tablero = db.MemoryObject(size).Get()
var run = false

func Index(c echo.Context) error {
	return c.Render(200, "index.html", entities.Juego{
		Tablero: tablero,
		Run:     run,
	})
}

func Kill(c echo.Context) error {
	x := c.Param("PosX")
	y := c.Param("PosY")

	xpos, errx := strconv.Atoi(x)
	if errx != nil {
		panic(errx)
	}
	ypos, erry := strconv.Atoi(y)
	if erry != nil {
		panic(erry)
	}

	tablero.Recuadros[xpos][ypos].Vivo = false

	return c.Render(200, "index.html", entities.Juego{
		Tablero: tablero,
		Run:     run,
	})
}

func Alive(c echo.Context) error {
	x := c.Param("PosX")
	y := c.Param("PosY")

	xpos, errx := strconv.Atoi(x)
	if errx != nil {
		panic(errx)
	}
	ypos, erry := strconv.Atoi(y)
	if erry != nil {
		panic(erry)
	}

	tablero.Recuadros[xpos][ypos].Vivo = true

	return c.Render(200, "index.html", entities.Juego{
		Tablero: tablero,
		Run:     run,
	})
}

func Advance(c echo.Context) error {
	var tableroTemp = entities.CopyTablero(tablero, size)
	for i := 0; i < len(tablero.Recuadros); i++ {
		for j := 0; j < len(tablero.Recuadros[i]); j++ {
			if i == 0 || j == 0 || i == (len(tablero.Recuadros)-1) || j == (len(tablero.Recuadros[i])-1) {
				tableroTemp.Recuadros[i][j].Vivo = entities.TratarBordes(tablero, i, j)
			} else {
				adyacentes := entities.Adyacentes{
					Centro: tablero.Recuadros[i][j],
					Alrededor: [8]entities.Recuadro{
						tablero.Recuadros[i-1][j-1],
						tablero.Recuadros[i-1][j],
						tablero.Recuadros[i-1][j+1],
						tablero.Recuadros[i][j-1],
						tablero.Recuadros[i][j+1],
						tablero.Recuadros[i+1][j-1],
						tablero.Recuadros[i+1][j],
						tablero.Recuadros[i+1][j+1],
					},
				}
				tableroTemp.Recuadros[i][j].Vivo = entities.CalcularVida(adyacentes)
			}
		}
	}
	tablero = tableroTemp
	return c.Render(200, "index.html", entities.Juego{
		Tablero: tablero,
		Run:     run,
	})
}

func Reset(c echo.Context) error {
	tablero = db.MemoryObject(size).Get()
	return c.Render(200, "index.html", entities.Juego{
		Tablero: tablero,
		Run:     run,
	})
}

func Start(c echo.Context) error {
	run = !run
	return c.Render(200, "index.html", entities.Juego{
		Tablero: tablero,
		Run:     run,
	})
}
