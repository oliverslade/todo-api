package grpc

import (
	"context"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/oliverslade/todo-api/internal/repository"
	todov1 "github.com/oliverslade/todo-api/proto/todo/v1"
)

type TodoHandler struct {
	todov1.UnimplementedTodoHandlerServer
	repo repository.TodoRepository
}

func NewTodoHandler(repo repository.TodoRepository) *TodoHandler {
	return &TodoHandler{
		repo: repo,
	}
}

func (h *TodoHandler) ListTodos(ctx context.Context, req *todov1.ListTodosRequest) (*todov1.ListTodosResponse, error) {
	todos, err := h.repo.GetAllTodos(ctx)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to get todos: %v", err)
	}

	var pbTodos []*todov1.Todo
	for _, todo := range todos {
		pbTodos = append(pbTodos, &todov1.Todo{
			Id:         todo.ID,
			Message:    todo.Message,
			IsFinished: todo.IsFinished,
		})
	}

	return &todov1.ListTodosResponse{
		Todos: pbTodos,
	}, nil
}

func (h *TodoHandler) GetTodo(ctx context.Context, req *todov1.GetTodoRequest) (*todov1.GetTodoResponse, error) {
	if req.Id <= 0 {
		return nil, status.Errorf(codes.InvalidArgument, "id must be positive")
	}

	todo, err := h.repo.GetTodoById(ctx, req.Id)
	if err != nil {
		return nil, status.Errorf(codes.NotFound, "todo not found: %v", err)
	}

	return &todov1.GetTodoResponse{
		Todo: &todov1.Todo{
			Id:         todo.ID,
			Message:    todo.Message,
			IsFinished: todo.IsFinished,
		},
	}, nil
}

func (h *TodoHandler) CreateTodo(ctx context.Context, req *todov1.CreateTodoRequest) (*todov1.CreateTodoResponse, error) {
	if req.Message == "" {
		return nil, status.Errorf(codes.InvalidArgument, "message is required")
	}

	todo := repository.Todo{
		Message:    req.Message,
		IsFinished: req.IsFinished,
	}

	err := h.repo.CreateTodo(ctx, &todo)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to create todo: %v", err)
	}

	return &todov1.CreateTodoResponse{
		Todo: &todov1.Todo{
			Id:         todo.ID,
			Message:    todo.Message,
			IsFinished: todo.IsFinished,
		},
	}, nil
}

func (h *TodoHandler) UpdateTodo(ctx context.Context, req *todov1.UpdateTodoRequest) (*todov1.UpdateTodoResponse, error) {
	if req.Id <= 0 {
		return nil, status.Errorf(codes.InvalidArgument, "id must be positive")
	}

	err := h.repo.SetTodoFinished(ctx, req.Id, req.IsFinished)
	if err != nil {
		return nil, status.Errorf(codes.NotFound, "failed to update todo: %v", err)
	}

	todo, err := h.repo.GetTodoById(ctx, req.Id)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to get updated todo: %v", err)
	}

	return &todov1.UpdateTodoResponse{
		Todo: &todov1.Todo{
			Id:         todo.ID,
			Message:    todo.Message,
			IsFinished: todo.IsFinished,
		},
	}, nil
}
