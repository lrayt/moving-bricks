package handler

import (
	"os"
	"testing"
)

func TestTaskHandler_Run(t *testing.T) {
	task := TaskHandler{source: "d://tmp", target: "d://abc"}
	task.exist()
	err := task.list()
	t.Log(err)
	task.mkdir()
}

func TestReadDir(t *testing.T) {
	files, err := os.ReadDir("d://tmp")
	t.Log(err)
	for _, o := range files {
		t.Log(o.IsDir(), o.Name())
	}
}
