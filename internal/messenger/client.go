package messenger

import (
	"context"
	"fmt"
	"io"
	"log"
	"time"

	pb "github.com/harryzcy/scheduler/internal/messenger/messenger"
	"google.golang.org/grpc"
)

const (
	address     = "localhost:50051"
	defaultName = "world"
)

func ListTask() {
	conn, err := grpc.Dial(address, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := pb.NewMessengerClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	stream, err := c.ListTasks(ctx, &pb.Empty{})
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}

	done := make(chan bool)

	go func() {
		for {
			resp, err := stream.Recv()
			if err == io.EOF {
				done <- true
				return
			}
			if err != nil {
				log.Fatalf("cannot receive %v", err)
			}
			fmt.Printf("%s %s %s %t", resp.GetName(), resp.GetCommand(), resp.GetSchedule(), resp.GetOnce())
		}
	}()

	<-done
}

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

func RemoveTask(name string) {
	conn, err := grpc.Dial(address, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("failed to connect: %v", err)
	}
	defer conn.Close()
	c := pb.NewMessengerClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	req := &pb.RemoveTaskRequest{
		Name: name,
	}
	_, err = c.RemoveTask(ctx, req)
	if err != nil {
		log.Fatalf("failed to remove the task: %v", err)
	}
}
