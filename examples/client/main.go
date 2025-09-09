package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	todov1 "github.com/oliverslade/todo-api/proto/todo/v1"
)

func main() {
	conn, err := grpc.NewClient("localhost:8080", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("failed to connect: %v", err)
	}
	defer conn.Close()

	client := todov1.NewTodoHandlerClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()

	fmt.Println("=== Creating Todo ===")
	createResp, err := client.CreateTodo(ctx, &todov1.CreateTodoRequest{
		Message:    "Learn gRPC",
		IsFinished: false,
	})
	if err != nil {
		log.Fatalf("failed to create todo: %v", err)
	}
	fmt.Printf("Created todo: %+v\n", createResp.Todo)

	fmt.Println("\n=== Listing Todos ===")
	listResp, err := client.ListTodos(ctx, &todov1.ListTodosRequest{})
	if err != nil {
		log.Fatalf("failed to list todos: %v", err)
	}
	for _, todo := range listResp.Todos {
		fmt.Printf("Todo: %+v\n", todo)
	}

	if len(listResp.Todos) > 0 {
		fmt.Println("\n=== Getting Todo ===")
		todoID := listResp.Todos[0].Id
		getTodoResp, err := client.GetTodo(ctx, &todov1.GetTodoRequest{
			Id: todoID,
		})
		if err != nil {
			log.Fatalf("failed to get todo: %v", err)
		}
		fmt.Printf("Got todo: %+v\n", getTodoResp.Todo)

		fmt.Println("\n=== Updating Todo ===")
		updateResp, err := client.UpdateTodo(ctx, &todov1.UpdateTodoRequest{
			Id:         todoID,
			IsFinished: true,
		})
		if err != nil {
			log.Fatalf("failed to update todo: %v", err)
		}
		fmt.Printf("Updated todo: %+v\n", updateResp.Todo)
	}

	fmt.Println("\n=== All tests completed successfully! ===")
}
