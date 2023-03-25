package test

import (
	"Otus/fcopy/internal/app"
	"github.com/stretchr/testify/assert"
	"testing"
)

var (
	src = "./test.txt"
	trg = "./test_copy.txt"
	lim = int64(0)
	off = int64(0)
)

func TestSourceOpenExistFile(t *testing.T) {
	err := app.Run(&src, &trg, &lim, &off)

	assert.NoError(t, err, "source correct. no error")
}

func TestSourceOpenNonExistFile(t *testing.T) {
	src = "./none_exist.txt"
	err := app.Run(&src, &trg, &lim, &off)

	assert.Error(t, err)
}

func TestTargetOpenNonExistFile(t *testing.T) {
	err := app.Run(&src, &trg, &lim, &off)

	assert.NoError(t, err, "target correct. no error")
}

func TestTargetOpenExistFile(t *testing.T) {
	err := app.Run(&src, &trg, &lim, &off)

	assert.ErrorContains(t, err, "file exist already")
}
