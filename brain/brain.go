package brain;

import (
	"log"
	"net"

	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"

	pb "github.com/jtgans/squidbot-grpc"
)

type BrainServer struct {
	port string
	configFile string
	grpc *grpc.Server
}

const version = "0.1.0"

func NewServer(port string, configFile string) *BrainServer {
	return &BrainServer {
		port,
		configFile,
		nil,
	}
}

func (brain *BrainServer) Run() {
	log.Println("Squidbot Brain v" + version + " starting")

	var sock, err = net.Listen("tcp", "0.0.0.0:" + brain.port)
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	brain.grpc = grpc.NewServer()
	pb.RegisterBrainServer(brain.grpc, brain)
	reflection.Register(brain.grpc)

	log.Println("Squidbot Brain serving on port " + brain.port)

	if err := brain.grpc.Serve(sock); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}

func (s *BrainServer) OnNewMessage(ctx context.Context, in *pb.MessageRequest) (*pb.MessageResponse, error) {
	return &pb.MessageResponse{
		UniqueId: in.UniqueId,
		MessageBody: "Hi. I'mma squidbot!",
	}, nil
}
