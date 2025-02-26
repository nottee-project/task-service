package tm

import (
	"context"
	"database/sql"
	"fmt"
	"log"

	"github.com/Masterminds/squirrel"
	models "github.com/nottee-project/task_service/internal/models/task"
	"github.com/pkg/errors"

	store "github.com/nottee-project/task_service/internal/adapter/store"
)

type TaskStore struct {
	Store *store.Store
}

func (t *TaskStore) CreateTask(ctx context.Context, task models.CreateTaskParams) (models.Task, error) {
	result := models.Task{}

	sqlStr, args, err := squirrel.
		Insert(tableNameTasks).
		Columns(fieldNameUserId, fieldNameTitle, fieldNameBody).
		Values(task.UserId, task.Title, task.Body).
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
		return models.Task{}, errors.Wrap(err, "ToSql")
	}

	log.Printf("SQL Query: %s, Args: %v", sqlStr, args)

	err = t.Store.DB.Get(&result, sqlStr, args...)
	if err != nil {
		return models.Task{}, errors.Wrap(err, "Get")
	}

	return result, nil
}

func (t *TaskStore) UpdateTask(ctx context.Context, task models.Task) (models.UpdateTaskResponse, error) {
	result := models.UpdateTaskResponse{}

	sqlStr, args, err := squirrel.
		Update(tableNameTasks).
		Set(fieldNameTitle, task.Title).
		Set(fieldNameBody, task.Body).
		Set(fieldNameUpdatedAt, squirrel.Expr("NOW()")).
		Where(squirrel.Eq{
			fieldNameTaskId: task.Id,
			fieldNameUserId: task.UserId,
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

func (t *TaskStore) GetTask(ctx context.Context, taskId, userId string) (models.GetTaskResponse, error) {
	result := models.GetTaskResponse{}

	query, args, err := squirrel.Select(
		fieldNameTaskId, fieldNameUserId, fieldNameTitle, fieldNameBody,
	).
		From(tableNameTasks).
		Where(squirrel.Eq{
			fieldNameTaskId: taskId,
			fieldNameUserId: userId,
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

func (t *TaskStore) ListTasks(ctx context.Context, listParams models.ListTasksParams) ([]models.GetTaskResponse, error) {
	result := make([]models.GetTaskResponse, 0)

	query := squirrel.Select(
		fieldNameTaskId, fieldNameUserId, fieldNameTitle, fieldNameBody,
	).
		From(tableNameTasks).
		Where(squirrel.Eq{
			fieldNameUserId: listParams.UserId,
		}).
		PlaceholderFormat(squirrel.Dollar)

	sqlStr, args, err := query.ToSql()
	if err != nil {
		log.Printf("Error generating SQL: %v", err)
		return nil, errors.Wrap(err, "ToSql")
	}

	log.Printf("Generated SQL: %s, Args: %v", sqlStr, args)

	err = t.Store.DB.SelectContext(ctx, &result, sqlStr, args...)
	if err != nil {
		log.Printf("Error executing query: %v", err)
		return nil, errors.Wrap(err, "SelectContext")
	}

	return result, nil
}
func (t *TaskStore) DeleteTask(ctx context.Context, taskID, userID string) error {
	sqlStr, args, err := squirrel.
		Delete(tableNameTasks).
		Where(squirrel.Eq{
			fieldNameTaskId: taskID,
			fieldNameUserId: userID,
		}).
		PlaceholderFormat(squirrel.Dollar).
		ToSql()
	if err != nil {
		return errors.Wrap(err, "ToSql")
	}

	_, err = t.Store.DB.ExecContext(ctx, sqlStr, args...)
	if err != nil {
		return errors.Wrap(err, "ExecContext")
	}

	return nil
}
