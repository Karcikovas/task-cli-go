package main

import (
	"task-cli-go/cmd/cli"
	cmdTask "task-cli-go/cmd/cli/task"
	"task-cli-go/internal/logger"
	"task-cli-go/internal/storage"
	"task-cli-go/internal/task"
)

func NewApp() *Application {
	log := logger.NewLogger()
	s := storage.CreateNewStorage(log)
	t := task.CreateNewTask(s, log)
	root := cli.NewCLi(
		cmdTask.NewAdd(t, log),
		log,
	)
	root.SetAvailableCommands(
		cmdTask.NewAdd(t, log),
		cmdTask.NewDelete(t, log),
		cmdTask.NewDone(t, log),
		cmdTask.NewList(t, log),
		cmdTask.NewFilter(t, log),
		cmdTask.NewProgress(t, log),
		cmdTask.NewUpdate(t, log),
	)

	application := NewApplication(log, *root)

	return application
}
