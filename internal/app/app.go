package app

import (
	"log/slog"

	"github.com/nottee-project/task_service/internal/adapter/store"
	taskStoreAdapter "github.com/nottee-project/task_service/internal/adapter/store/task"
	grpcapp "github.com/nottee-project/task_service/internal/app/grpc"
	"github.com/nottee-project/task_service/internal/config"
	"github.com/nottee-project/task_service/internal/services/grpc/task"
	"github.com/pkg/errors"
)

type App struct {
	GRPCServer *grpcapp.App
}

func New(log *slog.Logger, grpcPort int) (*App, error) {
	cfg, err := config.NewConfig()
	if err != nil {
		return nil, errors.Wrap(err, "failed to load config")
	}

	dbStore, err := store.New(cfg.Database)
	if err != nil {
		return nil, errors.Wrap(err, "failed to initialize database store")
	}

	taskStore := &taskStoreAdapter.TaskStore{Store: dbStore}

	taskService := task.New(log, taskStore, taskStore, taskStore, taskStore, taskStore)

	grpcApp := grpcapp.New(log, taskService, grpcPort)
	if grpcApp == nil {
		return nil, errors.New("failed to initialize GRPC server")
	}

	return &App{
		GRPCServer: grpcApp,
	}, nil
}
