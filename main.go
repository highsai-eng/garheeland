package main

import (
	"log"
	"os"

	"github.com/garheeland/constant"
	"github.com/garheeland/core"
	"github.com/garheeland/handler"
	"github.com/lxn/walk"
	. "github.com/lxn/walk/declarative"
)

func main() {

	defer log.Println("destroy garheeland...")

	log.Println("starting garheeland...")

	app := core.Application{}

	file := handler.File{App: &app}

	err := MainWindow{
		AssignTo: &app.MainWindow,
		Title:    constant.AppName,
		MinSize:  Size{Width: 300, Height: 200},
		Size:     Size{Width: 1000, Height: 1000},
		Layout:   HBox{MarginsZero: true},
		Children: []Widget{
			TextEdit{
				AssignTo: &app.EditTextArea,
			},
		},
		MenuItems: []MenuItem{
			Menu{
				Text: constant.File,
				Items: []MenuItem{
					Action{Text: constant.NewFile, OnTriggered: func() {}},
					Action{Text: constant.NewWindow, OnTriggered: func() {}},
					Action{Text: constant.Open, Shortcut: Shortcut{Modifiers: walk.ModControl, Key: walk.KeyO}, OnTriggered: file.Open},
					Action{Text: constant.SaveOverride, OnTriggered: func() {}},
					Action{Text: constant.SaveNaming, Shortcut: Shortcut{Modifiers: walk.ModControl, Key: walk.KeyO}, OnTriggered: file.SaveNaming},
					Separator{},
					Action{Text: constant.SaveAndClose, OnTriggered: func() {}},
					Action{Text: constant.Close, Shortcut: Shortcut{Modifiers: walk.ModControl | walk.ModShift, Key: walk.KeyS}, OnTriggered: file.Close},
					Separator{},
					Action{Text: constant.Print, OnTriggered: func() {}},
					Action{Text: constant.PrintPreview, OnTriggered: func() {}},
					Action{Text: constant.PrintPageSetting, OnTriggered: func() {}},
					Separator{},
					Action{Text: constant.FileProperty, OnTriggered: func() {}},
					Action{Text: constant.Browse, OnTriggered: func() {}},
					Separator{},
					Action{Text: constant.RecentlyUsedFiles, OnTriggered: func() {}},
					Action{Text: constant.RecentlyUsedFolders, OnTriggered: func() {}},
				},
			},
			Menu{
				Text: constant.Edit,
			},
			Menu{
				Text: constant.Convert,
			},
			Menu{
				Text: constant.Search,
			},
			Menu{
				Text: constant.Tool,
			},
			Menu{
				Text: constant.Setting,
			},
			Menu{
				Text: constant.Window,
			},
			Menu{
				Text: constant.Help,
			},
		},
	}.Create()

	if err != nil {
		panic(err.Error())
	}

	statusCode := app.Run()
	os.Exit(statusCode)
}
