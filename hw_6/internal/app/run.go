package app

import (
	"fmt"
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

	_, err = os.Stat(*trg)
	if err == nil {
		return fmt.Errorf("file exist already")
	}

	fileTo, err := os.OpenFile(*trg, os.O_CREATE|os.O_WRONLY, 0644)
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

	if *off != 0 {
		_, err = fileFrom.Seek(*off, io.SeekStart)
		if err != nil {
			return err
		}
	}

	for *off < *lim {
		read, err := fileFrom.Read(buf)
		*off += int64(read)
		if err != nil && err == io.EOF {
			return err
		}
		if read == 0 {
			break
		}
		if _, err = fileTo.Write(buf[:read]); err != nil {
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
