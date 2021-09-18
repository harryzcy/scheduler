package messenger

import (
	"context"
	"log"
	"net"
	"os"
	"os/signal"
	"sync"
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
	log.Printf("received request ListTasks\n")

	tasks, _ := core.ListTasks()

	var wg sync.WaitGroup
	idx := 1
	for _, task := range tasks {
		wg.Add(1)

		t := &pb.Task{
			Name:     task.Name,
			Command:  task.Command,
			Schedule: task.Schedule,
			Once:     task.Once,
		}

		go func(count int, task *pb.Task) {
			defer wg.Done()
			err := stream.Send(task)

			if err == nil {
				log.Printf("sent response %d for ListTasks\n", count)
			} else {
				log.Printf("sending response %d for ListTasks failed\n", count)
			}
		}(idx, t)

		idx++
	}
	wg.Wait()
	log.Println("request ListTasks completed successfully")

	return nil
}

func (s *server) AddTask(ctx context.Context, in *pb.Task) (*pb.Empty, error) {
	task := &core.Task{
		Name:     in.GetName(),
		Command:  in.GetCommand(),
		Schedule: in.GetSchedule(),
		Once:     in.GetOnce(),
	}

	log.Printf("received request AddTask - name: %s, command: %s, schedule: %s, once: %t\n",
		task.Name, task.Command, task.Schedule, task.Once,
	)

	err := core.AddTask(task)

	if err == nil {
		log.Printf("request AddTask completed successfully - name: %s\n", task.Name)
	} else {
		log.Printf("request AddTask failed - name: %s, error: %s\n", task.Name, err)
	}

	return &pb.Empty{}, err
}

func (s *server) RemoveTask(ctx context.Context, in *pb.RemoveTaskRequest) (*pb.Empty, error) {
	log.Printf("received request RemoveTask - name: %s\n", in.GetName())

	err := core.RemoveTask(in.GetName())

	if err == nil {
		log.Printf("request RemoveTask completed successfully - name: %s\n", in.GetName())
	} else {
		log.Printf("request RemoveTask failed - name: %s, error: %s\n", in.GetName(), err)
	}

	return &pb.Empty{}, err
}

func (s *server) RemoveAllTasks(ctx context.Context, in *pb.Empty) (*pb.Empty, error) {
	log.Printf("received request RemoveAllTasks\n")

	core.RemoveAllTasks()

	log.Printf("request RemoveAllTasks completed successfully\n")

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

	core.StoreTasks()

}
