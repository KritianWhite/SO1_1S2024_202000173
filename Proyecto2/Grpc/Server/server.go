package main

import (
	pb "SO1_1S2024_202000173/Proyecto2/Grpc/Server/Proto"

	"context"
	"fmt"
	"log"
	"net"

	"google.golang.org/grpc"
)


type server struct {
	pb.UnimplementedGetInfoServer
}

const (
	port = ":3001"
)

type Data struct {
	Name  string
	Album string
	Year  string
	Rank  string
}


func (s *server) ReturnInfo(ctx context.Context, in *pb.RequestId) (*pb.ReplyInfo, error) {
	fmt.Println("Recibí del cliente: ", in.GetName())
	data := Data{
		Name:      in.GetName(),
		Album:     in.GetAlbum(),
		Year:      in.GetYear(),
		Rank:      in.GetRank(),
	}
	fmt.Println(data)
	
	return &pb.ReplyInfo{Info: "Hola cliente, recibí la canción"}, nil
}

func main() {
	listen, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalln(err)
	}
	s := grpc.NewServer()
	pb.RegisterGetInfoServer(s, &server{})

	if err := s.Serve(listen); err != nil {
		log.Fatalln(err)
	}
}