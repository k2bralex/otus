package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

/*
Написать функцию для параллельного выполнения N заданий (т.е. в N параллельных горутинах).
Функция принимает на вход:
- слайс с заданиями []func() error;
- число заданий которые можно выполнять параллельно (N);
- максимальное число ошибок после которого нужно приостановить обработку.
*/

type Task struct {
	Message string
}

func (t *Task) ErrorResponse() {
	//time.Sleep(200 * time.Millisecond)
	t.Message = "Stopped"
}

func (t *Task) NoErrorResponse() {
	//time.Sleep(300 * time.Millisecond)
	t.Message = "Done"
}

func main() {
	taskPoll := []Task{
		{}, {}, {}, {}, {}, {}, {}, {}, {}, {}, {},
		{}, {}, {}, {}, {}, {}, {}, {}, {}, {}, {},
		{}, {}, {}, {}, {}, {}, {}, {}, {}, {}, {},
		{}, {}, {}, {}, {}, {}, {}, {}, {}, {}, {},
		{}, {}, {}, {}, {}, {}, {}, {}, {}, {}, {},
		{}, {}, {}, {}, {}, {}, {}, {}, {}, {}, {},
		{}, {}, {}, {}, {}, {}, {}, {}, {}, {}, {},
		{}, {}, {}, {}, {}, {}, {}, {}, {}, {}, {},
		{}, {}, {}, {}, {}, {}, {}, {}, {}, {}, {},
	}

	TaskRun(&taskPoll, 5, 3)
}

func TaskRun(p *[]Task, ng, me int) {
	in := make(chan *Task)
	out := make(chan *Task)

	var (
		ctx, cancel = context.WithCancel(context.Background())
		wg          = &sync.WaitGroup{}
		mu          = &sync.Mutex{}
		errCount    = 0
	)
	defer cancel()

	for i := 0; i <= ng; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			worker(ctx, in, out, &errCount, mu)
		}()
	}

	go func() {
		defer close(in)
		for _, task := range *p {
			if errCount == me {
				cancel()
			}
			in <- &task
		}
	}()

	go func() {
		wg.Wait()
		close(out)
	}()

	for task := range out {
		fmt.Println(task.Message)
	}

	fmt.Println(errCount)

}

func worker(ctx context.Context, in, out chan *Task, err *int, m *sync.Mutex) {
	for {
		select {
		case <-ctx.Done():
			return
		case task, ok := <-in:
			if !ok {
				return
			}
			time.Sleep(100 * time.Millisecond)
			out <- responseHandler(task, err, m)
		}
	}
}

func responseHandler(task *Task, err *int, m *sync.Mutex) *Task {
	if time.Now().Unix()%2 == 0 {
		task.NoErrorResponse()
	} else {
		task.ErrorResponse()
		m.Lock()
		*err++
		m.Unlock()
	}
	return task
}
