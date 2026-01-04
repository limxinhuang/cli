package todo

import (
	"fmt"
	"strconv"
	"time"

	utils "github.com/limxinhuang/cli/pkg/utlis"
)

const timeFormat = "2006-01-02 15:04:05"

func Add(title string) {
	todos, _ := loadTodos()

	newTodo := Todo{
		ID:        GetMaxId(),
		Title:     title,
		Completed: false,
		CreatedAt: time.Now(),
	}

	todos = append(todos, newTodo)
	saveTodos(todos)
	fmt.Println("已添加任务：", title)
}

func List() {
	todos, _ := loadTodos()

	table := utils.NewTableWriterDefault([]string{"ID", "任务", "状态", "创建时间", "完成时间"})

	for _, t := range todos {
		status := "\033[32mX\033[0m"
		id := strconv.Itoa(t.ID)
		completedAt := "-"
		if t.Completed {
			status = "\033[31m%s\033[0m"
			completedAt = t.CompletedAt.Format(timeFormat)
		}

		row := []string{
			id,
			t.Title,
			status,
			t.CreatedAt.Format(timeFormat),
			completedAt,
		}

		table.Append(row)
	}

	table.Render()
}

func Completed(id int) {
	todos, _ := loadTodos()
	found := false
	var todo Todo

	for i, t := range todos {
		if t.ID == id {
			todos[i].Completed = true
			todos[i].CompletedAt = time.Now()
			found = true
			todo = todos[i]
			break
		}
	}

	if found {
		saveTodos(todos)
		fmt.Printf("任务 %s 已完成\n", todo.Title)
	} else {
		fmt.Printf("任务ID %d 不存在\n", id)
	}
}

func DeleteTodo(id int) {
	todos, _ := loadTodos()
	var newTodos []Todo

	for _, t := range todos {
		if t.ID != id {
			newTodos = append(newTodos, t)
		}
	}

	saveTodos(newTodos)
	fmt.Println("任务已删除")
}

func UpdateTitle(id int, title string) {
	todos, _ := loadTodos()
	var oldTitle string

	for i, t := range todos {
		if t.ID == id {
			oldTitle = todos[i].Title
			todos[i].Title = title
			break
		}
	}

	if oldTitle == "" {
		fmt.Printf("任务ID %d 不存在\n", id)
		return
	}

	saveTodos(todos)
	fmt.Printf("任务ID %d 的标题修改成功 %s -> %s\n", id, oldTitle, title)
}
