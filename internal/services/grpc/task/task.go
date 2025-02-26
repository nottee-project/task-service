package task

import (
	"context"
	"log/slog"

	"github.com/nottee-project/task_service/internal/domain/models"
	"github.com/nottee-project/task_service/internal/lib/logger/sl"
)

type TaskService struct {
	log         *slog.Logger
	taskCreator TaskCreator
	taskGetter  TaskGetter
	tasksLister TasksLister
	taskUpdator TaskUpdator
	taskDeleter TaskDeleter
	// appProvider AppProvider
}

type TaskCreator interface {
	CreateTask(
		ctx context.Context,
		req models.CreateTaskRequest,
	) (models.CreateTaskResponse, error)
}
type TaskUpdator interface {
	UpdateTask(
		ctx context.Context,
		req models.UpdateTaskRequest,
	) (models.UpdateTaskResponse, error)
}
type TaskGetter interface {
	GetTask(
		ctx context.Context,
		req models.GetTaskRequest,
	) (models.GetTaskResponse, error)
}
type TasksLister interface {
	ListTasks(
		ctx context.Context,
		req models.ListTasksRequest,
	) (models.ListTasksResponse, error)
}
type TaskDeleter interface {
	DeleteTask(
		ctx context.Context,
		req models.DeleteTaskRequest,
	) (models.DeleteTaskResponse, error)
}

type AppProvider interface {
	App(ctx context.Context, appID int) (models.App, error)
}

func New(
	log *slog.Logger,
	taskCreator TaskCreator,
	taskGetter TaskGetter,
	tasksLister TasksLister,
	taskUpdator TaskUpdator,
	taskDeleter TaskDeleter,
	// appProvider AppProvider,
) *TaskService {
	return &TaskService{
		log:         log,
		taskCreator: taskCreator,
		taskGetter:  taskGetter,
		tasksLister: tasksLister,
		taskUpdator: taskUpdator,
		taskDeleter: taskDeleter,
		// appProvider: appProvider,
	}
}

func (t *TaskService) CreateTask(ctx context.Context, req models.CreateTaskRequest) (res models.CreateTaskResponse, err error) {
	const op = "Task.CreateTask"

	log := t.log.With(
		slog.String("op", op),
	)

	log.Info("attempting to create task")

	res, err = t.taskCreator.CreateTask(ctx, req)
	if err != nil {
		log.Error("failed to create task", sl.Err(err))

		return models.CreateTaskResponse{}, err
	}

	return res, nil
}

func (t *TaskService) UpdateTask(ctx context.Context, req models.UpdateTaskRequest) (res models.UpdateTaskResponse, err error) {
	const op = "Task.UpdateTask"

	log := t.log.With(
		slog.String("op", op),
	)

	log.Info("attempting to update task")

	res, err = t.taskUpdator.UpdateTask(ctx, req)
	if err != nil {
		log.Error("failed to create task", sl.Err(err))

		return models.UpdateTaskResponse{}, err
	}

	return res, nil
}

func (t *TaskService) GetTask(ctx context.Context, req models.GetTaskRequest) (res models.GetTaskResponse, err error) {
	const op = "Task.GetTask"

	log := t.log.With(
		slog.String("op", op),
	)

	log.Info("attempting to update task")

	res, err = t.taskGetter.GetTask(ctx, req)
	if err != nil {
		log.Error("failed to create task", sl.Err(err))

		return models.GetTaskResponse{}, err
	}

	return res, nil
}

func (t *TaskService) ListTasks(ctx context.Context, req models.ListTasksRequest) (res models.ListTasksResponse, err error) {
	const op = "Task.ListTasks"

	log := t.log.With(
		slog.String("op", op),
	)

	log.Info("attempting to update task")

	res, err = t.tasksLister.ListTasks(ctx, req)
	if err != nil {
		log.Error("failed to create task", sl.Err(err))

		return models.ListTasksResponse{}, err
	}

	return res, nil
}

func (t *TaskService) DeleteTask(ctx context.Context, req models.DeleteTaskRequest) (res models.DeleteTaskResponse, err error) {
	const op = "Task.DeleteTask"

	log := t.log.With(
		slog.String("op", op),
	)

	log.Info("attempting to update task")

	res, err = t.taskDeleter.DeleteTask(ctx, req)
	if err != nil {
		log.Error("failed to create task", sl.Err(err))

		return models.DeleteTaskResponse{}, err
	}

	return res, nil
}
