package task

import (
	"errors"
	"fmt"
	"regexp"
	"strconv"
	"task-cli-go/internal/console"
	"task-cli-go/internal/dto"
	"task-cli-go/internal/logger"
	"task-cli-go/internal/task"
)

type Filter struct {
	service task.Service
	logger  logger.Service
}

func NewFilter(service task.Service, logger logger.Service) *Filter {
	return &Filter{
		service: service,
		logger:  logger,
	}
}

func (c *Filter) Run(args string) {
	flagFilterRegex := regexp.MustCompile(`-s\s+"([^"]+)"`)
	flag := flagFilterRegex.FindString(args)

	if len(flag) == 0 {
		c.logger.LogError(ErrMissingStatusFlag.Error())

		return
	}

	statusRegex := regexp.MustCompile(`"([^"]+)"`)
	status := statusRegex.FindString(args)

	if len(status) == 0 {
		c.logger.LogError(ErrEmptyStatus.Error())
	}

	status, err := extractValue(args)

	if err != nil {
		c.logger.LogError(ErrFailedToExtractStatus.Error())

		return
	}

	var tasks []dto.TaskDTO

	switch status {
	case task.DONE:
		tasks = append(tasks, c.service.FilterByStatus(task.DONE)...)
	case task.InPROGRESS:
		tasks = append(tasks, c.service.FilterByStatus(task.InPROGRESS)...)
	case task.TODO:
		tasks = append(tasks, c.service.FilterByStatus(task.TODO)...)
	default:
		c.logger.LogError(ErrWrongStatusTypePassed.Error())

		return
	}

	for _, t := range tasks {
		c.logger.LogInfo(
			fmt.Sprintf(
				`ID: %s Description: %s Status: %s Updated: %s`,
				strconv.Itoa(*t.ID),
				t.Description,
				t.Status,
				*t.UpdatedAt,
			))
	}
}

func (c *Filter) GetCmd() *console.Console {
	command := console.NewConsoleCommand("filter", "get list item by status", c.Run)

	return command
}

func extractValue(input string) (string, error) {
	var flag, value string
	n, err := fmt.Sscanf(input, "%s %q", &flag, &value)
	if err != nil || n != 2 {
		return "", errors.New("invalid format")
	}
	return value, nil
}
