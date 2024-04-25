// nombre de la base de datos
var dbName = 'proyecto2';

// usuario administrador con un nombre y contraseña personalizados
db.createUser({
  user: 'root',
  pwd: 'root',
  roles: ['readWrite', 'dbAdmin'],
  passwordDigestor: 'server',
});

// Usar la base de datos especificada
db = db.getSiblingDB(dbName);

// Crear una colección de ejemplo
db.createCollection('voto');