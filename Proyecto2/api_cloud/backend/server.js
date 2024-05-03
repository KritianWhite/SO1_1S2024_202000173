const express = require("express");
const cors = require("cors");
const morgan = require("morgan");
const { MongoClient } = require("mongodb");

const app = express();
const PORT =  process.env.PORT || 2024;
const uri = 'mongodb://34.173.164.182:27017';
const dbName = 'votos';

app.use(cors());
app.use(morgan("dev"));
app.use(express.json());

const client = new MongoClient(uri);

// Configuración de opciones deprecated
client.connect({ useNewUrlParser: true, useUnifiedTopology: true })
  .then(() => {
    console.log("Conexión exitosa a MongoDB");
  })
  .catch((error) => {
    console.error("Error de conexión a MongoDB:", error);
    process.exit(1);
  });

// Ruta para obtener los últimos registros de la colección "dataset-votos"
app.get("/votos", async (req, res) => {
  try {
    const db = client.db(dbName);
    const votos = await db.collection('dataset-votos').find().limit(20).toArray();
    var votes = [];
    for (const voto of votos) {
      const datetimString = JSON.stringify(voto.inserted);
      //Dividir la cadena en fecha y hora utilizando el método split()
      const [datePart, timePart] = datetimString.replace(/"/g, '').split('T');
      voto.message = "{" + voto.message + ",\"fecha\":\"" + datePart + "\",\"hora\":\"" + timePart + "\"}";
      votes.push(JSON.parse(voto.message));
      console.log(voto.message);
    }
    res.json(votes);
  } catch (error) {
    console.error("Error al obtener los registros:", error);
    res.status(500).json({ error: "Error al obtener los registros" });
  }
});

app.listen(PORT, () => {
  console.log(`Servidor API en http://localhost:${PORT}`);
});
