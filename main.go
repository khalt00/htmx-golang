package main

import (
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"golang.org/x/exp/slices"
)

type TodoList struct {
	ID          int
	Task        string
	CreatedAt   string
	IsCompleted bool
}
type ErrorMap map[string]string

func (t *TodoList) validate() ErrorMap {
	var errors ErrorMap = make(ErrorMap)

	if t.Task == "" {
		errors["task"] = "Task is required"
	}

	return errors

}

var todoList = []TodoList{
	{
		ID:          1,
		Task:        "Buy groceries",
		CreatedAt:   time.Now().Format("2006-01-02 15:04:05"),
		IsCompleted: false,
	},
	{
		ID:          2,
		Task:        "Finish homework",
		CreatedAt:   time.Now().Format("2006-01-02 15:04:05"),
		IsCompleted: false,
	},
	{
		ID:          3,
		Task:        "Go for a run",
		CreatedAt:   time.Now().Format("2006-01-02 15:04:05"),
		IsCompleted: true,
	},
}

func deleteTodo(id int) error {

	for i, v := range todoList {
		if id == v.ID {
			todoList = slices.Delete(todoList, i, i+1)
		}
	}
	fmt.Print(todoList)
	return nil
}

type EditCreateTodoReq struct {
	Task        string
	CreatedAt   time.Time
	IsCompleted bool
}

type NewTodoPage struct {
}

func createTodo(req EditCreateTodoReq) error {

	newTodo := TodoList{
		ID:          len(todoList) + 1,
		Task:        req.Task,
		CreatedAt:   time.Now().Format("2006-01-02 15:04:05"),
		IsCompleted: req.IsCompleted,
	}

	todoList = append(todoList, newTodo)

	return nil
}

func getTodoByID(id int) (TodoList, error) {
	for i, v := range todoList {
		if v.ID == id {
			return todoList[i], nil
		}
	}
	return TodoList{}, nil
}

func editTodo(id int, req EditCreateTodoReq) ([]TodoList, error) {

	return nil, nil
}

func main() {

	mygin := gin.Default()

	mygin.LoadHTMLGlob("templates/*")

	mygin.GET("/index", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index", gin.H{
			"data":  todoList,
			"title": "Main website",
		})

	})
	mygin.DELETE("/:id", func(ctx *gin.Context) {
		tid := ctx.Param("id")
		id, err := strconv.Atoi(tid)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, "bad id")
			return
		}
		if err := deleteTodo(id); err != nil {
			ctx.JSON(http.StatusInternalServerError, "bad id")
			return
		}

		ctx.String(http.StatusOK, "")
	})
	mygin.GET("/new", func(ctx *gin.Context) {
		ctx.HTML(200, "create-edit-todo", nil)
	})
	mygin.GET("/edit/:id", func(ctx *gin.Context) {
		id := ctx.Param("id")
		tid, _ := strconv.Atoi(id)

		todo, err := getTodoByID(tid)
		if err != nil {
			return
		}

		fmt.Print(todo)
		ctx.HTML(200, "create-edit-todo", gin.H{
			"Todo": todo,
		})
	})
	mygin.POST("/new", func(ctx *gin.Context) {
		isDone, _ := strconv.ParseBool(ctx.Request.FormValue("isCompleted"))
		fmt.Print(isDone)
		req := EditCreateTodoReq{
			Task:        ctx.Request.FormValue("task"),
			IsCompleted: isDone,
		}

		if err := createTodo(req); err != nil {
			ctx.String(500, "")
			return
		}

		ctx.Redirect(303, "/index")
	})
	// mygin.PUT("/:id", func(ctx *gin.Context) {
	// 	tid := ctx.Param("id")
	// 	id, err := strconv.Atoi(tid)
	// 	if err != nil {
	// 		ctx.JSON(http.StatusBadRequest, "bad id")
	// 		return
	// 	}
	// 	if err := editTodo(id, req); err != nil {
	// 		ctx.String(http.StatusInternalServerError, "wrong...")
	// 	}
	// })

	mygin.Run(":9000")
}
