package v1

import (
	"github.com/brian-nunez/bsuite-api/internal/handlers/v1/tasks"
	"github.com/labstack/echo/v4"
)

func RegisterRoutes(e *echo.Echo) {
	v1Group := e.Group("/api/v1")
	v1Group.GET("/health", HealthHandler)

	schedulerGroup := v1Group.Group("/scheduler")
	schedulerGroup.GET("/tasks", tasks.ReadAllTasksHandler)
	schedulerGroup.GET("/tasks/completed", tasks.ReadAllCompletedTasksHandler)
	schedulerGroup.GET("/tasks/pending", tasks.ReadAllPendingTasksHandler)
	schedulerGroup.GET("/tasks/failed", tasks.ReadAllFailedTasksHandler)
	schedulerGroup.GET("/tasks/running", tasks.ReadAllRunningTasksHandler)
	schedulerGroup.POST("/task/m3u8", tasks.CreateM3U8Task)
}
