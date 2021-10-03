package controller

import (
	"fmt"

	"github.com/SpicyChickenFLY/never-todo-cmd/constant"
	"github.com/SpicyChickenFLY/never-todo-cmd/model"
)

// ShowHelp for user
func ShowHelp() error {
	return nil
}

// StartUI interface
func StartUI() error {

	return nil
}

// ShowSummary like the total amount of todo/tag/schedule
func ShowSummary() error {
	// 展示logo，用法，当前待办数和标签数
	fmt.Println(constant.ColorfulLogo)
	fmt.Println(constant.Descirption)
	fmt.Println(constant.Separator)
	var todoTotal, doneTotal, deletedTotal, tagTotal int
	for _, task := range model.DB.Data.Tasks {
		switch task.Status {
		case model.ProjectTodo:
			todoTotal++
		case model.ProjectDone:
			doneTotal++
		case model.ProjectDeleted:
			deletedTotal++
		}
	}
	for _, tag := range model.DB.Data.Tags {
		if !tag.Deleted {
			tagTotal++
		}
	}
	fmt.Printf("todo: %d, done: %d, tag: %d\n", todoTotal, doneTotal, tagTotal)

	return nil
}
