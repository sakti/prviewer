package main

import (
	"github.com/andlabs/ui"
)

var (
	btnGetPR    *ui.Button
	entryPR     *ui.MultilineEntry
	comboBranch *ui.Combobox

	progressProcess *ui.ProgressBar
)

func reloadComboBranch() {
	for _, v := range cfg.MonitoredBranches {
		comboBranch.Append(v)
	}
	comboBranch.SetSelected(0)
}

func onBtnGetPrClick(btn *ui.Button) {
	value := progressProcess.Value() + 10
	if value > 100 {
		value = 0
	}

	branchIdx := comboBranch.Selected()
	branchLabel := cfg.MonitoredBranches[branchIdx]

	progressProcess.SetValue(value)

	Infof("starting get PR info from branch %s, from repo %s/%s", branchLabel, cfg.RepoOwner, cfg.RepoName)

	btnGetPR.Disable()
	comboBranch.Disable()
	go getPR(branchLabel)
}

func buildViewerUI() ui.Control {
	btnGetPR = ui.NewButton("Get PR Summary")
	btnGetPR.OnClicked(onBtnGetPrClick)
	entryPR = ui.NewMultilineEntry()
	progressProcess = ui.NewProgressBar()
	comboBranch = ui.NewCombobox()

	vbox := ui.NewVerticalBox()
	vbox.Append(progressProcess, false)
	vbox2 := ui.NewVerticalBox()
	vbox2.SetPadded(true)
	vbox2.Append(btnGetPR, false)
	vbox2.Append(comboBranch, false)

	hbox := ui.NewHorizontalBox()
	hbox.SetPadded(true)
	hbox.Append(vbox2, false)
	hbox.Append(entryPR, true)
	vbox.Append(hbox, true)

	reloadComboBranch()
	entryPR.SetReadOnly(true)

	return vbox
}
