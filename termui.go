package main

import (
	"github.com/epiclabs-io/winman"
	"github.com/rivo/tview"
)

type TermUI struct{}

func (tui *TermUI) Run() error {
	wm := winman.NewWindowManager()
	wm.NewWindow().
		Show().
		SetRoot(tui.appPanel()).
		SetBorder(false).
		Maximize()

	app := tview.NewApplication()
	err := app.SetRoot(wm, true).
		EnableMouse(true).
		Run()
	if err != nil {
		return err
	}

	return nil
}

func (tui *TermUI) appPanel() tview.Primitive {
	sidebar := tview.NewFlex().
		SetDirection(tview.FlexRow).
		AddItem(tui.sidebarTreePanel(), 0, 5, true).
		AddItem(tui.sidebarButtonPanel(), 0, 2, true).
		AddItem(tui.sidebarAppInfo(), 0, 1, true)

	layout := tview.NewFlex().
		SetDirection(tview.FlexColumn).
		AddItem(sidebar, 0, 1, true).
		AddItem(tui.diffViewerPanel(), 0, 2, true)

	return layout
}

func (tui *TermUI) sidebarTreePanel() tview.Primitive {
	view := tview.NewBox().
		SetTitle("Browse").
		SetTitleAlign(tview.AlignCenter).
		SetBorder(true).
		SetBorderPadding(1, 1, 1, 1)

	return view
}

func (tui *TermUI) sidebarButtonPanel() tview.Primitive {
	view := tview.NewBox().
		SetTitle("Button").
		SetTitleAlign(tview.AlignCenter).
		SetBorder(true).
		SetBorderPadding(1, 1, 1, 1)

	return view
}

func (tui *TermUI) sidebarAppInfo() tview.Primitive {
	view := tview.NewBox().
		SetBorder(true).
		SetBorderPadding(1, 1, 1, 1)

	return view
}

func (tui *TermUI) diffViewerPanel() tview.Primitive {
	view := tview.NewBox().
		SetTitle("Diff Viewer").
		SetTitleAlign(tview.AlignCenter).
		SetBorder(true).
		SetBorderPadding(1, 1, 1, 1)

	return view
}
