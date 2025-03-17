package main

import (
	"task-cli-go/cmd/Cli"
	cmdTask "task-cli-go/cmd/Cli/task"
	"task-cli-go/internal/logger"
	"task-cli-go/internal/storage"
	"task-cli-go/internal/task"
)

func NewApp() (*Application, error) {
	log := logger.NewLogger()
	s := storage.CreateNewStorage(log)
	t := task.CreateNewTask(s, log)
	cli := Cli.NewCLi(cmdTask.NewAdd(t, log))
	cli.SetAvailableCommands(
		cmdTask.NewAdd(t, log),
		cmdTask.NewDelete(t, log),
		cmdTask.NewDone(t, log),
		cmdTask.NewList(t, log),
		cmdTask.NewFilter(t, log),
		cmdTask.NewProgress(t, log),
		cmdTask.NewUpdate(t, log),
	)

	application := NewApplication(log, *cli)

	return application, nil
}
