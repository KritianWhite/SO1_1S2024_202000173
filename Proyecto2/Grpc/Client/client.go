package main

import (
	pb "SO1_1S2024_202000173/Proyecto2/Grpc/Client/Proto" // Importa el paquete generado a partir de tu archivo .proto
	"context"
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/joho/godotenv"
	"github.com/segmentio/kafka-go"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

// CargarVariablesEntorno carga las variables de entorno desde el archivo .env
func CargarVariablesEntorno() {
	err := godotenv.Load()
	if err != nil {
		fmt.Println("No se encuentra el archivo...")
	}
}

var ctx = context.Background()

type BandsData struct {
	name  string
	album string
	year  string
	rank  string
}

func sendToKafka(voto BandsData) {
	// Configura el escritor de Kafka
	writer := &kafka.Writer{
		Addr:         kafka.TCP(os.Getenv("KAFKA_BROKER")),
		Topic:        os.Getenv("KAFKA_TOPIC"),
		Balancer:     &kafka.LeastBytes{},
		RequiredAcks: kafka.RequireAll,
	}
	defer writer.Close()

	// Datos a enviar como un mensaje a Kafka
	message := fmt.Sprintf("name: %s, album: %s, year: %s, rank: %s", voto.name, voto.album, voto.year, voto.rank)

	// Crea un mensaje Kafka y lo envía
	msg := kafka.Message{
		Key:   []byte(voto.name),
		Value: []byte(message),
	}

	if err := writer.WriteMessages(ctx, msg); err != nil {
		log.Fatalf("Error al enviar datos a Kafka: %s", err)
	}
	fmt.Println("Datos enviados a Kafka exitosamente.")
}

func insertData(c *fiber.Ctx) error {
	var data map[string]string
	e := c.BodyParser(&data)
	if e != nil {
		return e
	}

	voto := BandsData{
		name:  data["name"],
		album: data["album"],
		year:  data["year"],
		rank:  data["rank"],
	}

	go sendServer(voto)
	fmt.Println("SE ENVIARAN LOS DATOS A KAFKA...")
	sendToKafka(voto)
	return nil
}

func sendServer(voto BandsData) {
	client_host := obtenerHostCliente()
	client_port := obtenerPuertoCliente()
	fmt.Printf("La conexion con el cliente es en %s:%d\n", client_host, client_port)
	serverAddr := fmt.Sprintf("%s:%d", client_host, client_port)
	conn, err := grpc.Dial(serverAddr, grpc.WithTransportCredentials(insecure.NewCredentials()),
		grpc.WithBlock())
	if err != nil {
		log.Fatalln(err)
	}

	cl := pb.NewBandServiceClient(conn)
	defer func(conn *grpc.ClientConn) {
		err := conn.Close()
		if err != nil {
			log.Fatalln(err)
		}
	}(conn)

	ret, err := cl.SendBandInfo(ctx, &pb.Band{
		Name:  voto.name,
		Album: voto.album,
		Year:  voto.year,
		Rank:  voto.rank,
	})
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Println("Respuesta del server " + ret.GetMessage())
}

func getMessage(c *fiber.Ctx) error {
	// Devuelve una respuesta apropiada
	return c.JSON(fiber.Map{
		"message": "Se recupero la informacion desde el servidor",
	})
}

func obtenerHostCliente() string {
	HostStr := os.Getenv("CLIENT_HOST")
	return HostStr
}

func obtenerPuertoCliente() int {
	portStr := os.Getenv("CLIENT_PORT")
	port, err := strconv.Atoi(portStr)
	if err != nil {
		// Manejar el error en caso de que el valor de la variable de entorno no sea un número
		fmt.Println("Error al convertir el valor de PORT a un número:", err)
	}
	return port
}

func obtenerPuertoServer() int {
	portStr := os.Getenv("SERVER_PORT")
	port, err := strconv.Atoi(portStr)
	if err != nil {
		// Manejar el error en caso de que el valor de la variable de entorno no sea un número
		fmt.Println("Error al convertir el valor de PORT a un número:", err)
	}
	return port
}

func main() {
	CargarVariablesEntorno()
	server_port := obtenerPuertoServer()
	fmt.Printf("Transmitiendo por el puerto: %d\n", server_port)
	app := fiber.New()

	// Middleware para CORS
	app.Use(cors.New(cors.Config{
		AllowOrigins: "*",
		AllowMethods: "GET,POST",
	}))

	app.Get("/", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{
			"res": "todo bien",
		})
	})
	app.Post("/insert", insertData)

	// Endpoint GET para recibir mensajes
	app.Get("/receive", getMessage)

	err := app.Listen(fmt.Sprintf(":%d", server_port))
	if err != nil {
		return
	}
}
