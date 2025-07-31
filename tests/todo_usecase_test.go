package tests

import (
	"context"
	"errors"
	"github.com/gictorbit/ice/internal/todo/domain"
	"github.com/gictorbit/ice/internal/todo/usecase"
	mock_domain "github.com/gictorbit/ice/tests/mocks"
	"github.com/golang/mock/gomock"
	"github.com/google/uuid"
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
		wantErr       bool
	}{
		"success": {
			todo: &domain.TodoItem{
				ID:          uuid.New(),
				Description: "",
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
			wantErr: false,
		},
		"repo error": {
			todo: &domain.TodoItem{
				ID:          uuid.New(),
				Description: "",
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
			wantErr: true,
		},
		"publisher error": {
			todo: &domain.TodoItem{
				ID:          uuid.New(),
				Description: "",
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
			wantErr: true,
		},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			uc := usecase.NewTodoUseCase(tc.mockRepo(), tc.mockPublisher())
			err := uc.CreateTodo(context.Background(), tc.todo)
			if (err != nil) != tc.wantErr {
				t.Errorf("Create() error = %v, wantErr %v", err, tc.wantErr)
			}
		})
	}
}
