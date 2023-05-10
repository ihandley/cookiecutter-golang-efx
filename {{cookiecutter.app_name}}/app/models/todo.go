package models

// Todo is a model
type Todo struct {
	Base
	Task *string `json:"task" validate:"required"`
	Done bool    `json:"done"`
}

// TableName It will returns table name
func (a *Todo) TableName() string {
	return "todos"
}

func init() {
}

// TodoQuery is non db operational model
type TodoQuery struct {
	Todo
	Pagination *Pagination
	Sort       *Order
}
