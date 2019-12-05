package handler

import (
	"io/ioutil"
	"log"
	"strings"

	"github.com/garheeland/constant"

	"github.com/garheeland/core"
	"github.com/lxn/walk"
)

type File struct {
	App *core.Application
}

func (f *File) Open() {

	dialog := walk.FileDialog{
		Title:    "ファイルを開く",
		FilePath: f.App.Path,
	}

	ok, err := dialog.ShowOpen(f.App)

	if err != nil {
		log.Fatalln("file dialog error.", err)
		return
	}

	if !ok {
		return
	}

	f.App.Path = dialog.FilePath

	if err := f.App.MainWindow.SetTitle(constant.AppName + " - " + dialog.FilePath); err != nil {
		log.Fatalln("set title error.", err)
		return
	}

	bytes, err := ioutil.ReadFile(dialog.FilePath)
	if err != nil {
		log.Fatalln("file open error.", err)
		return
	}

	str := strings.ReplaceAll(string(bytes), "\n", "\r\n")

	if err := f.App.EditTextArea.SetText(str); err != nil {
		log.Println("edit text error.", err)
	}
}

func (f *File) Exit() {
	if err := f.App.Close(); err != nil {
		log.Fatalln(err.Error())
	}
}
