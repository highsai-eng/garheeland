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
				//Items: []MenuItem{
				//	Action{
				//		//AssignTo: &openAction,
				//		Text: "&Open",
				//		//Image:       "../img/open.png",
				//		Enabled:  Bind("enabledCB.Checked"),
				//		Visible:  Bind("!openHiddenCB.Checked"),
				//		Shortcut: Shortcut{walk.ModControl, walk.KeyO},
				//		//OnTriggered: mw.openAction_Triggered,
				//	},
				//	Menu{
				//		//AssignTo: &recentMenu,
				//		Text: "Recent",
				//	},
				//	Separator{},
				//	Action{
				//		Text: "E&xit",
				//		//OnTriggered: func() { mw.Close() },
				//	},
				//},
				Items: []MenuItem{
					Action{Text: constant.NewFile, OnTriggered: func() {}},
					Action{Text: constant.NewWindow, OnTriggered: func() {}},
					Action{Text: constant.Open, Shortcut: Shortcut{Modifiers: walk.ModControl, Key: walk.KeyO}, OnTriggered: file.Open},
					Action{Text: constant.SaveOverride, OnTriggered: func() {}},
					Action{Text: constant.SaveOverrideAll, OnTriggered: func() {}},
					Action{Text: constant.SaveNaming, OnTriggered: func() {}},
					Separator{},
					Action{Text: constant.SaveAndClose, OnTriggered: func() {}},
					Action{Text: constant.Close, OnTriggered: func() {}},
					Action{Text: constant.CloseAndAnonymous, OnTriggered: func() {}},
					Action{Text: constant.CloseAndOpen, OnTriggered: func() {}},
					Action{Text: constant.ReOpen, OnTriggered: func() {}},
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
					Separator{},
					Action{Text: constant.Exit, OnTriggered: file.Exit},
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
