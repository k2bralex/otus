package main

import (
	"flag"
	"fmt"
	"io"
	"os"
)

//	Домашнее задание
//	Копирование файлов Цель: Реализовать утилиту копирования файлов
//	Утилита должна принимать следующие аргументы:
//	* файл источник (From)
//	* файл копия (To)
//	* Отступ в источнике (Offset), по умолчанию - 0
//	* Количество копируемых байт (Limit), по умолчанию - весь файл из From
//	Выводить в консоль прогресс копирования в %, например с помощью github.com/cheggaaa/pb
//	Программа может НЕ обрабатывать файлы, у которых не известна длинна (например /dev/urandom).
//	Завести в репозитории отдельный пакет (модуль) для этого ДЗ
//	Реализовать функцию вида Copy(from string, to string, limit int, offset int) error
//	Написать unit-тесты на функцию Copy Реализовать функцию main, анализирующую параметры
//	командной строки и вызывающую Copy Проверить установку и работу утилиты руками
//	Критерии оценки:
//	Функция должна проходить все тесты
//	Все необходимые для тестов файлы должны создаваться в самом тесте
//	Код должен проходить проверки go vet и golint
//	Должна быть возможность скачать и установить пакет с помощью go get / go test / go install

var (
	source = flag.String("from", "", "file path copy FROM")
	target = flag.String("to", "", "file path copy TO")
	limit  = flag.Int64("l", 0, "bytes to copy")
	offset = flag.Int64("o", 0, "offset to start copy")
)

func main() {
	flag.Parse()

	if err := Run(source, target, limit, offset); err != nil {
		fmt.Println(err.Error())
	}
}

func Run(src, trg *string, lim, off *int64) error {
	fileFrom, err := os.Open(*src)
	if err != nil {
		return err
	}
	defer fileFrom.Close()

	if *lim == 0 {
		if err = limitHandler(src, lim); err != nil {
			return err
		}
	}

	buf := make([]byte, *lim)

	for *off < *lim {
		read, err := fileFrom.Read(buf[*off:])
		*off += int64(read)
		if err == io.EOF {
			break
		}
	}

	fileTo, err := os.OpenFile(*target, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return err
	}
	defer fileTo.Close()

	fileTo.Write(buf)

	/*go func() {
		bar := pb.StartNew(100)
		defer bar.Finish()

	}()*/

	return nil
}

func limitHandler(path *string, lim *int64) error {
	info, err := os.Stat(*path)
	if err != nil {
		return err
	}
	*lim = info.Size()
	return nil
}
