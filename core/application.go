package core

import "github.com/lxn/walk"

type Application struct {
	*walk.MainWindow
	EditTextArea *walk.TextEdit
	Path         string
}
