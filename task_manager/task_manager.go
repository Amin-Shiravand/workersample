package task_manager

import (
	"time"
	"workersample/task_manager/task"
	"workersample/task_manager/worker"
	"workersample/utils"
)

type TaskManager struct {
	Workers []*worker.Worker
}

func NewTaskManager(numWorkers int) *TaskManager {
	workers := make([]*worker.Worker, numWorkers)
	for i := 0; i < numWorkers; i++ {
		workers[i] = worker.NewWorker(i+1, utils.GenerateRandomString(i))
		workers[i].Execute()
	}
	return &TaskManager{Workers: workers}
}

func (tm *TaskManager) DispatchTasks(tasks []task.Task, delay time.Duration) {
	for _, t := range tasks {
		for _, currentWorker := range tm.Workers {
			currentWorker.Tasks <- t
			time.Sleep(delay)
		}
	}
}

func (tm *TaskManager) Close() {
	for _, currentWorker := range tm.Workers {
		close(currentWorker.Tasks)
	}
}
