package workerpool

import "fmt"

type Task struct {
	Err  error
	Data interface{}
	f    func(interface{}) error
}

func NewTask(f func(interface{}) error, Data interface{}) *Task {
	return &Task{f: f, Data: Data}
}

func Process(workerID int, task *Task) {
	fmt.Printf("Worker %d processes task %v\n", workerID, task.Data)
	task.Err = task.f(task.Data)
}
