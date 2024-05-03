const { MongoClient } = require('mongodb');

// URL de conexión a tu base de datos MongoDB
const uri = 'mongodb://localhost:27017'; // Cambia esto según tu configuración

// Nombre de la base de datos
const dbName = 'nombre_de_tu_base_de_datos'; // Cambia esto según el nombre de tu base de datos

// Función para conectar con MongoDB
async function connectToMongoDB() {
  const client = new MongoClient(uri);

  try {
    // Conectarse al servidor MongoDB
    await client.connect();
    console.log('Conexión exitosa a MongoDB');

    // Seleccionar la base de datos
    const db = client.db(dbName);

    // Retornar la instancia de la base de datos
    return db;
  } catch (error) {
    console.error('Error al conectar a MongoDB', error);
    throw error;
  }
}

// Exportar la función para conectar a MongoDB
module.exports = { connectToMongoDB };
