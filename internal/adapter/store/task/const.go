package tm

import "github.com/pkg/errors"

const (
	tableNameTasks = "tasks"

	fieldNameTaskId    = "id"
	fieldNameUserId    = "user_id"
	fieldNameTitle     = "title"
	fieldNameBody      = "body"
	fieldNameCreatedAt = "created_at"
	fieldNameUpdatedAt = "updated_at"
)

var TaskNotFound = errors.New("this task not found")
