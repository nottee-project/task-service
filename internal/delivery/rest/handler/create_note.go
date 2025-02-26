package handler

import (
	"context"
	"log"
	"net/http"

	"github.com/labstack/echo/v4"
	models "github.com/nottee-project/task_service/internal/models/task"
)

func (t *TaskHandler) CreateTask(c echo.Context) error {
	userId, ok := c.Get("user_id").(string)
	if !ok || userId == "" {
		log.Println("user_id is missing or invalid")
		return c.JSON(http.StatusUnauthorized, map[string]string{
			"error": "Unauthorized",
		})
	}

	log.Printf("user_id: %s", userId)

	var params models.CreateTaskParams
	if err := c.Bind(&params); err != nil {
		log.Printf("Failed to bind request: %v", err)
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error": "Invalid request body",
		})
	}

	log.Printf("CreateTaskParams: %+v", params)

	if params.Title == "" {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error": "Title is required",
		})
	}

	params.UserId = userId
	log.Printf("Final CreateTaskParams: %+v", params)

	task, err := t.TaskSrv.CreateTask(context.Background(), params)
	if err != nil {
		log.Printf("Failed to create task: %v", err)
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"error": "Failed to create task",
		})
	}

	return c.JSON(http.StatusCreated, task)
}
