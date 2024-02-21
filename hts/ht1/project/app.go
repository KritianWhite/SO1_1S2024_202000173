package main

import (
	"context"
	"fmt"
	"os/exec"
	"strconv"
)

// App struct
type App struct {
	ctx context.Context
}

// NewApp creates a new App application struct
func NewApp() *App {
	return &App{}
}

// startup is called when the app starts. The context is saved
// so we can call the runtime methods
func (a *App) startup(ctx context.Context) {
	a.ctx = ctx
}

func (a *App) Ram() int {

	// Nombre del archivo
	archivoProc := "/proc/ram_202000173"
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
