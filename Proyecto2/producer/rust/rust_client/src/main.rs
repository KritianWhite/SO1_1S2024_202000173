use actix_web::{web, App, HttpResponse, HttpServer, Responder};
use serde::{Deserialize, Serialize};
use reqwest::Client;

// Estructura para deserializar el cuerpo JSON
#[derive(Debug, Deserialize, Serialize)]
struct RequestBody {
    name: String,
    album: String,
    year: String,
    rank: String,
}

async fn server_handler(body: web::Json<RequestBody>) -> impl Responder {
    // Realizar la solicitud HTTP al servicio externo
    println!("Cayendo solicITUD...");
    let client = Client::new();
    let response = client
        .post("http://localhost:5004/mensaje")
        .header("Content-Type", "application/json")
        .body(serde_json::to_string(&body.into_inner()).unwrap())
        .send()
        .await;

    match response {
        Ok(response) => {
            if response.status().is_success() {
                match response.text().await {
                    Ok(text) => HttpResponse::Ok().body(text),
                    Err(_) => HttpResponse::InternalServerError().finish(),
                }
            } else {
                // Imprimir el error en la consola
                println!("Error en la solicitud HTTP: {:?}", response.status());
                HttpResponse::InternalServerError().finish()
            }
        }
        Err(e) => {
            // Imprimir el error en la consola
            println!("Error en la solicitud HTTP: {}", e);
            HttpResponse::InternalServerError().finish()
        }
    }
}

async fn ping() -> impl Responder {
    HttpResponse::Ok().body("Pong")
}

#[actix_web::main]
async fn main() -> std::io::Result<()> {
    println!("Iniciando la API... client");
    HttpServer::new(|| {
        App::new()
            // Definir la ruta /server de tipo POST y manejarla con la función server_handler
            .route("/rust/server", web::post().to(server_handler))
            .route("/ping", web::get().to(ping))
    })
    .bind("0.0.0.0:5003")? // Cambia la dirección y el puerto según tu configuración
    .run()
    .await
}
