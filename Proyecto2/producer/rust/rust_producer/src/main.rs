use actix_web::{get, post, web, App, HttpResponse, HttpServer, Responder};
use rdkafka::config::ClientConfig;
use rdkafka::producer::{FutureProducer, FutureRecord};
use serde_json::Value;
use std::time::Duration;

#[post("/mensaje")]
async fn recibir_mensaje(data: web::Json<Value>) -> impl Responder {
    // Obtener los valores de los campos del JSON recibido
    let name = data.get("name").and_then(Value::as_str).unwrap_or("");
    let album = data.get("album").and_then(Value::as_str).unwrap_or("");
    let year = data.get("year").and_then(Value::as_str).unwrap_or("");
    let rank = data.get("rank").and_then(Value::as_str).unwrap_or("");

    // Imprimir los valores en la consola del servidor
    println!("Name: {}", name);
    println!("Album: {}", album);
    println!("Year: {}", year);
    println!("Rank: {}", rank);

    // Configurar el productor de Kafka
    let producer: FutureProducer = ClientConfig::new()
        .set("bootstrap.servers", "my-cluster-kafka-bootstrap:9092") // Cambiar si es necesario
        .set("message.timeout.ms", "5000")
        .create()
        .expect("Producer creation error");

    // Serializar el mensaje JSON
    let message = serde_json::to_string(&data.0)
        .expect("Error serializing JSON message");

    // Publicar el mensaje en Kafka
    match     producer
    .send(
        FutureRecord::to("topic-votos")
            .key("") // No se establece una clave en este ejemplo
            .payload(&message),
        Duration::from_secs(0), // Partición a la que enviar el mensaje
     // Partición a la que enviar el mensaje
    ).await {
        Ok(_) => {
            // Responder al cliente con un mensaje de confirmación
            actix_web::HttpResponse::Ok().body("Mensaje recibido en el servidor")
        }
        Err(err) => {
            // Manejar el error al enviar el mensaje a Kafka
            eprintln!("Error sending message to Kafka: {:?}", err);
            // Devolver un error HTTP al cliente
            actix_web::HttpResponse::InternalServerError().finish()
        }
    }
}


#[get("/ping")]
async fn ping() -> impl Responder {
    HttpResponse::Ok().body("Pong")
}

#[actix_web::main]
async fn main() -> std::io::Result<()> {
    // Mensaje de inicio
    println!("Iniciando la API... server");

    // Crear y ejecutar el servidor
    HttpServer::new(|| {
        App::new()
            .service(recibir_mensaje)
            .service(ping)
    })
    .bind("0.0.0.0:5004")?
    .run()
    .await
}
