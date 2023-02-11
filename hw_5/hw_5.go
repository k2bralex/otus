package hw_5

import (
	"fmt"
	"log"
)

/*
Написать функцию для параллельного выполнения N заданий (т.е. в N параллельных горутинах).
Функция принимает на вход:
- слайс с заданиями []func() error;
- число заданий которые можно выполнять параллельно (N);
- максимальное число ошибок после которого нужно приостановить обработку.
*/

func FuncPool(pool []func() error, n, maxErr int) {
	errCount := 0
	funcChan := make(chan func() error, n)
	errChan := make(chan error, n)

	go func() {
		for _, f := range pool {
			funcChan <- f
			go func() {
				someFunc := <-funcChan
				err := someFunc()
				errChan <- err
			}()
		}
		close(funcChan)
	}()

	for err := range errChan {
		if errCount >= maxErr {
			fmt.Println("error limit")
			return
		}
		errCount++
		log.Println(err)
	}
}

/*func FuncPool(pool []func() error, errCount, maxErr int) {
	var mu sync.Mutex
	for _, f := range pool {
		if maxErr >= errCount {
			break
		}
		f := f
		go func(errCount int) {
			err := f()
			if err != nil {
				mu.Lock()
				errCount++
				mu.Lock()
				return
			}
		}(errCount)
	}
}

func withError() error {
	return fmt.Errorf("some error")
}

func noError() {
	time.Sleep(10 * time.Second)
	fmt.Println("done")
}*/
