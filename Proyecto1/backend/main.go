package main

import (
	"fmt"
	"os/exec"
	"strconv"
	"strings"
    "encoding/json"
)

type InformacionProcesos struct {
	Porcentaje string `json:"porcentaje"`
	Procesos   string `json:"procesos"`
}

func main() {
	fmt.Println("¡Hola, mundo!")
	// Llamada a la función Ram
	Ram()
	CPU()
}

func Ram() int {
	// Nombre del archivo
	archivoProc := "/proc/ram_so1_1s2024"
	// Creación del comando cat
	comando := exec.Command("cat", archivoProc)
	// Captura de la salida del comando (Contenido del archivo)[Se obtiene en array de bytes]
	salida, err := comando.Output()
	if err != nil {
		fmt.Println("Error al acedder a la información de la ram.", err)
		return 0
	}
	// Convertir la salida de bytes a una cadena
	salidaStr := string(salida)

	// Parsear la cadena a un número entero
	numero, err := strconv.Atoi(salidaStr)
	if err != nil {
		fmt.Println("Error al convertir la salida a entero:", err)
		return 0
	}

	fmt.Println(numero)

	return numero
}

func CPU() {
    salida, _ := exec.Command("cat", "/proc/cpu_so1_1s2024").Output()
	lineas := strings.Split(string(salida), "\n")

	info := InformacionProcesos{
		Porcentaje: strings.TrimPrefix(lineas[0], "Porcentaje de uso de CPU: "),
	}

	for _, linea := range lineas[1:] {
		if linea != "" {
			info.Procesos += linea + "\n"
		}
	}

	jsonData, _ := json.Marshal(info)
	fmt.Println(string(jsonData))

}