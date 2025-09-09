package main

import (
	"database/sql"
	"log/slog"
	"net"
	"os"

	_ "github.com/mattn/go-sqlite3"
	"google.golang.org/grpc"

	grpcService "github.com/oliverslade/todo-api/internal/grpc"
	"github.com/oliverslade/todo-api/internal/repository"
	todov1 "github.com/oliverslade/todo-api/proto/todo/v1"
)

func main() {
	db, err := sql.Open("sqlite3", "todo.db")
	if err != nil {
		slog.Error("failed to open db", "error", err)
		os.Exit(1)
	}
	defer db.Close()

	repo := repository.NewTodoRepo(db)
	todoHandler := grpcService.NewTodoHandler(repo)

	lis, err := net.Listen("tcp", ":8080")
	if err != nil {
		slog.Error("failed to listen", "error", err)
		os.Exit(1)
	}

	server := grpc.NewServer()
	todov1.RegisterTodoHandlerServer(server, todoHandler)

	slog.Info("starting server", "port", 8080)
	if err := server.Serve(lis); err != nil {
		slog.Error("failed to serve", "error", err)
		os.Exit(1)
	}
}
