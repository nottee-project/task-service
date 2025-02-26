package handler

import (
	"context"
	"net/http"

	"github.com/labstack/echo/v4"
	models "github.com/nottee-project/task_service/internal/models/task"
)

func (t *TaskHandler) UpdateTask(c echo.Context) error {
	// userID, ok := c.Get("user_id").(string)
	// if !ok || userID == "" {
	// 	return c.JSON(http.StatusUnauthorized, map[string]string{
	// 		"error": "Unauthorized",
	// 	})
	// }

	taskID := c.Param("id")
	if taskID == "" {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error": "Task ID is required",
		})
	}

	var params models.Task
	if err := c.Bind(&params); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error": "Invalid request body",
		})
	}

	params.Id = taskID

	task, err := t.TaskSrv.UpdateTask(context.Background(), params)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"error": "Failed to update task",
		})
	}

	return c.JSON(http.StatusOK, task)
}
