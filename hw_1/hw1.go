//Написать программу печатающую текущее время / точное время с использованием библиотеки NTP.
//Программа должна корректно обрабатывать ошибки библиотеки: распечатывать их в STDERR и
//возвращать ненулевой код выхода.

package hw_1

import (
	"fmt"
	"github.com/beevik/ntp"
	"time"
)

func MyTime() string {
	timeLocal := time.Now()
	timeWeb, err := ntp.Time("0.beevik-ntp.pool.ntp.org")
	if err != nil {
		return err.Error()
	}
	return fmt.Sprintf("Local time: %v\nWebserver time: %v\n\n", timeLocal, timeWeb)
}
