package handler

import (
	"io/ioutil"
	"log"
	"os"
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

	bytes, err := ioutil.ReadFile(dialog.FilePath)
	if err != nil {
		log.Fatalln("file open error.", err)
		return
	}

	str := strings.ReplaceAll(string(bytes), "\n", "\r\n")

	if err := f.App.EditTextArea.SetText(str); err != nil {
		log.Println("edit text error.", err)
		return
	}

	if err := f.App.MainWindow.SetTitle(constant.AppName + " - " + dialog.FilePath); err != nil {
		log.Fatalln("set title error.", err)
	}
}

func (f *File) SaveNaming() {

	dialog := walk.FileDialog{
		Title:    "名前を付けて保存",
		FilePath: f.App.Path,
	}

	ok, err := dialog.ShowSave(f.App)

	if err != nil {
		log.Fatalln("file dialog error.", err)
		return
	}

	if !ok {
		return
	}

	f.App.Path = dialog.FilePath

	if err := ioutil.WriteFile(dialog.FilePath, []byte(f.App.EditTextArea.Text()), os.ModePerm); err != nil {
		log.Fatalln("file save error.", err)
	}

	if err := f.App.MainWindow.SetTitle(constant.AppName + " - " + dialog.FilePath); err != nil {
		log.Fatalln("set title error.", err)
	}
}

func (f *File) Close() {
	if err := f.App.Close(); err != nil {
		log.Fatalln(err.Error())
	}
}
