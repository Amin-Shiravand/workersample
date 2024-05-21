package main

import (
	"fmt"
	"time"
	"workersample/task_manager"
	"workersample/task_manager/task"
	"workersample/utils"
)

func main() {
	numWorkers := 2
	numTasks := 5
	taskDelay := time.Millisecond * 100

	taskManager := task_manager.NewTaskManager(numWorkers)
	tasks := make([]task.Task, numTasks)
	for i := 0; i < numTasks; i++ {
		tasks[i] = task.Task{Message: utils.GenerateRandomString(i)}
	}

	taskManager.DispatchTasks(tasks, taskDelay)
	taskManager.Close()
	fmt.Println("All tasks completed")
}
