package main

import (
	"fmt"
	"time"

	"github.com/andlabs/ui"
	_ "github.com/andlabs/ui/winmanifest"
)

const (
	windowHeight = 600
	windowWidth  = 1000
)

var (
	winMain      *ui.Window
	btnLogin     *ui.Button
	lblStatus    *ui.Label
	entrySummary *ui.MultilineEntry

	entryLog *ui.MultilineEntry

	Version = "v0.1.0"
)

func updateLogger() {
	n := 0
	for {
		if entryLog == nil {
			time.Sleep(1 * time.Second)
		}
		n++
		text := entryLog.Text()
		text += fmt.Sprintf("counter  = %d \n", n)
		ui.QueueMain(func() {
			entryLog.SetText(text)
		})
		time.Sleep(1 * time.Second)
	}
}

func setupUI() {
	btnLogin = ui.NewButton("Github Login")
	lblStatus = ui.NewLabel("status")
	entrySummary = ui.NewMultilineEntry()

	winMain = ui.NewWindow("PR Viewer "+Version, windowWidth, windowHeight, true)
	winMain.SetMargined(true)

	winMain.OnClosing(func(*ui.Window) bool {
		ui.Quit()
		return true
	})

	ui.OnShouldQuit(func() bool {
		winMain.Destroy()
		return true
	})
	tab := ui.NewTab()
	winMain.SetChild(tab)
	winMain.SetMargined(true)

	tab.Append("PR Viewer", buildViewerUI())
	tab.SetMargined(0, true)

	tab.Append("Configuration ", buildConfigUI())
	tab.SetMargined(1, true)

	tab.Append("Log ", buildLoggerUI())
	tab.SetMargined(2, true)

	winMain.Show()
}

func main() {
	//go updateLogger()
	go processLog()
	err := ConfigSetup()
	if err != nil {
		panic(err)
	}
	ui.Main(setupUI)
}
