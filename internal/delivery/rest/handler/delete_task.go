package handler

import (
	"context"
	"net/http"

	"github.com/labstack/echo/v4"
)

func (t *TaskHandler) DeleteTask(c echo.Context) error {
	userId, ok := c.Get("user_id").(string)
	if !ok || userId == "" {
		return c.JSON(http.StatusUnauthorized, map[string]string{
			"error": "Unauthorized",
		})
	}

	taskId := c.Param("id")
	if taskId == "" {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error": "Task ID is required",
		})
	}

	err := t.TaskSrv.DeleteTask(context.Background(), taskId, userId)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"error": "Failed to delete task",
		})
	}

	return c.JSON(http.StatusOK, map[string]string{
		"message": "Task deleted successfully",
	})

}
