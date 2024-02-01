# TAREA 1 SO1 - 202000173


## BACKEND
Este código en Go (Golang) es un programa simple que crea un servidor web básico utilizando el paquete `net/http` y habilita el soporte para CORS (Cross-Origin Resource Sharing) mediante el uso del paquete `github.com/rs/cors`.

1. **Importación de paquetes:**
   ```go
   import (
       "fmt"
       "net/http"
       "github.com/rs/cors"
   )
   ```
   - `fmt`: Paquete estándar de Go utilizado para la entrada/salida de datos.
   - `net/http`: Paquete estándar de Go para construir aplicaciones HTTP.
   - `github.com/rs/cors`: Un paquete externo utilizado para manejar configuración CORS.

2. **Función `main`:**
   ```go
   func main() {
       mux := http.NewServeMux()
   ```
   - Se crea un nuevo multiplexor de servidores (`http.ServeMux`). Un multiplexor es un enrutador simple que decide qué función se ejecutará en función de la ruta de la solicitud.

   ```go
       mux.HandleFunc("/data", func(w http.ResponseWriter, r *http.Request) {
           fmt.Fprintf(w, "Christian Blanco - 202000173")
       })
   ```
   - Se define un manejador de función para la ruta "/data". Cuando un cliente hace una solicitud a "/data", la función asociada se ejecutará y enviará la cadena "Christian Blanco - 202000173" como respuesta.

   ```go
       c := cors.AllowAll()
   ```
   - Se crea una instancia de la estructura `cors.Cors` con la configuración predeterminada que permite todas las solicitudes CORS.

   ```go
       handler := c.Handler(mux)
   ```
   - Se envuelve el multiplexor (`mux`) con el manejador CORS creado anteriormente. Esto habilita CORS para todas las rutas manejadas por el multiplexor.

   ```go
       fmt.Println("Servidor escuchando en el puerto 8080...")
       http.ListenAndServe(":8080", handler)
   ```
   - Se imprime un mensaje en la consola y se inicia el servidor en el puerto 8080 utilizando `http.ListenAndServe`. El manejador CORS se aplica a todas las solicitudes manejadas por el multiplexor.

El Dockerfile utilizado para construir una imagen de contenedor para una aplicación escrita en Go (Golang) es el siguiente:

1. **Base de la imagen:**
   ```dockerfile
   # Usa una imagen de Go como base
   FROM golang:1.21.6-alpine
   ```
   - Se utiliza la imagen oficial de Go en su versión 1.21.6-alpine como la base para la construcción del contenedor. Alpine es una distribución Linux ligera.

2. **Directorio de trabajo:**
   ```dockerfile
   # Establece el directorio de trabajo en la aplicación
   WORKDIR /app
   ```
   - Establece el directorio de trabajo en el contenedor como "/app". Es el directorio en el que se copiarán los archivos de la aplicación y donde se ejecutarán los comandos siguientes.

3. **Copia de archivos:**
   ```dockerfile
   # Copia los archivos de tu aplicación al contenedor
   COPY . .
   ```
   - Copia todos los archivos del directorio actual (donde reside el Dockerfile) al directorio de trabajo "/app" en el contenedor.

4. **Compilación de la aplicación:**
   ```dockerfile
   # Compila la aplicación
   RUN go build -o main .
   ```
   - Utiliza el comando `go build` para compilar la aplicación Go. El resultado se guarda en un archivo ejecutable llamado "main" en el mismo directorio de trabajo.

5. **Puerto expuesto:**
   ```dockerfile
   # Expone el puerto en el que se ejecuta la aplicación
   EXPOSE 8080
   ```
   - Indica que el contenedor escuchará en el puerto 8080. Sin embargo, esto no abre el puerto en el host, es simplemente informativo.

6. **Comando de ejecución:**
   ```dockerfile
   # Comando para ejecutar la aplicación
   CMD ["./main"]
   ```
   - Especifica el comando que se ejecutará cuando se inicie el contenedor. En este caso, ejecutará el archivo ejecutable "main" que fue compilado anteriormente.


## FRONTEND

Este fragmento de código es un componente funcional escrito en React para una aplicación de front-end.

1. **Importación de React y useState:**
   ```jsx
   import React, { useState } from 'react';
   ```
   - Importa React y la función `useState` desde la biblioteca de React. `useState` es un gancho (hook) que permite agregar estado a componentes funcionales.

2. **Importación de estilos y creación del componente `App`:**
   ```jsx
   import './App.css';

   function App() {
   ```
   - Importa un archivo de estilos (`App.css`) y luego define el componente funcional `App`.

