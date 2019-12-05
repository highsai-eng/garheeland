package main

import (
	"log"
	"os"

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

	var in *walk.TextEdit

	err := MainWindow{
		AssignTo: &app.MainWindow,
		Title:    "Garhee Land",
		MinSize:  Size{Width: 300, Height: 200},
		Size:     Size{Width: 1000, Height: 1000},
		Layout:   HBox{MarginsZero: true},
		Children: []Widget{
			TextEdit{
				AssignTo: &in,
			},
		},
		MenuItems: []MenuItem{
			Menu{
				Text: File,
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
					Action{Text: NewFile, OnTriggered: func() {}},
					Action{Text: NewWindow, OnTriggered: func() {}},
					Action{Text: Open, OnTriggered: func() {}},
					Action{Text: SaveOverride, OnTriggered: func() {}},
					Action{Text: SaveOverrideAll, OnTriggered: func() {}},
					Action{Text: SaveNaming, OnTriggered: func() {}},
					Separator{},
					Action{Text: SaveAndClose, OnTriggered: func() {}},
					Action{Text: Close, OnTriggered: func() {}},
					Action{Text: CloseAndAnonymous, OnTriggered: func() {}},
					Action{Text: CloseAndOpen, OnTriggered: func() {}},
					Action{Text: ReOpen, OnTriggered: func() {}},
					Separator{},
					Action{Text: Print, OnTriggered: func() {}},
					Action{Text: PrintPreview, OnTriggered: func() {}},
					Action{Text: PrintPageSetting, OnTriggered: func() {}},
					Separator{},
					Action{Text: FileProperty, OnTriggered: func() {}},
					Action{Text: Browse, OnTriggered: func() {}},
					Separator{},
					Action{Text: RecentlyUsedFiles, OnTriggered: func() {}},
					Action{Text: RecentlyUsedFolders, OnTriggered: func() {}},
					Separator{},
					Action{Text: Exit, OnTriggered: file.Exit},
				},
			},
			Menu{
				Text: Edit,
			},
			Menu{
				Text: Convert,
			},
			Menu{
				Text: Search,
			},
			Menu{
				Text: Tool,
			},
			Menu{
				Text: Setting,
			},
			Menu{
				Text: Window,
			},
			Menu{
				Text: Help,
			},
		},
	}.Create()

	if err != nil {
		panic(err.Error())
	}

	statusCode := app.Run()
	os.Exit(statusCode)
}
