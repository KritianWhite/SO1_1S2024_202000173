const express = require('express');
const bodyParser = require('body-parser');
const fs = require('fs');
const path = require('path');
const cors = require('cors');
const mongoose = require('mongoose');

const app = express();
app.use(cors());
// Configura body-parser para que maneje las peticiones con cuerpos en formato JSON
app.use(bodyParser.json({ limit: '50mb' }));

// Conecta a la base de datos
mongoose.connect('mongodb://localhost:27018/tarea2', {
    useNewUrlParser: true,
    useUnifiedTopology: true,
});

// Definimos el esquema de la imagen
const imageSchema = new mongoose.Schema({
    image: String, // Para la imagen en formato base64
    createdAt: { type: Date, default: Date.now }
});

const Image = mongoose.model('Image', imageSchema);


// Endpoint para almacenar la imagen en la base de datos
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

// Optiene la última imagen añadida a la base de datos
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

// Define el puerto y pon a escuchar al servidor
const PORT = 8000;
app.listen(PORT, () => {
    console.log(`Servidor corriendo en el puerto ${PORT}`);
});
