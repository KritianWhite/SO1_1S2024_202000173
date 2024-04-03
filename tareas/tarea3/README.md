# TAREA 3 | 202000173

## SUBSCRIBED

1. Importaciones:
   ```python
   import time, redis
   ```
   - Se importan los módulos `time` y `redis`. El módulo `time` se utiliza para manejar el tiempo y el módulo `redis` es una biblioteca de Python que proporciona soporte para trabajar con la base de datos de Redis, un almacén de estructura de datos en memoria.

2. Definición de la clase `RedisSub`:
   ```python
   class RedisSub:
       def __init__(self, host, port, db=0):
           self.r = redis.Redis(host=host, port=port, db=db, decode_responses=True)
           self.pubsub = self.r.pubsub()
   ```
   - Se define la clase `RedisSub`, que se utiliza para realizar la suscripción a un canal de Redis y recibir mensajes.
   - El método `__init__` inicializa la clase y crea una instancia de `redis.Redis` con la dirección del host, el puerto y la base de datos especificados. `decode_responses=True` se utiliza para decodificar las respuestas de Redis a cadenas Unicode.
   - Se crea un objeto `pubsub` que se utiliza para realizar la suscripción y recibir mensajes.

3. Método `subscribe`:
   ```python
       def subscribe(self, channel, callback):
           self.pubsub.subscribe(**channel)
           print('subscribed to', channel.keys())
           while True:
               message = self.pubsub.get_message()
               if message and not message['type'] == 'subscribe':
                   callback(message)
               time.sleep(0.02)
   ```
   - El método `subscribe` toma un diccionario `channel` que contiene el nombre del canal como clave y una función de `callback` como valor.
   - Se suscribe al canal utilizando `self.pubsub.subscribe(**channel)` y muestra un mensaje de confirmación.
   - Luego, inicia un bucle infinito donde obtiene mensajes del canal usando `self.pubsub.get_message()`.
   - Si el mensaje es válido (no es un mensaje de suscripción) llama a la función de `callback` proporcionada con el mensaje como argumento.
   - Espera brevemente antes de verificar nuevamente los mensajes en el canal.

4. Función `handle_message`:
   ```python
   def handle_message(message):
       print("received message:", {message['data']})
   ```
   - Define una función `handle_message` que toma un mensaje como argumento e imprime el mensaje recibido.

5. Ejecución principal:
   ```python
   if __name__ == '__main__':
       host = "10.22.69.155"
       port = 6379
       channel = {"test": handle_message}
       
       subscriber = RedisSub(host, port)
       subscriber.subscribe(channel, handle_message)
   ```
   - En la ejecución principal del programa, se define la dirección del host, el puerto y el canal al que suscribirse.
   - Se crea una instancia de `RedisSub` con la dirección del host y el puerto, y se llama al método `subscribe` con el canal y la función de `callback` para manejar los mensajes recibidos.

## PUBLISHED

1. Importación de módulos y definición de la función asincrónica:
   ```javascript
   const { createClient } = require('redis');
   
   (async () => {
       // More code here
   })();
   ```
   - Se importa la función `createClient` del módulo 'redis', que se utiliza para crear un cliente Redis en Node.js.
   - Se define una función asincrónica autoinvocada (IIFE) que contiene el código principal.

2. Creación y configuración del cliente Redis:
   ```javascript
   const client = createClient({
       socket: {
           host: '10.22.69.155',
           port: 6379,
       }
   });
   ```
   - Se crea un cliente Redis utilizando `createClient` con la configuración del host y puerto proporcionados.

3. Manejo de errores del cliente Redis:
   ```javascript
   client.on('error', (error) => {
       console.error("Error on redis: ",error);
   });
   ```
   - Se agrega un evento para manejar errores que puedan ocurrir en el cliente Redis. Cuando se produce un error, se imprime un mensaje de error en la consola.

4. Conexión al servidor Redis:
   ```javascript
   await client.connect();
   console.log("Redis connected");
   ```
   - Se espera a que el cliente se conecte al servidor Redis utilizando `await client.connect()`. Una vez conectado, se imprime un mensaje de confirmación en la consola.

5. Publicación periódica de mensajes:
   ```javascript
   setInterval(async () => {
       const msg = JSON.stringify({msg: "Hello, everyone!"});
       console.log("Publishing message: ",msg);
       try{
           const result = await client.publish("test: ",msg);
           console.log("Published message successfully: ",result);
       }catch(error){
           console.error("Error publishing message: ",error);
       }
   }, 3000);
   ```
   - Se utiliza `setInterval` para ejecutar una función asincrónica de manera periódica cada 3 segundos.
   - En cada iteración, se crea un mensaje JSON y se publica en el canal "test" utilizando `client.publish`.
   - Se imprime el mensaje publicado y se manejan posibles errores que puedan ocurrir durante la publicación.

## [LINK DEL VIDEO DE YUTUB]()