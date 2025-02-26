package bootstrap

// import (
// 	"github.com/nottee-project/task_service/internal/adapter/store"
// 	task_store "github.com/nottee-project/task_service/internal/adapter/store/task"
// 	"github.com/nottee-project/task_service/internal/config"
// 	"github.com/nottee-project/task_service/internal/service/rest/task"
// 	"github.com/pkg/errors"
// )

// package bootstrap

// import (
// 	"github.com/nottee-project/task_service/internal/adapter/store"
// 	taskStoreAdapter "github.com/nottee-project/task_service/internal/adapter/store/task"
// 	"github.com/nottee-project/task_service/internal/config"
// 	"github.com/nottee-project/task_service/internal/service/task"
// 	"github.com/pkg/errors"
// )

// // CreateTaskService инициализирует сервис задач.
// func CreateTaskService() (*task.TaskService, error) {
// 	// Загрузка новой конфигурации.
// 	cfg, err := config.LoadConfig()
// 	if err != nil {
// 		return nil, errors.Wrap(err, "Ошибка загрузки конфигурации")
// 	}

// 	// Инициализация хранилища базы данных с использованием настроек cfg.Database.
// 	dbStore, err := store.New(cfg.Database)
// 	if err != nil {
// 		return nil, err
// 	}

// 	// Создание адаптера хранилища задач.
// 	taskStore := &taskStoreAdapter.TaskStore{Store: dbStore}

// 	// Создание и возврат нового TaskService.
// 	// Обратите внимание: здесь для всех зависимостей используется один и тот же адаптер.
// 	return task.NewTaskService(
// 		// Здесь можно передать экземпляр логгера (например, настроенный *slog.Logger).
// 		// Например, logger,
// 		taskStore,
// 		taskStore,
// 		taskStore,
// 		taskStore,
// 		taskStore,
// 		taskStore,
// 	), nil
// }
