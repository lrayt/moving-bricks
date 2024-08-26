package handler

import (
	"io/fs"
	"log"
	"os"
	"path/filepath"
)

type TaskHandler struct {
	source     string
	target     string
	dirList    []string
	fileList   []string
	sourceInfo fs.FileInfo
}

func (h TaskHandler) Run() {

}

func (h *TaskHandler) exist() error {
	if info, err := os.Stat(h.source); err != nil {
		return err
	} else {
		h.sourceInfo = info
	}
	return nil
}

func (h *TaskHandler) list() error {
	h.dirList = make([]string, 0)
	h.fileList = make([]string, 0)

	sourceInfo, err := os.Stat(h.source)
	if err != nil {
		return err
	}
	return filepath.Walk(h.source, func(path string, info fs.FileInfo, err error) error {
		if os.SameFile(sourceInfo, info) {
			return err
		}
		if info.IsDir() {
			if h.isLeafDir(path) {
				h.dirList = append(h.dirList, path)
			}
		} else {
			h.fileList = append(h.fileList, path)
		}
		return err
	})
}

func (h TaskHandler) mkdir() error {
	//os.MkdirAll()
	for _, o := range h.dirList {
		if diff, err := filepath.Rel(h.source, o); err != nil {
			return err
		} else {
			log.Println(filepath.Join(h.target, diff))
		}
	}
	return nil
}

func (h TaskHandler) read() {

}

func (h TaskHandler) write() {

}

func (h TaskHandler) send() {

}

func (h TaskHandler) isLeafDir(dir string) bool {
	dirArr, err := os.ReadDir(dir)
	if err != nil {
		return false
	}
	for _, o := range dirArr {
		if o.IsDir() {
			return false
		}
	}
	return true
}
