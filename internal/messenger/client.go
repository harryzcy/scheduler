package messenger

import (
	"context"
	"log"
	"time"

	pb "github.com/harryzcy/scheduler/internal/messenger/messenger"
	"google.golang.org/grpc"
)

const (
	address     = "localhost:50051"
	defaultName = "world"
)

func AddTask(name, command, schedule string, once bool) {
	conn, err := grpc.Dial(address, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("failed to connect: %v", err)
	}
	defer conn.Close()
	c := pb.NewMessengerClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	task := &pb.Task{
		Name:     name,
		Command:  command,
		Schedule: schedule,
		Once:     once,
	}
	_, err = c.AddTask(ctx, task)
	if err != nil {
		log.Fatalf("failed to add the task: %v", err)
	}
}
