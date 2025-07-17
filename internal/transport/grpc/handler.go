package grpc

import (
	"context"
	"log"

	taskpb "github.com/Brrocat/project-protos/proto/task"
	userpb "github.com/Brrocat/project-protos/proto/user"
	"github.com/Brrocat/tasks-service/internal/task"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type Handler struct {
	svc        *task.Service
	userClient userpb.UserServiceClient
	taskpb.UnimplementedTaskServiceServer
}

func NewHandler(svc *task.Service, uc userpb.UserServiceClient) *Handler {
	return &Handler{
		svc:        svc,
		userClient: uc,
	}
}

func (h *Handler) CreateTask(ctx context.Context, req *taskpb.CreateTaskRequest) (*taskpb.CreateTaskResponse, error) {
	log.Printf("Creating task with title: %s for user: %d", req.GetTitle(), req.GetUserId())

	// 1. Проверяем существование пользователя
	_, err := h.userClient.GetUser(ctx, &userpb.User{Id: req.GetUserId()})
	if err != nil {
		log.Printf("User verification failed: %v", err)
		return nil, status.Errorf(codes.NotFound, "user verification failed: %v", err)
	}

	// 2. Создаем задачу
	newTask := &task.Task{
		Title:  req.GetTitle(),
		UserID: uint(req.GetUserId()), // Теперь поле существует
	}

	createdTask, err := h.svc.CreateTask(newTask)
	if err != nil {
		log.Printf("Failed to create task: %v", err)
		return nil, status.Error(codes.Internal, "failed to create task")
	}

	// 3. Формируем ответ
	return &taskpb.CreateTaskResponse{
		Task: &taskpb.Task{
			Id:    uint32(createdTask.ID),
			Title: createdTask.Title,
		},
	}, nil
}

func (h *Handler) GetTask(ctx context.Context, req *taskpb.Task) (*taskpb.Task, error) {
	log.Printf("Getting task with ID: %d", req.GetId())

	t, err := h.svc.GetTaskByID(uint(req.GetId()))
	if err != nil {
		log.Printf("Task not found: %v", err)
		return nil, status.Error(codes.NotFound, "task not found")
	}

	return &taskpb.Task{
		Id:    uint32(t.ID),
		Title: t.Title,
	}, nil
}

func (h *Handler) UpdateTask(ctx context.Context, req *taskpb.UpdateTaskRequest) (*taskpb.UpdateTaskResponse, error) {
	log.Printf("Updating task with ID: %d", req.GetId())

	// 1. Проверяем существование задачи
	existingTask, err := h.svc.GetTaskByID(uint(req.GetId()))
	if err != nil {
		log.Printf("Task not found: %v", err)
		return nil, status.Error(codes.NotFound, "task not found")
	}

	// 2. Обновляем задачу
	existingTask.Title = req.GetTitle()

	updatedTask, err := h.svc.UpdateTask(existingTask)
	if err != nil {
		log.Printf("Failed to update task: %v", err)
		return nil, status.Error(codes.Internal, "failed to update task")
	}

	return &taskpb.UpdateTaskResponse{
		Task: &taskpb.Task{
			Id:    uint32(updatedTask.ID),
			Title: updatedTask.Title,
		},
	}, nil
}

func (h *Handler) DeleteTask(ctx context.Context, req *taskpb.DeleteTaskRequest) (*taskpb.DeleteTaskResponse, error) {
	log.Printf("Deleting task with ID: %d", req.GetId())

	if err := h.svc.DeleteTask(uint(req.GetId())); err != nil {
		log.Printf("Failed to delete task: %v", err)
		return nil, status.Error(codes.Internal, "failed to delete task")
	}

	return &taskpb.DeleteTaskResponse{
		Success: true,
	}, nil
}

func (h *Handler) ListTasks(ctx context.Context, req *taskpb.ListTaskRequest) (*taskpb.ListTaskResponse, error) {
	log.Printf("Listing tasks with limit: %d, offset: %d", req.GetLimit(), req.GetOffset())

	tasks, err := h.svc.ListTasks(int(req.GetLimit()), int(req.GetOffset()))
	if err != nil {
		log.Printf("Failed to list tasks: %v", err)
		return nil, status.Error(codes.Internal, "failed to list tasks")
	}

	var pbTasks []*taskpb.Task
	for _, t := range tasks {
		pbTasks = append(pbTasks, &taskpb.Task{
			Id:    uint32(t.ID),
			Title: t.Title,
		})
	}

	return &taskpb.ListTaskResponse{
		Tasks: pbTasks,
	}, nil
}

func (h *Handler) ListUserTasks(ctx context.Context, req *taskpb.ListUserTasksRequest) (*taskpb.ListTaskResponse, error) {
	log.Printf("Listing tasks for user ID: %d", req.GetUserId())

	tasks, err := h.svc.ListUserTasks(
		uint(req.GetUserId()),
		int(req.GetLimit()),
		int(req.GetOffset()),
	)

	if err != nil {
		log.Printf("Failed to list user tasks: %v", err)
		return nil, status.Error(codes.Internal, "failed to list user tasks")
	}

	var pbTasks []*taskpb.Task
	for _, t := range tasks {
		pbTasks = append(pbTasks, &taskpb.Task{
			Id:     uint32(t.ID),
			Title:  t.Title,
			UserId: uint32(t.UserID),
		})
	}

	return &taskpb.ListTaskResponse{Tasks: pbTasks}, nil
}
