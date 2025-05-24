package tasks

import (
	"strings"

	"github.com/brian-nunez/bsuite-api/internal/handlers/errors"
	worker "github.com/brian-nunez/task-orchestration"
	"github.com/labstack/echo/v4"
)

var pool *worker.WorkerPool

func init() {
	pool = &worker.WorkerPool{
		Concurreny:   10,
		LogPath:      "logs",
		DatabasePath: "tasks.db",
	}
	pool.Start()
}

func ReadAllTasksHandler(c echo.Context) error {
	tasks, err := pool.GetAllTasks()
	if err != nil {
		response := errors.InternalServerError().Build()
		return c.JSON(response.HTTPStatusCode, response)
	}

	return c.JSON(200, tasks)
}

func ReadAllCompletedTasksHandler(c echo.Context) error {
	tasks, err := pool.GetCompletedTasks()
	if err != nil {
		response := errors.InternalServerError().Build()
		return c.JSON(response.HTTPStatusCode, response)
	}

	return c.JSON(200, tasks)
}

func ReadAllPendingTasksHandler(c echo.Context) error {
	tasks, err := pool.GetPendingTasks()
	if err != nil {
		response := errors.InternalServerError().Build()
		return c.JSON(response.HTTPStatusCode, response)
	}

	return c.JSON(200, tasks)
}

func ReadAllRunningTasksHandler(c echo.Context) error {
	tasks, err := pool.GetRunningTasks()
	if err != nil {
		response := errors.InternalServerError().Build()
		return c.JSON(response.HTTPStatusCode, response)
	}

	return c.JSON(200, tasks)
}

func ReadAllFailedTasksHandler(c echo.Context) error {
	tasks, err := pool.GetFailedTasks()
	if err != nil {
		response := errors.InternalServerError().Build()
		return c.JSON(response.HTTPStatusCode, response)
	}

	return c.JSON(200, tasks)
}

type CreateM3U8TaskBody struct {
	URL    string `json:"url,required"`
	Output string `json:"output,required"`
}

func CreateM3U8Task(c echo.Context) error {
	var body CreateM3U8TaskBody

	err := c.Bind(&body)
	if err != nil {
		response := errors.InvalidRequest().Build()
		return c.JSON(response.HTTPStatusCode, response)
	}

	if strings.TrimSpace(body.URL) == "" {
		response := errors.
			InvalidRequest().
			WithMessage("url is required").
			Build()
		return c.JSON(response.HTTPStatusCode, response)
	}

	if strings.TrimSpace(body.Output) == "" {
		response := errors.
			InvalidRequest().
			WithMessage("output is required").
			Build()
		return c.JSON(response.HTTPStatusCode, response)
	}

	task, err := pool.AddTask(&M3U8Task{
		URL:    body.URL,
		Output: body.Output,
	})
	if err != nil {
		response := errors.InternalServerError().Build()
		return c.JSON(response.HTTPStatusCode, response)
	}

	return c.JSON(200, task)
}
