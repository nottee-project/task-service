package rest

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"

	"github.com/nottee-project/task_service/internal/bootstrap"
	"github.com/nottee-project/task_service/internal/delivery/rest/handler"
	mw_auth "github.com/nottee-project/task_service/internal/delivery/rest/middleware"
)

const prefix = "/api/v1"

func RegisterRoutes(e *echo.Echo, authServiceURL string) error {
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	taskSrv, err := bootstrap.CreateTaskService()
	if err != nil {
		return err
	}

	n := e.Group(prefix + "/task")
	n.Use(mw_auth.AuthMiddleware(authServiceURL))

	taskHandler := handler.NewTaskHandler(taskSrv)

	n.POST("", taskHandler.CreateTask)
	n.PUT("/:id", taskHandler.UpdateTask)
	n.GET("", taskHandler.ListTasks)
	n.GET("/:id", taskHandler.GetTask)
	n.DELETE("/:id", taskHandler.DeleteTask)

	// e.POST("/webhook", handler.TelegramWebhookHandler)

	e.GET("/test", taskHandler.TestTask)

	return nil
}
