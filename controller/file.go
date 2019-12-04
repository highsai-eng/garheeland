package controller

import (
	"log"

	"github.com/garheeland/core"
)

type File struct {
	App *core.Application
}

func (f *File) Exit() {
	if err := f.App.Close(); err != nil {
		log.Fatalln(err.Error())
	}
}
