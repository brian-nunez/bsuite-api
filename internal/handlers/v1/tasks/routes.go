package tasks

import (
	"fmt"
	"path/filepath"
	"strconv"
	"strings"

	"github.com/brian-nunez/bsuite-api/internal/handlers/errors"
	"github.com/brian-nunez/bsuite-api/internal/utils"
	worker "github.com/brian-nunez/task-orchestration"
	"github.com/labstack/echo/v4"
)

var pool *worker.WorkerPool

func init() {
	pool = &worker.WorkerPool{
		Concurrency:  10,
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
	URL    string `json:"url"`
	Output string `json:"output"`
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

type GetTaskBtProcessIdResponse struct {
	Task *worker.TaskInfo `json:"task"`
}

func GetTaskByProcessId(c echo.Context) error {
	processId := c.QueryParam("processId")
	if processId == "" {
		response := errors.
			InvalidRequest().
			WithMessage("processId is required").
			Build()
		return c.JSON(response.HTTPStatusCode, response)
	}

	task, err := pool.GetTaskByProcessId(worker.GetTaskByProcessIdParams{
		ProcessId: processId,
	})
	if err != nil {
		response := errors.InternalServerError().WithMessage(err.Error()).Build()
		return c.JSON(response.HTTPStatusCode, response)
	}

	return c.JSON(200, &GetTaskBtProcessIdResponse{
		Task: task,
	})
}

type ReadProcessLogResponse struct {
	Data string `json:"data"`
}

func ReadProcessLogAsJSON(c echo.Context) error {
	readLog := ReadProcessLog(func(code int, data *ReadProcessLogResponse) error {
		return c.JSON(code, data)
	})

	return readLog(c)
}
func ReadProcessLogAsRaw(c echo.Context) error {
	readLog := ReadProcessLog(func(code int, data *ReadProcessLogResponse) error {
		return c.String(code, data.Data)
	})

	return readLog(c)
}

func ReadProcessLog(responseFunction func(code int, i *ReadProcessLogResponse) error) func(c echo.Context) error {
	return func(c echo.Context) error {
		processID := c.Param("processId")
		if processID == "" {
			response := errors.InvalidRequest().Build()
			return c.JSON(response.HTTPStatusCode, response)
		}

		logPath := filepath.Join("logs", fmt.Sprintf("%s.log", processID))

		offsetParam := c.QueryParam("offset")
		var offset int64
		if offsetParam != "" {
			parsed, err := strconv.ParseInt(offsetParam, 10, 64)
			if err == nil {
				offset = parsed
			}
		}

		byteParam := c.QueryParam("bytes")
		var byteLength int64
		if byteParam != "" {
			parsed, err := strconv.ParseInt(byteParam, 10, 64)
			if err == nil {
				byteLength = parsed
			}
		} else {
			byteLength = 4096
		}

		seekingData, err := utils.ReadFileBySeeking(utils.ReadFileBySeekingParams{
			Offset:   offset,
			Bytes:    byteLength,
			FilePath: logPath,
		})
		if err != nil {
			response := errors.InternalServerError().Build()
			return c.JSON(response.HTTPStatusCode, response)
		}

		return responseFunction(200, &ReadProcessLogResponse{
			Data: seekingData.Data,
		})
	}
}
