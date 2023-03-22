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
	Message  string
	Response error
}

func (t *Task) ErrorResponse() {
	//time.Sleep(200 * time.Millisecond)
	t.Message = "Stopped"
	t.Response = fmt.Errorf("some error")

}

func (t *Task) NoErrorResponse() {
	//time.Sleep(300 * time.Millisecond)
	t.Message = "Done"
	t.Response = nil
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

	TaskRun(&taskPoll, 5, 10)
}

func TaskRun(p *[]Task, gNum, maxErr int) {
	in := make(chan *Task, gNum)
	out := make(chan *Task, gNum)

	var (
		ctx, cancel = context.WithCancel(context.Background())
		wg          = &sync.WaitGroup{}
		mu          = &sync.Mutex{}
		errCount    = 0
	)

	defer cancel()

	for i := 0; i < gNum; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			worker(ctx, in, out, &errCount, mu)
		}()
	}

	go func(err *int) {
		defer close(in)
		for _, task := range *p {
			time.Sleep(100 * time.Millisecond)
			if *err >= maxErr {
				cancel()
			}
			in <- &task
		}
	}(&errCount)

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
			out <- responseHandler(task, err, m)
		}
	}
}

func responseHandler(task *Task, err *int, m *sync.Mutex) *Task {
	time.Sleep(100 * time.Millisecond)
	if time.Now().Unix()%2 == 0 {
		task.NoErrorResponse()
	} else {
		m.Lock()
		*err++
		m.Unlock()
		task.ErrorResponse()
	}
	return task
}
