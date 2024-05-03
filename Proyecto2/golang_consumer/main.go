// consumer/main.go
package main

import (
	"context"
	"os"
	"fmt"
	"log"
	"time"
	"syscall"
	"os/signal"
	"strings"
	//"encoding/json"
	"github.com/google/uuid"
	"github.com/redis/go-redis/v9"
	"github.com/segmentio/kafka-go"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var ctx = context.Background()

// Definir una estructura para deserializar el JSON
type Message struct {
	Name  string `json:"name"`
	Album string `json:"album"`
	Year  string `json:"year"`
	Rank  string `json:"rank"`
}

func main() {
	fmt.Println("Starting Kafka consumer...")

	// Configura la dirección del servidor Kafka
	kafkaBrokers := []string{"my-cluster-kafka-bootstrap:9092"}

	// Configura el topic del que deseas consumir mensajes
	topic := "topic-votos"

	// Configura la dirección del servidor Redis
	redisAddr := "redis-service:6379"

	// Crea un cliente Redis
	rdb := redis.NewClient(&redis.Options{
		Addr:     redisAddr,
		Password: "", // si no hay contraseña
		DB:       0,  // usa el DB predeterminado
	})

	// Cierra la conexión de Redis al finalizar
	defer rdb.Close()

	// Configura la dirección de conexión a MongoDB
	mongoURI := "mongodb://mongodb-service:27017"
	clientOptions := options.Client().ApplyURI(mongoURI)

	// Conecta con MongoDB
	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		fmt.Printf("Error connecting to MongoDB: %v\n", err)
		return
	}
	defer client.Disconnect(ctx)

	// Selecciona la base de datos y la colección en MongoDB
	collection := client.Database("votos").Collection("dataset-votos")

	// Crea un nuevo lector Kafka
	r := kafka.NewReader(kafka.ReaderConfig{
		Brokers:      kafkaBrokers,
		Topic:        topic,
		MinBytes:     10e3, // 10KB
		MaxBytes:     10e6, // 10MB
		MaxAttempts:  5,
		GroupID:      "group-id",
		StartOffset:  kafka.LastOffset,
	})

	// Cierra el lector de Kafka al finalizar
	defer r.Close()

	// Configura el canal para capturar las señales de terminación
	sigterm := make(chan os.Signal, 1)
	signal.Notify(sigterm, syscall.SIGINT, syscall.SIGTERM)

	// Inicia un bucle para leer los mensajes de Kafka
	for {
		select {
		case <-sigterm:
			// Señal de terminación recibida, salir del bucle
			fmt.Println("Received termination signal. Existing...")
			return
		default:
			// Lee un mensaje del topic
			message, err := r.ReadMessage(ctx)
			if err != nil {
				fmt.Printf("Error reading message: %v\n", err)
				return
			}

			// Convertir el mensaje en una cadena
			msgStr := string(message.Value)
			fmt.Printf("Received JSON message: %s\n", msgStr)

			msgStr = strings.Trim(msgStr, "{}")

			// Dividir el mensaje en partes usando la coma (`,`) como delimitador
			parts := strings.Split(msgStr, ",")
			// Variables para almacenar los valores de cada clave
			var name, album, year, rank string
			
			// Iterar sobre las partes y asignar los valores correspondientes a las variables
			for _, part := range parts {
				// Eliminar las comillas dobles de cada elemento
				cleanedPart := strings.ReplaceAll(part, `"`, "")

				// Dividir cada parte en clave y valor usando ":" como delimitador
				keyValue := strings.Split(cleanedPart, ":")
				if len(keyValue) != 2 {
					continue // Saltar si la parte no tiene un formato clave:valor válido
				}

				// Obtener la clave y el valor limpios
				key := strings.TrimSpace(keyValue[0])
				value := strings.TrimSpace(keyValue[1])

				// Asignar el valor a la variable correspondiente según la clave
				switch key {
				case "name":
					name = value
				case "album":
					album = value
				case "year":
					year = value
				case "rank":
					rank = value
				}
			}

			// Imprimir los valores asignados a cada variable
			fmt.Printf("Name-> %s\n", name)
			fmt.Printf("Album-> %s\n", album)
			fmt.Printf("Year-> %s\n", year)
			fmt.Printf("Rank-> %s\n", rank)

			uuid, err := uuid.NewRandom()
			if err != nil {
				log.Printf("Error al generar un UUID: %v", err)
				continue
			}

			// Generar la clave para Redis
			redisKey := name + "_" + album + "_" + year + "_" + rank + "_"+ uuid.String()
			fmt.Printf("Redis Key: %s\n", redisKey)

			// Incrementar el contador de votos en Redis
			err = rdb.HIncrBy(ctx,"votos",redisKey, 1).Err()
			if err != nil {
				fmt.Printf("Error incrementing vote count in Redis: %v\n", err)
				return
			}
			
			data := fmt.Sprintf(`uuid: "%s", album: "%s", year: "%s", name: "%s", rank: %s`, uuid, album, year, name, rank)
			_, err = rdb.LPush(ctx, "votos_list", data).Result()
			
			if err != nil {
				log.Printf("Error al guardar en Redis: %v", err)
				continue
			}

			fmt.Printf("Successfully incremented vote count for %s in Redis\n", redisKey)

			// Insertar el mensaje completo en MongoDB junto con la fecha y hora de inserción
			document := bson.M{
				"message":  msgStr,
				"inserted": time.Now(),
			}

			// Insertar el documento en la colección de MongoDB
			_, err = collection.InsertOne(ctx, document)
			if err != nil {
				fmt.Printf("Error inserting document into MongoDB: %v\n", err)
				return
			}

			fmt.Println("Successfully inserted document into MongoDB")
		}
	}
}
