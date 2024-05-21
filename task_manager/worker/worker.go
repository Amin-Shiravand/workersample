package worker

import (
	"fmt"
	"time"
	"workersample/task_manager/task"
)

type Worker struct {
	ID      int
	Message string
	Tasks   chan task.Task
}

func NewWorker(id int, message string) *Worker {
	return &Worker{
		ID:      id,
		Message: message,
		Tasks:   make(chan task.Task),
	}
}

func (w *Worker) Execute() {
	go func() {
		for currentTask := range w.Tasks {
			w.Process(currentTask)
		}
	}()
}

func (w *Worker) Process(task task.Task) {
	fmt.Printf("Worker %d is processing task %d\n", w.ID, task.Message)
	time.Sleep(time.Millisecond * 100)
}
