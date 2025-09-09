package grpc

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/oliverslade/todo-api/internal/repository/inmemory"
	todov1 "github.com/oliverslade/todo-api/proto/todo/v1"
)

func TestTodoHandler_CreateTodo(t *testing.T) {
	repo := inmemory.NewInMemoryTodoRepo()
	handler := NewTodoHandler(repo)
	ctx := context.Background()

	t.Run("should create a todo successfully with valid request", func(t *testing.T) {
		req := &todov1.CreateTodoRequest{
			Message:    "Test todo",
			IsFinished: false,
		}

		resp, err := handler.CreateTodo(ctx, req)
		require.NoError(t, err)
		assert.NotNil(t, resp)
		assert.NotNil(t, resp.Todo)
		assert.Equal(t, "Test todo", resp.Todo.Message)
		assert.False(t, resp.Todo.IsFinished)
		assert.Greater(t, resp.Todo.Id, int64(0))
	})

	t.Run("should return InvalidArgument when message is empty", func(t *testing.T) {
		req := &todov1.CreateTodoRequest{
			Message:    "",
			IsFinished: false,
		}

		resp, err := handler.CreateTodo(ctx, req)
		assert.Nil(t, resp)
		require.Error(t, err)

		st := status.Convert(err)
		assert.Equal(t, codes.InvalidArgument, st.Code())
		assert.Contains(t, st.Message(), "message is required")
	})
}

func TestTodoHandler_ListTodos(t *testing.T) {
	repo := inmemory.NewInMemoryTodoRepo()
	handler := NewTodoHandler(repo)
	ctx := context.Background()

	t.Run("should return all todos successfully", func(t *testing.T) {
		req := &todov1.ListTodosRequest{}

		resp, err := handler.ListTodos(ctx, req)
		require.NoError(t, err)
		assert.NotNil(t, resp)
		assert.Len(t, resp.Todos, 1) // In-memory repo has one default todo
		assert.Equal(t, "Buy groceries", resp.Todos[0].Message)
		assert.False(t, resp.Todos[0].IsFinished)
	})
}

func TestTodoHandler_GetTodo(t *testing.T) {
	repo := inmemory.NewInMemoryTodoRepo()
	handler := NewTodoHandler(repo)
	ctx := context.Background()

	t.Run("should return a todo successfully with valid id", func(t *testing.T) {
		req := &todov1.GetTodoRequest{Id: 1}

		resp, err := handler.GetTodo(ctx, req)
		require.NoError(t, err)
		assert.NotNil(t, resp)
		assert.NotNil(t, resp.Todo)
		assert.Equal(t, int64(1), resp.Todo.Id)
		assert.Equal(t, "Buy groceries", resp.Todo.Message)
		assert.False(t, resp.Todo.IsFinished)
	})

	t.Run("should return InvalidArgument when id is not positive", func(t *testing.T) {
		req := &todov1.GetTodoRequest{Id: 0}

		resp, err := handler.GetTodo(ctx, req)
		assert.Nil(t, resp)
		require.Error(t, err)

		st := status.Convert(err)
		assert.Equal(t, codes.InvalidArgument, st.Code())
		assert.Contains(t, st.Message(), "id must be positive")
	})

	t.Run("should return NotFound when todo does not exist", func(t *testing.T) {
		req := &todov1.GetTodoRequest{Id: 999}

		resp, err := handler.GetTodo(ctx, req)
		assert.Nil(t, resp)
		require.Error(t, err)

		st := status.Convert(err)
		assert.Equal(t, codes.NotFound, st.Code())
	})
}

func TestTodoHandler_UpdateTodo(t *testing.T) {
	repo := inmemory.NewInMemoryTodoRepo()
	handler := NewTodoHandler(repo)
	ctx := context.Background()

	t.Run("should update a todo successfully", func(t *testing.T) {
		req := &todov1.UpdateTodoRequest{
			Id:         1,
			IsFinished: true,
		}

		resp, err := handler.UpdateTodo(ctx, req)
		require.NoError(t, err)
		assert.NotNil(t, resp)
		assert.NotNil(t, resp.Todo)
		assert.Equal(t, int64(1), resp.Todo.Id)
		assert.True(t, resp.Todo.IsFinished)
	})

	t.Run("should return InvalidArgument when id is not positive", func(t *testing.T) {
		req := &todov1.UpdateTodoRequest{
			Id:         0,
			IsFinished: true,
		}

		resp, err := handler.UpdateTodo(ctx, req)
		assert.Nil(t, resp)
		require.Error(t, err)

		st := status.Convert(err)
		assert.Equal(t, codes.InvalidArgument, st.Code())
		assert.Contains(t, st.Message(), "id must be positive")
	})

	t.Run("should return NotFound when todo does not exist", func(t *testing.T) {
		req := &todov1.UpdateTodoRequest{
			Id:         999,
			IsFinished: true,
		}

		resp, err := handler.UpdateTodo(ctx, req)
		assert.Nil(t, resp)
		require.Error(t, err)

		st := status.Convert(err)
		assert.Equal(t, codes.NotFound, st.Code())
	})
}
