# TAREA 2 SO1 - 202000173

## BACKEND


1. **Importación de módulos**: 

    ```javascript
    const express = require('express');
    const bodyParser = require('body-parser');
    const fs = require('fs');
    const path = require('path');
    const cors = require('cors');
    const mongoose = require('mongoose');
    ```
    Se importan los módulos necesarios para el servidor, como Express, body-parser, fs (para operaciones de sistema de archivos), path (para manipulación de rutas de archivos), cors (para permitir solicitudes de recursos desde otros dominios), y mongoose (para interactuar con MongoDB).


2. **Configuración de Express**: 

    ```javascript
    const app = express();
    app.use(cors());
    ```
    Se crea una instancia de la aplicación Express llamada `app`, y se configura para utilizar el middleware CORS para permitir solicitudes de recursos desde cualquier origen (esto puede ser útil si la aplicación va a ser consumida desde un cliente en un dominio diferente al del servidor).

3. **Configuración de body-parser**: 

    ```javascript
    app.use(bodyParser.json({ limit: '50mb' }));
    ```
    Se configura el middleware body-parser para analizar el cuerpo de las solicitudes HTTP en formato JSON con un límite de 50MB.


4. **Conexión a la base de datos MongoDB**:

    ```javascript
    mongoose.connect('mongodb://localhost:27017/tarea2', {
        useNewUrlParser: true,
        useUnifiedTopology: true,
    });
    ```
    Se utiliza Mongoose para conectar la aplicación a una base de datos MongoDB en el servidor local en el puerto 27017, y se especifica el nombre de la base de datos como 'tarea2'.


5. **Definición del esquema de la imagen**: 

    ```javascript
    const imageSchema = new mongoose.Schema({
        image: String, // Para la imagen en formato base64
        createdAt: { type: Date, default: Date.now }
    });
    ```
    Se define un esquema utilizando Mongoose para representar los datos de las imágenes que se guardarán en la base de datos. Este esquema incluye un campo `image` que almacenará la imagen en formato base64, y un campo `createdAt` que registrará la fecha y hora en que se guardó la imagen.


6. **Creación del modelo Image**: 

    ```javascript
    const Image = mongoose.model('Image', imageSchema);
    ```
    Se crea un modelo llamado `Image` utilizando el esquema definido anteriormente. Este modelo se utilizará para interactuar con la colección de imágenes en la base de datos.



7. **Endpoint para almacenar imágenes**: 

    ```javascript
    app.post('/upload', async (req, res) => {
        const base64Image = req.body.image.split(';base64,').pop();
        
        // Crear una instancia del modelo con la imagen
        const image = new Image({
            image: base64Image
        });

        // Guardar la instancia en la base de datos
        try {
            await image.save();
            res.send('Imagen guardada con éxito');
        } catch (error) {
            console.error(error);
            res.status(500).send('Error al guardar la imagen');
        }
    });
    ```
    Se define un endpoint `POST /upload` que acepta solicitudes para almacenar una imagen en la base de datos. Cuando se recibe una solicitud en este endpoint, se extrae la imagen en formato base64 del cuerpo de la solicitud, se crea una instancia del modelo `Image` con la imagen proporcionada, y se guarda en la base de datos. Se responde con un mensaje indicando si la imagen se guardó correctamente o si hubo un error.


8. **Endpoint para obtener la última imagen**: 

    ```javascript
    app.get('/latest-image', async (req, res) => {
        try {
            // Encuentra la última imagen añadida a la base de datos
            const latestImage = await Image.findOne().sort({ _id: -1 });

            if (latestImage) {
                res.json({ image: latestImage.image });
            } else {
                res.status(404).send('No se encontraron imágenes');
            }
        } catch (error) {
            console.error(error);
            res.status(500).send('Error al obtener la imagen');
        }
    });
    ```
    Se define un endpoint `GET /latest-image` que devuelve la última imagen añadida a la base de datos. Cuando se recibe una solicitud en este endpoint, se busca la última imagen en la base de datos ordenándolas por el campo `_id` de forma descendente. Si se encuentra una imagen, se devuelve su contenido en formato JSON. Si no se encuentra ninguna imagen, se responde con un mensaje de error.


9. **Inicio del servidor**: 

    ```javascript
    const PORT = 8000;
    app.listen(PORT, () => {
        console.log(`Servidor corriendo en el puerto ${PORT}`);
    });
    ```
    Se especifica el puerto en el que se ejecutará el servidor (8000 en este caso) y se inicia el servidor para que escuche las solicitudes entrantes en ese puerto. Se imprime un mensaje en la consola indicando que el servidor está corriendo y en qué puerto.


## FRONTEND

1. **Importación de módulos**:
    ```javascript
    import React, { useRef, useState, useEffect } from 'react';
    import './App.css';
    ```

    El código importa las funciones necesarias de React, como `useRef`, `useState`, `useEffect`, y también importa un archivo de estilo CSS (`App.css`).

2. **Componente funcional CameraApp**:
    ```javascript
    function CameraApp() {
    // Declaración de referencias a elementos DOM
    const videoRef = useRef(null);
    const canvasRef = useRef(null);
    // Estado para controlar si se está capturando una imagen
    const [capturing, setCapturing] = useState(false);
    ```
    Se define un componente funcional `CameraApp`. Dentro de este componente, se utilizan refs para acceder a los elementos `<video>` y `<canvas>` del DOM. También se utiliza un estado `capturing` para controlar si se está capturando una imagen.

