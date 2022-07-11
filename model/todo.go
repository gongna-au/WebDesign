package model

import (
	"gorm.io/gorm"
)

type TodoModel struct {
	gorm.Model
	Title     string `json:"title"`
	Completed int    `json:"completed"`
}

// transformedTodo 代表格式化的 todo 结构体
type TransformedTodo struct {
	ID        uint   `json:"id"`
	Title     string `json:"title"`
	Completed bool   `json:"completed"`
}
