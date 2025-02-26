package handler

import (
	"context"
	"net/http"

	"github.com/labstack/echo/v4"
	models "github.com/nottee-project/task_service/internal/models/task"
)

func (t *TaskHandler) ListTasks(c echo.Context) error {
	var params models.ListTasksParams

	userId, ok := c.Get("user_id").(string)
	if !ok || userId == "" {
		return c.JSON(http.StatusUnauthorized, map[string]string{
			"error": "Unauthorized",
		})
	}

	if err := c.Bind(&params); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error": "Invalid request body",
		})
	}

	params.UserId = userId

	tasks, err := t.TaskSrv.ListTasks(context.Background(), params)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"error": "Failed to retrieve tasks",
		})
	}

	return c.JSON(http.StatusOK, tasks)
}