3. **Efecto secundario con useEffect**:
    ```javascript
    useEffect(() => {
        getVideoStream();
        // Asegurarse de detener el stream de video cuando el componente se desmonte
        return () => {
        if (videoRef.current && videoRef.current.srcObject) {
            const tracks = videoRef.current.srcObject.getTracks();
            tracks.forEach(track => track.stop());
        }
        };
    }, []);
    ```
    Se utiliza `useEffect` para obtener acceso a la cámara del dispositivo y configurar el stream de video cuando el componente se monta por primera vez. Además, se especifica una función de limpieza para detener el stream de video cuando el componente se desmonta.

4. **Obtención del stream de video**:
    ```javascript
    const getVideoStream = async () => {
        try {
        const videoStream = await navigator.mediaDevices.getUserMedia({ video: true });
        videoRef.current.srcObject = videoStream;
        setCapturing(true);
        } catch (error) {
        console.error("Error accessing the camera", error);
        }
    };
    ```
    Esta función asincrónica `getVideoStream` utiliza `navigator.mediaDevices.getUserMedia` para obtener acceso al stream de video desde la cámara del dispositivo. Luego, establece este stream como la fuente de video del elemento `<video>`, y actualiza el estado `capturing` a `true` para indicar que se está capturando.

5. **Captura de una foto**:
    ```javascript
    const takePhoto = () => {
        const video = videoRef.current;
        const canvas = canvasRef.current;
        const context = canvas.getContext('2d');
        context.drawImage(video, 0, 0, canvas.width, canvas.height);

        // Convertir la imagen a base64
        const imageDataUrl = canvas.toDataURL('image/png');

        // Extraer la parte base64 de la cadena de datos de la imagen
        const base64Image = imageDataUrl.split(';base64,').pop();

        // Enviar la imagen al servidor
        fetch('http://localhost:8000/upload', {
        method: 'POST',
        headers: {
            'Content-Type': 'application/json',
        },
        body: JSON.stringify({ image: base64Image }),
        })
        .then(response => response.text())
        .then(data => {
            console.log('Success:', data);
        })
        .catch((error) => {
            console.error('Error:', error);
        });
    };
    ```
    Esta función `takePhoto` se ejecuta cuando se hace clic en el botón "Tomar Foto". Captura la imagen actual del stream de video utilizando el contexto del canvas, la convierte a formato base64, y luego envía la imagen al servidor utilizando una solicitud POST a la ruta '/upload'.

6. **Recuperación de la última imagen capturada**:
    ```javascript
    const getPhoto = () => {
        fetch('http://localhost:8000/latest-image')
        .then(response => response.json())
        .then(data => {
            let image = document.getElementById('latest-image');
            // Si no existe el elemento de imagen, lo crea y le asigna un id
            if (!image) {
            image = new Image();
            image.id = 'latest-image'; // Asigna un id al elemento de imagen
            image.className = 'latest-image';
            document.body.appendChild(image);
            }
            // Actualiza el atributo src de la imagen existente
            image.src = `data:image/png;base64,${data.image}`;
        })
        .catch((error) => {
            console.error('Error:', error);
        });
    };
    ```
    Esta función `getPhoto` se ejecuta cuando se hace clic en el botón "Obtener Foto". Realiza una solicitud GET a la ruta '/latest-image' en el servidor para obtener la última imagen capturada. Luego, crea un elemento `<img>` en el DOM (si no existe uno) y actualiza su atributo `src` con la imagen base64 recibida del servidor.

7. **Renderización del componente**:
    ```javascript
    return (
        <div>
        <video ref={videoRef} autoPlay hidden={!capturing}></video>
        <canvas ref={canvasRef} hidden={capturing}></canvas>
        <button onClick={takePhoto} disabled={!capturing}>Tomar Foto</button>
        <button onClick={getPhoto}>Obtener Foto</button>
        </div>
    );
    }
    ```
    Finalmente, el componente renderiza un contenedor `<div>` que contiene un elemento `<video>` (para mostrar el stream de video), un elemento `<canvas>` (para capturar la imagen), y dos botones para tomar y obtener la foto respectivamente. Dependiendo del estado `capturing`, se oculta el elemento `<canvas>` cuando no se está capturando una imagen, y se oculta el elemento `<video>` cuando se está capturando una imagen.

8. **Exportación del componente**:
    ```javascript
    export default CameraApp;
    ```
    El componente `CameraApp` se exporta como el componente predeterminado de este archivo, lo que permite su uso en otros componentes de la aplicación React.

## BASE DE DATOS

```bash
docker run -d -p 27017:27017 -v mongodata:/data/db --name mongodb mongo
```

El comando utilizado para ejecutar el contenedor Docker de MongoDB se explica acontinuación:

- `docker run`: Este comando se utiliza para ejecutar un contenedor a partir de una imagen Docker.
- `-d`: Esta bandera indica que deseas que el contenedor se ejecute en segundo plano (modo demonio).
- `-p 27017:27017`: Esta opción mapea el puerto 27017 del contenedor (puerto predeterminado de MongoDB) al puerto 27017 del host (tu máquina local). Esto permite acceder a MongoDB desde fuera del contenedor a través del puerto 27017 en tu máquina local.
- `-v mongodata:/data/db`: Esta opción monta un volumen Docker llamado "mongodata" en la ruta "/data/db" dentro del contenedor. Esto se hace para persistir los datos de MongoDB en tu sistema de archivos local, en lugar de almacenarlos solo en el contenedor, lo que los haría efímeros.
- `--name mongodb`: Con esta opción le das un nombre al contenedor, en este caso, "mongodb".
- `mongo`: Es el nombre de la imagen Docker que se utilizará para crear el contenedor. En este caso, es la imagen oficial de MongoDB disponible en Docker Hub.

## DOCKER COMPOSE

