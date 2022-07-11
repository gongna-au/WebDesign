package todo

import (
	"github.com/WebDesign/database"
	"github.com/WebDesign/model"
	"github.com/gin-gonic/gin"

	"net/http"
	"strconv"
)

func CreateTodo(c *gin.Context) {
	completed, _ := strconv.Atoi(c.PostForm("completed"))
	todo := model.TodoModel{Title: c.PostForm("title"), Completed: completed}
	database.GetDBInstance().Save(&todo)
	c.JSON(http.StatusCreated, gin.H{"status": http.StatusCreated, "message": "Todo item created successfully!", "resourceId": todo.ID})
}

// fetchAllTodo 返回所有的 todo 数据
func FetchAllTodo(c *gin.Context) {
	var todos []model.TodoModel
	var _todos []model.TransformedTodo
	//
	database.GetDBInstance().Find(&todos)

	if len(todos) <= 0 {
		c.JSON(http.StatusNotFound, gin.H{"status": http.StatusNotFound, "message": "No todo found!"})
		return
	}

	//转化 todos 数据，用来格式化
	for _, item := range todos {
		completed := false
		if item.Completed == 1 {
			completed = true
		} else {
			completed = false
		}
		_todos = append(_todos, model.TransformedTodo{ID: item.ID, Title: item.Title, Completed: completed})
	}
	c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "data": _todos})
}

// fetchSingleTodo方法返回一条 todo 数据
func FetchSingleTodo(c *gin.Context) {
	var todo model.TodoModel
	todoID := c.Param("id")

	database.GetDBInstance().First(&todo, todoID)

	if todo.ID == 0 {
		c.JSON(http.StatusNotFound, gin.H{"status": http.StatusNotFound, "message": "No todo found!"})
		return
	}

	completed := false
	if todo.Completed == 1 {
		completed = true
	} else {
		completed = false
	}

	_todo := model.TransformedTodo{ID: todo.ID, Title: todo.Title, Completed: completed}
	c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "data": _todo})
}

// updateTodo 方法 更新 todo 数据
func UpdateTodo(c *gin.Context) {
	var todo model.TodoModel
	todoID := c.Param("id")

	database.GetDBInstance().First(&todo, todoID)

	if todo.ID == 0 {
		c.JSON(http.StatusNotFound, gin.H{"status": http.StatusNotFound, "message": "No todo found!"})
		return
	}

	database.GetDBInstance().Model(&todo).Update("title", c.PostForm("title"))
	completed, _ := strconv.Atoi(c.PostForm("completed"))
	database.GetDBInstance().Model(&todo).Update("completed", completed)
	c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "message": "Todo updated successfully!"})
}

// deleteTodo 方法依据 id 删除一条todo 数据
func DeleteTodo(c *gin.Context) {
	var todo model.TodoModel
	todoID := c.Param("id")

	database.GetDBInstance().First(&todo, todoID)

	if todo.ID == 0 {
		c.JSON(http.StatusNotFound, gin.H{"status": http.StatusNotFound, "message": "No todo found!"})
		return
	}

	database.GetDBInstance().Delete(&todo)
	c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "message": "Todo deleted successfully!"})
}
