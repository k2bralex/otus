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
	errChan := make(chan error, n)
	funcChan := make(chan func() error, n)

	go func(c *int) {
		for _, f := range pool {
			if *c >= maxErr {
				return
			}
			go func(c *int) {
				if *c >= maxErr {
					close(errChan)
					return
				}
				fn := <-funcChan
				err := fn()
				errChan <- err
			}(&errCount)
			funcChan <- f
		}
		close(funcChan)
	}(&errCount)

	for err := range errChan {
		if errCount >= maxErr {
			fmt.Println("error limit")
			return
		}
		errCount++
		log.Println(err)
	}
}
