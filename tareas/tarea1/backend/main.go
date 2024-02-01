package main

import (
    "fmt"
    "net/http"

    "github.com/rs/cors"
)

func main() {
    mux := http.NewServeMux()

    mux.HandleFunc("/data", func(w http.ResponseWriter, r *http.Request) {
        fmt.Fprintf(w, "Christian Blanco - 202000173")
    })

    // Configurar CORS
    c := cors.AllowAll()

    // Iniciar el servidor en el puerto 8000 con CORS habilitado
    handler := c.Handler(mux)
    fmt.Println("Servidor escuchando en el puerto 8080...")
    http.ListenAndServe(":8080", handler)
}