3. **Definición de estados iniciales con `useState`:**
   ```jsx
     const [mensaje, setMensaje] = useState('');
     const [mostrarDatos, setMostrarDatos] = useState(false);
   ```
   - Se utilizan dos estados locales: `mensaje` para almacenar el mensaje obtenido del servidor y `mostrarDatos` para controlar si se debe mostrar o no el mensaje en la interfaz de usuario.

4. **Función `handleClick`:**
   ```jsx
     const handleClick = async () => {
       try {
         // Realizar una solicitud al servidor Go para obtener el mensaje
         const response = await fetch('http://localhost:8080/data');
         const data = await response.text();

         // Actualizar el estado del mensaje
         setMensaje(data);

         // Alternar el estado de mostrarDatos
         setMostrarDatos(!mostrarDatos);
       } catch (error) {
         console.error('Error al obtener el mensaje:', error);
       }
     };
   ```
   - Esta función asincrónica `handleClick` se llama cuando se hace clic en el botón. Realiza una solicitud HTTP a un servidor Go (presumiblemente el mismo servidor del código de servidor Go que proporcionaste anteriormente) para obtener datos.

   - Se utiliza `fetch` para hacer la solicitud y se maneja la respuesta. Luego, actualiza el estado `mensaje` con los datos obtenidos y alterna el estado `mostrarDatos` para mostrar u ocultar el mensaje en la interfaz de usuario.

5. **Renderizado del componente:**
   ```jsx
     return (
       <div className="App">
         <header className="App-header">
           <button className='btn-mostrar-ocultar' onClick={handleClick}>
             {mostrarDatos ? 'Ocultar datos' : 'Mostrar datos'}
           </button>
           {mostrarDatos && <h1>{mensaje}</h1>}
         </header>
       </div>
     );
   }
   ```
   - En el método `render` del componente, se devuelve JSX que representa la interfaz de usuario. Hay un botón que muestra "Mostrar datos" o "Ocultar datos" según el estado de `mostrarDatos`. Además, si `mostrarDatos` es `true`, se muestra un elemento `h1` que contiene el mensaje obtenido del servidor.


El Dockerfile utilizado para construir una imagen de contenedor para una aplicación de Node.js es el siguiente:

1. **Base de la imagen:**
   ```dockerfile
   # Usa una imagen de Node.js como base
   FROM node:16.20.2-alpine
   ```
   - Se utiliza la imagen oficial de Node.js en su versión 16.20.2-alpine como base para la construcción del contenedor. Alpine es una distribución Linux ligera.

2. **Directorio de trabajo:**
   ```dockerfile
   # Establece el directorio de trabajo en la aplicación
   WORKDIR /usr/src/app
   ```
   - Establece el directorio de trabajo en el contenedor como "/usr/src/app". Es el directorio en el que se copiarán los archivos de la aplicación y donde se ejecutarán los comandos siguientes.

3. **Copia de archivos del paquete y bloqueo:**
   ```dockerfile
   # Copia los archivos del paquete y el bloqueo para instalar las dependencias
   COPY package*.json ./
   ```
   - Copia los archivos `package.json` y `package-lock.json` del directorio actual (donde reside el Dockerfile) al directorio de trabajo en el contenedor.

4. **Instalación de dependencias:**
   ```dockerfile
   # Instala las dependencias
   RUN npm install
   ```
   - Utiliza el comando `npm install` para instalar las dependencias de la aplicación. Esto se basa en la información proporcionada por los archivos `package.json` y `package-lock.json`.

5. **Copia del resto de los archivos de la aplicación:**
   ```dockerfile
   # Copia los archivos del resto de la aplicación
   COPY . .
   ```
   - Copia todos los archivos del directorio actual (donde reside el Dockerfile) al directorio de trabajo en el contenedor. Esto incluirá los archivos de la aplicación aparte de los archivos de configuración y código fuente.

6. **Compilación de la aplicación:**
   ```dockerfile
   # Compila la aplicación
   RUN npm run build
   ```
   - Utiliza el comando `npm run build` para compilar la aplicación. Este comando generalmente se utiliza para construir una aplicación de React, Angular o similar.

7. **Puerto expuesto:**
   ```dockerfile
   # Expone el puerto en el que se ejecuta la aplicación
   EXPOSE 3000
   ```
   - Indica que el contenedor escuchará en el puerto 3000. Al igual que en el caso anterior, esto no abre el puerto en el host, es solo informativo.

8. **Comando de ejecución:**
   ```dockerfile
   # Inicia la aplicación
   CMD ["npm", "start"]
   ```
   - Especifica el comando que se ejecutará cuando se inicie el contenedor. En este caso, se inicia la aplicación con el comando `npm start`.



### [LINK DEL VIDEO DE YUTUB](https://youtu.be/TI-4p2TNIu4)
