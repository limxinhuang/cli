package todo

import (
	"encoding/json"
	"os"
	"path/filepath"
)

// 数据文件
const dbFile = ".todo/todos.json"

func getDbFile() string {
	home, _ := os.UserHomeDir()
	return filepath.Join(home, dbFile)
}

func GetMaxId() int {
	todos, _ := loadTodos()
	var maxId = 0

	for _, t := range todos {
		if t.ID > maxId {
			maxId = t.ID
		}
	}

	return maxId + 1
}

// 读取数据
func loadTodos() ([]Todo, error) {
	if _, err := os.Stat(getDbFile()); os.IsNotExist(err) {
		return []Todo{}, nil
	}

	data, err := os.ReadFile(getDbFile())
	if err != nil {
		return nil, err
	}

	var todos []Todo
	err = json.Unmarshal(data, &todos)
	return todos, err
}

// 保存数据
func saveTodos(todos []Todo) error {
	data, err := json.MarshalIndent(todos, "", "  ")
	if err != nil {
		return err
	}
	// 0644 文件权限 rw-r--r--
	return os.WriteFile(getDbFile(), data, 0644)
}
