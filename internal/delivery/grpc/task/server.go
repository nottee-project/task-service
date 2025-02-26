package taskgrpc

import (
	"context"
	"log"

	"github.com/google/uuid"
	"github.com/nottee-project/task_service/internal/domain/models"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/timestamppb"

	taskv1 "github.com/nottee-project/protos/fleap_protos"
)

type Task interface {
	CreateTask(ctx context.Context, req models.CreateTaskRequest) (models.CreateTaskResponse, error)
	UpdateTask(ctx context.Context, req models.UpdateTaskRequest) (models.UpdateTaskResponse, error)
	GetTask(ctx context.Context, req models.GetTaskRequest) (models.GetTaskResponse, error)
	ListTasks(ctx context.Context, req models.ListTasksRequest) (models.ListTasksResponse, error)
	DeleteTask(ctx context.Context, req models.DeleteTaskRequest) (models.DeleteTaskResponse, error)
}

type serverAPI struct {
	taskv1.UnimplementedTaskServiceServer
	task Task
}

func Register(gRPCServer *grpc.Server, task Task) {
	taskv1.RegisterTaskServiceServer(gRPCServer, &serverAPI{task: task})
}

func (s *serverAPI) CreateTask(
	ctx context.Context,
	in *taskv1.CreateTaskRequest,
) (*taskv1.CreateTaskResponse, error) {

	params := models.CreateTaskRequest{
		Title: in.GetTitle(),
		Body:  in.GetBody(),
	}

	task, err := s.task.CreateTask(ctx, params)
	if err != nil {
		log.Printf("Failed to create task: %v", err)
		return nil, err
	}

	return &taskv1.CreateTaskResponse{
		Task: &taskv1.Task{
			Id:        task.Id,
			UserId:    task.UserId.String(),
			Title:     task.Title,
			Body:      task.Body,
			CreatedAt: timestamppb.New(task.CreatedAt),
			UpdatedAt: timestamppb.New(task.UpdatedAt),
		},
	}, nil
}

func (s *serverAPI) ListTasks(ctx context.Context, in *taskv1.ListTasksRequest) (*taskv1.ListTasksResponse, error) {
	log.Println("ListTasks called with:", in)

	userID, err := uuid.Parse(in.GetUserId())
	if err != nil {
		log.Printf("Invalid user ID: %v", err)
		return nil, status.Errorf(codes.InvalidArgument, "invalid user ID: %v", err)
	}

	params := models.ListTasksRequest{
		UserId: userID,
	}

	tasksResponse, err := s.task.ListTasks(ctx, params)
	if err != nil {
		log.Printf("Failed to list tasks: %v", err)
		return nil, status.Errorf(codes.Internal, "failed to list tasks: %v", err)
	}

	taskList := &taskv1.ListTasksResponse{
		Tasks: &taskv1.Tasks{
			Tasks: make([]*taskv1.Task, 0, len(tasksResponse.Tasks.Tasks)),
		},
	}

	for _, t := range tasksResponse.Tasks.Tasks {
		taskList.Tasks.Tasks = append(taskList.Tasks.Tasks, &taskv1.Task{
			Id:     t.Id,
			UserId: t.UserId.String(),
			Title:  t.Title,
			Body:   t.Body,
		})
	}

	return taskList, nil
}

func (s *serverAPI) GetTask(ctx context.Context, in *taskv1.GetTaskRequest) (*taskv1.GetTaskResponse, error) {
	log.Println("GetTask called with:", in)

	params := models.GetTaskRequest{
		TaskId: in.GetId(),
		UserId: in.GetUserId(),
	}

	task, err := s.task.GetTask(ctx, params)
	if err != nil {
		log.Printf("Failed to get task: %v", err)
		return nil, err
	}

	return &taskv1.GetTaskResponse{
		Id:     task.ID,
		UserId: task.UserID.String(),
		Title:  task.Title,
		Body:   task.Body,
	}, nil
}

func (s *serverAPI) UpdateTask(ctx context.Context, in *taskv1.UpdateTaskRequest) (*taskv1.UpdateTaskResponse, error) {
	log.Println("UpdateTask called with:", in)

	//TODO добавить в прото собственно поля для изменения

	params := models.UpdateTaskRequest{
		Id:     in.GetId(),
		UserId: uuid.MustParse(in.UserId),
	}

	task, err := s.task.UpdateTask(ctx, params)
	if err != nil {
		log.Printf("Failed to update task: %v", err)
		return nil, err
	}

	return &taskv1.UpdateTaskResponse{
		Id:        task.Id,
		UserId:    task.UserId.String(),
		Title:     task.Title,
		Body:      task.Body,
		UpdatedAt: timestamppb.New(task.UpdatedAt),
	}, nil
}

func (s *serverAPI) DeleteTask(ctx context.Context, in *taskv1.DeleteTaskRequest) (*taskv1.DeleteTaskResponse, error) {
	log.Println("DeleteTask called with:", in)

	params := models.DeleteTaskRequest{
		TaskId: in.GetId(),
		UserId: in.GetUserId(),
	}

	msg, err := s.task.DeleteTask(ctx, params)
	if err != nil {
		log.Printf("Failed to delete task: %v", err)
		return nil, err
	}

	return &taskv1.DeleteTaskResponse{Message: msg.Message}, nil
}
