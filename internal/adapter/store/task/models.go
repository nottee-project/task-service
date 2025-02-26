package tm

// import (
// 	"database/sql"
// 	"time"
// )

// type Task struct {
// 	Id        string    `db:"id"`
// 	Title     string    `db:"title"`
// 	Body      string    `db:"body"`
// 	Completed bool      `db:"completed"`
// 	CreatedAt time.Time `db:"created_at"`
// 	UpdatedAt time.Time `db:"updated_at"`
// }

// type Tasks struct {
// 	Tasks []Task
// }

// type ListTasksParams struct {
// 	Completed bool   `db:"completed"`
// 	Order     string `db:"order"`
// 	SortBy    string `db:"sort_by"`
// }