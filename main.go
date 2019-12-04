package main

import (
	"log"
	"os"

	"github.com/garheeland/controller"
	"github.com/garheeland/core"
	"github.com/lxn/walk"
	. "github.com/lxn/walk/declarative"
)

func main() {

	defer log.Println("destroy garheeland...")

	log.Println("starting garheeland...")

	app := core.Application{}

	file := controller.File{App: &app}

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
				Text: "&ファイル",
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
					Action{
						Text:        "&終了",
						OnTriggered: file.Exit,
					},
				},
			},
			Menu{
				Text: "&編集",
			},
			Menu{
				Text: "&変換",
			},
			Menu{
				Text: "&検索",
			},
			Menu{
				Text: "&ツール",
			},
			Menu{
				Text: "&設定",
			},
			Menu{
				Text: "&ウィンドウ",
			},
			Menu{
				Text: "&ヘルプ",
			},
		},
	}.Create()

	if err != nil {
		panic(err.Error())
	}

	statusCode := app.Run()
	os.Exit(statusCode)
}
