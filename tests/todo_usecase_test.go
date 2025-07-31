package tests

import (
	"context"
	"errors"
	"github.com/gictorbit/ice/internal/todo/domain"
	"github.com/gictorbit/ice/internal/todo/usecase"
	mock_domain "github.com/gictorbit/ice/tests/mocks"
	"github.com/golang/mock/gomock"
	"github.com/google/uuid"
	"go.uber.org/zap"
	"gotest.tools/v3/assert"
	"testing"
	"time"
)

func TestTodoUseCase_Create(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	tests := map[string]struct {
		todo          *domain.TodoItem
		mockRepo      func() domain.TodoRepository
		mockPublisher func() domain.StreamPublisher
		wantErr       error
	}{
		"success": {
			todo: &domain.TodoItem{
				ID:          uuid.New(),
				Description: "success todo",
				DueDate:     time.Now().Add(24 * time.Hour),
			},
			mockRepo: func() domain.TodoRepository {
				m := mock_domain.NewMockTodoRepository(ctrl)
				m.EXPECT().CreateTodo(gomock.Any(), gomock.Any()).Return(nil)
				return m
			},
			mockPublisher: func() domain.StreamPublisher {
				m := mock_domain.NewMockStreamPublisher(ctrl)
				m.EXPECT().Publish(gomock.Any(), gomock.Any()).Return(nil)
				return m
			},
			wantErr: nil,
		},
		"repo error": {
			todo: &domain.TodoItem{
				ID:          uuid.New(),
				Description: "repo fail",
				DueDate:     time.Now().Add(24 * time.Hour),
			},
			mockRepo: func() domain.TodoRepository {
				m := mock_domain.NewMockTodoRepository(ctrl)
				m.EXPECT().CreateTodo(gomock.Any(), gomock.Any()).Return(errors.New("db error"))
				return m
			},
			mockPublisher: func() domain.StreamPublisher {
				return mock_domain.NewMockStreamPublisher(ctrl)
			},
			wantErr: errors.New("db error"),
		},
		"publisher error": {
			todo: &domain.TodoItem{
				ID:          uuid.New(),
				Description: "publish fail",
				DueDate:     time.Now().Add(24 * time.Hour),
			},
			mockRepo: func() domain.TodoRepository {
				m := mock_domain.NewMockTodoRepository(ctrl)
				m.EXPECT().CreateTodo(gomock.Any(), gomock.Any()).Return(nil)
				return m
			},
			mockPublisher: func() domain.StreamPublisher {
				m := mock_domain.NewMockStreamPublisher(ctrl)
				m.EXPECT().Publish(gomock.Any(), gomock.Any()).Return(errors.New("redis fail"))
				return m
			},
			wantErr: errors.New("redis fail"),
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {

			logger, err := zap.NewProduction()
			assert.NilError(t, err)

			uc := usecase.NewTodoUseCase(test.mockRepo(), test.mockPublisher(), logger)

			err = uc.CreateTodo(context.Background(), test.todo)
			if test.wantErr == nil {
				assert.NilError(t, err)
			} else {
				assert.Error(t, err, test.wantErr.Error())
			}
		})
	}
}
