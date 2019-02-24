package main

import "github.com/andlabs/ui"

func buildLoggerUI() ui.Control {
	entryLog = ui.NewMultilineEntry()
	vbox := ui.NewVerticalBox()
	vbox.Append(ui.NewLabel("Log:"), false)
	vbox.Append(entryLog, true)

	entryLog.SetReadOnly(true)

	return vbox
}
