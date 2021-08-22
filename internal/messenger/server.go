package messenger

import (
	"context"
	"log"
	"net"

	pb "github.com/harryzcy/scheduler/internal/messenger/messenger"
	"google.golang.org/grpc"
)

const (
	port = ":50051"
)

type server struct {
	pb.UnimplementedMessengerServer
}

func (s *server) ListTasks(in *pb.Empty, stream pb.Messenger_ListTasksServer) error {

	return nil
}

func (s *server) AddTask(ctx context.Context, in *pb.Task) (*pb.Empty, error) {
	log.Printf("Received: %v", in.GetName())
	return &pb.Empty{}, nil
}

func (s *server) RemoveTask(ctx context.Context, in *pb.Task) (*pb.Empty, error) {
	return &pb.Empty{}, nil
}

func (s *server) RemoveAllTasks(ctx context.Context, in *pb.Empty) (*pb.Empty, error) {
	return &pb.Empty{}, nil
}

func Start() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterMessengerServer(s, &server{})
	log.Printf("server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
