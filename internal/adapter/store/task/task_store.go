package task

import (
	"context"
	"database/sql"
	"fmt"
	"log"

	"github.com/Masterminds/squirrel"
	models "github.com/nottee-project/task_service/internal/domain/models"
	"github.com/pkg/errors"

	store "github.com/nottee-project/task_service/internal/adapter/store"
)

type TaskStore struct {
	Store *store.Store
}

func (t *TaskStore) CreateTask(ctx context.Context, payload models.CreateTaskRequest) (result models.CreateTaskResponse, err error) {
	sqlStr, args, err := squirrel.
		Insert(tableNameTasks).
		Columns(fieldNameUserId, fieldNameTitle, fieldNameBody).
		Values(payload.UserId, payload.Title, payload.Body).
		Suffix(fmt.Sprintf("RETURNING %s, %s, %s, %s, %s, %s",
			fieldNameTaskId,
			fieldNameUserId,
			fieldNameTitle,
			fieldNameBody,
			fieldNameCreatedAt,
			fieldNameUpdatedAt)).
		PlaceholderFormat(squirrel.Dollar).
		ToSql()
	if err != nil {
		return models.CreateTaskResponse{}, errors.Wrap(err, "ToSql")
	}

	log.Printf("SQL Query: %s, Args: %v", sqlStr, args)

	err = t.Store.DB.Get(&result, sqlStr, args...)
	if err != nil {
		return models.CreateTaskResponse{}, errors.Wrap(err, "Get")
	}

	return result, nil
}

func (t *TaskStore) UpdateTask(ctx context.Context, payload models.UpdateTaskRequest) (result models.UpdateTaskResponse, err error) {
	sqlStr, args, err := squirrel.
		Update(tableNameTasks).
		// Set(fieldNameTitle, task.Title).
		// Set(fieldNameBody, task.Body).
		Set(fieldNameUpdatedAt, squirrel.Expr("NOW()")).
		Where(squirrel.Eq{
			fieldNameTaskId: payload.Id,
			fieldNameUserId: payload.UserId,
		}).
		Suffix(fmt.Sprintf("RETURNING %s, %s, %s, %s, %s",
			fieldNameTaskId,
			fieldNameUserId,
			fieldNameTitle,
			fieldNameBody,
			fieldNameUpdatedAt)).
		PlaceholderFormat(squirrel.Dollar).
		ToSql()
	if err != nil {
		return models.UpdateTaskResponse{}, errors.Wrap(err, "ToSql")
	}

	err = t.Store.DB.QueryRow(sqlStr, args...).
		Scan(
			&result.Id,
			&result.UserId,
			&result.Title,
			&result.Body,
			&result.UpdatedAt,
		)
	if err != nil {
		if err == sql.ErrNoRows {
			return models.UpdateTaskResponse{}, errors.New("task not found")
		}
		return models.UpdateTaskResponse{}, errors.Wrap(err, "QueryRowContext")
	}

	return result, nil
}

func (t *TaskStore) GetTask(ctx context.Context, payload models.GetTaskRequest) (result models.GetTaskResponse, err error) {
	query, args, err := squirrel.Select(
		fieldNameTaskId, fieldNameUserId, fieldNameTitle, fieldNameBody,
	).
		From(tableNameTasks).
		Where(squirrel.Eq{
			fieldNameTaskId: payload.TaskId,
			fieldNameUserId: payload.UserId,
		}).
		PlaceholderFormat(squirrel.Dollar).
		ToSql()
	if err != nil {
		return models.GetTaskResponse{}, errors.Wrap(err, "ToSql")
	}

	err = t.Store.DB.Get(&result, query, args...)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return models.GetTaskResponse{}, TaskNotFound
		}
		return models.GetTaskResponse{}, errors.Wrap(err, "Get")
	}

	return result, nil
}

func (t *TaskStore) ListTasks(ctx context.Context, payload models.ListTasksRequest) (result models.ListTasksResponse, err error) {
	res := make([]models.Task, 0)

	query := squirrel.Select(
		fieldNameTaskId, fieldNameUserId, fieldNameTitle, fieldNameBody,
	).
		From(tableNameTasks).
		Where(squirrel.Eq{
			fieldNameUserId: payload.UserId,
		}).
		PlaceholderFormat(squirrel.Dollar)

	sqlStr, args, err := query.ToSql()
	if err != nil {
		log.Printf("Error generating SQL: %v", err)
		return models.ListTasksResponse{}, errors.Wrap(err, "ToSql")
	}

	log.Printf("Generated SQL: %s, Args: %v", sqlStr, args)

	err = t.Store.DB.SelectContext(ctx, &res, sqlStr, args...)
	if err != nil {
		log.Printf("Error executing query: %v", err)
		return models.ListTasksResponse{}, errors.Wrap(err, "SelectContext")
	}

	result = models.ListTasksResponse{
		Tasks: models.Tasks{
			Tasks: res,
		},
	}

	return result, nil
}

func (t *TaskStore) DeleteTask(ctx context.Context, payload models.DeleteTaskRequest) (models.DeleteTaskResponse, error) {

	sqlStr, args, err := squirrel.
		Delete(tableNameTasks).
		Where(squirrel.Eq{
			fieldNameTaskId: payload.TaskId,
			fieldNameUserId: payload.UserId,
		}).
		PlaceholderFormat(squirrel.Dollar).
		ToSql()
	if err != nil {
		return models.DeleteTaskResponse{}, errors.Wrap(err, "ошибка формирования SQL-запроса")
	}

	res, err := t.Store.DB.ExecContext(ctx, sqlStr, args...)
	if err != nil {
		return models.DeleteTaskResponse{}, errors.Wrap(err, "ошибка выполнения запроса")
	}

	rowsAffected, err := res.RowsAffected()
	if err != nil {
		return models.DeleteTaskResponse{}, errors.Wrap(err, "ошибка получения количества затронутых строк")
	}

	var message string
	if rowsAffected > 0 {
		message = "Задача успешно удалена"
	} else {
		message = "Задача не найдена или уже удалена"
	}

	response := models.DeleteTaskResponse{
		Message: message,
	}

	return response, nil
}
