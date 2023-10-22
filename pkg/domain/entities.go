package entities

type Juego struct {
	Tablero Tablero
	Run     bool
}
type Tablero struct {
	Recuadros [][]Recuadro
}

type Recuadro struct {
	Vivo bool
	PosX int
	PosY int
}

type Adyacentes struct {
	Centro    Recuadro
	Alrededor [8]Recuadro
}

func CalcularVida(cuadro Adyacentes) bool {
	var vecinosVivos = 0
	for i := 0; i < len(cuadro.Alrededor); i++ {
		if cuadro.Alrededor[i].Vivo {
			vecinosVivos++
		}
	}
	if cuadro.Centro.Vivo {
		if vecinosVivos < 2 {
			return false
		}
		if vecinosVivos > 3 {
			return false
		}
		return true
	}
	if vecinosVivos == 3 {
		return true
	}
	return false
}

func CopyTablero(tablero Tablero, size int) Tablero {
	var recuadros = make([][]Recuadro, size)
	for i := range recuadros {
		recuadros[i] = make([]Recuadro, size)
	}
	var tableroTemp = Tablero{recuadros}
	for i := 0; i < len(tablero.Recuadros); i++ {
		for j := 0; j < len(tablero.Recuadros[i]); j++ {
			tableroTemp.Recuadros[i][j] = tablero.Recuadros[i][j]
		}
	}
	return tableroTemp
}

func TratarBordes(tablero Tablero, i int, j int) bool {
	invalidValue := -99
	var adyacentes = [8]Recuadro{}
	if i == 0 {
		adyacentes[0] = Recuadro{false, invalidValue, invalidValue}
		adyacentes[1] = Recuadro{false, invalidValue, invalidValue}
		adyacentes[2] = Recuadro{false, invalidValue, invalidValue}
	}
	if j == 0 {
		adyacentes[0] = Recuadro{false, invalidValue, invalidValue}
		adyacentes[3] = Recuadro{false, invalidValue, invalidValue}
		adyacentes[5] = Recuadro{false, invalidValue, invalidValue}
	}
	if i == len(tablero.Recuadros)-1 {
		adyacentes[5] = Recuadro{false, invalidValue, invalidValue}
		adyacentes[6] = Recuadro{false, invalidValue, invalidValue}
		adyacentes[7] = Recuadro{false, invalidValue, invalidValue}
	}
	if j == len(tablero.Recuadros[i])-1 {
		adyacentes[2] = Recuadro{false, invalidValue, invalidValue}
		adyacentes[4] = Recuadro{false, invalidValue, invalidValue}
		adyacentes[7] = Recuadro{false, invalidValue, invalidValue}
	}

	if adyacentes[0].PosX != invalidValue {
		adyacentes[0] = tablero.Recuadros[i-1][j-1]
	}
	if adyacentes[1].PosX != invalidValue {
		adyacentes[1] = tablero.Recuadros[i-1][j]
	}
	if adyacentes[2].PosX != invalidValue {
		adyacentes[2] = tablero.Recuadros[i-1][j+1]
	}
	if adyacentes[3].PosX != invalidValue {
		adyacentes[3] = tablero.Recuadros[i][j-1]
	}
	if adyacentes[4].PosX != invalidValue {
		adyacentes[4] = tablero.Recuadros[i][j+1]
	}
	if adyacentes[5].PosX != invalidValue {
		adyacentes[5] = tablero.Recuadros[i+1][j-1]
	}
	if adyacentes[6].PosX != invalidValue {
		adyacentes[6] = tablero.Recuadros[i+1][j]
	}
	if adyacentes[7].PosX != invalidValue {
		adyacentes[7] = tablero.Recuadros[i+1][j+1]
	}

	adyacentesCalculables := Adyacentes{
		Centro:    tablero.Recuadros[i][j],
		Alrededor: adyacentes,
	}
	return CalcularVida(adyacentesCalculables)
}
