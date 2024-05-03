// server/main.go
package main

import (
	"context"
	"fmt"
	pb "golang-producer/grpc"
	"log"
	"net"

	"github.com/segmentio/kafka-go"
	"google.golang.org/grpc"
)

type server struct {
	pb.UnimplementedGetInfoServer
}

type Data struct {
	Name  string
	Album string
	Year  string
	Rank  string
}

const (
	KafkaBroker = "my-cluster-kafka-bootstrap:9092"
	KafkaTopic  = "topic-votos"
	port        = ":5001"
)

func produceToKafka(data Data) error {
	// Configurar el escritor de Kafka
	writer := kafka.NewWriter(kafka.WriterConfig{
		Brokers: []string{KafkaBroker},
		Topic:   KafkaTopic,
	})

	// Escribir el mensaje en Kafka
	err := writer.WriteMessages(context.Background(), kafka.Message{
		Key:   []byte(data.Name),
		Value: []byte(fmt.Sprintf("{\"album\":\"%s\",\"name\":\"%s\",\"rank\":\"%s\",\"year\":\"%s\"}", data.Album, data.Name, data.Rank, data.Year)),
		//Value: []byte(fmt.Sprintf("name: %s, album: %s, year: %s, rank: %s", data.Name, data.Album, data.Year, data.Rank)),
	})
	if err != nil {
		return err
	}

	// Cerrar el escritor de Kafka
	writer.Close()
	return nil
}

func (s *server) ReturnInfo(ctx context.Context, in *pb.RequestId) (*pb.ReplyInfo, error) {
	fmt.Println("Receiving from customer")
	data := Data{
		Name:  in.GetName(),
		Album: in.GetAlbum(),
		Rank:  in.GetRank(),
		Year:  in.GetYear(),
	}
	fmt.Println(data)

	// Producir el mensaje en Kafka
	err := produceToKafka(data)
	if err != nil {
		fmt.Println("Error producing message in Kafka:", err)
		return nil, err
	}

	// Devolver la respuesta al cliente
	return &pb.ReplyInfo{Info: "Hello client, I received the album"}, nil
}

func main() {
	fmt.Println("Server running on port", port)
	listen, err := net.Listen("tcp", port)
	if err != nil {
		fmt.Println("Error listening:", err)
		log.Fatalln(err)
	}
	s := grpc.NewServer()
	pb.RegisterGetInfoServer(s, &server{})

	if err := s.Serve(listen); err != nil {
		fmt.Println("Error serving:", err)
		log.Fatalln(err)
	}
}
