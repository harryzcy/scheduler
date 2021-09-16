package messenger

import (
	"context"
	"fmt"
	"log"
	"net"
	"os"
	"os/signal"
	"syscall"

	"github.com/harryzcy/scheduler/internal/core"
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
	task := core.Task{
		Name:     in.GetName(),
		Command:  in.GetCommand(),
		Schedule: in.GetSchedule(),
		Once:     in.GetOnce(),
	}

	fmt.Println(task)

	core.AddTask(task)
	return &pb.Empty{}, nil
}

func (s *server) RemoveTask(ctx context.Context, in *pb.Task) (*pb.Empty, error) {
	return &pb.Empty{}, nil
}

func (s *server) RemoveAllTasks(ctx context.Context, in *pb.Empty) (*pb.Empty, error) {
	return &pb.Empty{}, nil
}

func StartServer() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterMessengerServer(s, &server{})

	sig := make(chan os.Signal, 1)
	signal.Notify(sig, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		if err := s.Serve(lis); err != nil {
			log.Fatalf("failed to serve: %v", err)
		}
	}()

	log.Printf("server listening at %v", lis.Addr())

	<-sig
	log.Print("server stopping gracefully...")

	s.GracefulStop()

	log.Print("server existed properly")

}
