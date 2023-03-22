package app

import (
	"github.com/cheggaaa/pb"
	"io"
	"os"
)

func Run(src, trg *string, lim, off *int64) error {
	fileFrom, err := os.Open(*src)
	if err != nil {
		return err
	}
	defer fileFrom.Close()

	fileTo, err := os.OpenFile(*trg, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return err
	}
	defer fileTo.Close()

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
		if _, err = fileTo.Write(buf); err != nil {
			return err
		}
	}

	go func() {
		bar := pb.StartNew(100)
		defer bar.Finish()

	}()

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
