package main

import (
	"strings"

	"github.com/andlabs/ui"
)

var (
	btnConfigSave          *ui.Button
	entryGithubToken       *ui.Entry
	entryRepoOwner         *ui.Entry
	entryRepoName          *ui.Entry
	checkSaveToken         *ui.Checkbox
	entryMonitoredBranches *ui.Entry
)

func onBtnConfigSaveClick(btn *ui.Button) {
	saveToken := false
	if checkSaveToken.Checked() {
		saveToken = true
	}
	cfg.GithubToken = entryGithubToken.Text()

	branches := strings.Split(entryMonitoredBranches.Text(), ",")

	if len(branches) == 0 || entryMonitoredBranches.Text() == "" {
		ui.MsgBoxError(winMain, "Entry branches", "Please insert minimum one branch")
		return
	}
	cfg.MonitoredBranches = branches

	if entryRepoName.Text() == "" || entryRepoOwner.Text() == "" {
		ui.MsgBoxError(winMain, "Repo information", "Please insert repo information")
		return
	}
	cfg.RepoName = entryRepoName.Text()
	cfg.RepoOwner = entryRepoOwner.Text()

	ConfigSave(saveToken)
	ui.MsgBox(winMain, "Saving Config", "saving config success")
}

func loadConfigValue() {
	entryGithubToken.SetText(cfg.GithubToken)
	if cfg.GithubToken != "" {
		checkSaveToken.SetChecked(true)
	}
	entryRepoOwner.SetText(cfg.RepoOwner)
	entryRepoName.SetText(cfg.RepoName)
	entryMonitoredBranches.SetText(strings.Join(cfg.MonitoredBranches, ","))
}

func buildConfigUI() ui.Control {
	entryGithubToken = ui.NewPasswordEntry()
	checkSaveToken = ui.NewCheckbox("save token in config?")
	entryMonitoredBranches = ui.NewEntry()
	entryRepoName = ui.NewEntry()
	entryRepoOwner = ui.NewEntry()

	btnConfigSave := ui.NewButton("Save")
	btnConfigSave.OnClicked(onBtnConfigSaveClick)

	vbox := ui.NewVerticalBox()
	vbox.Append(ui.NewLabel("Configuration"), false)

	frmCfg := ui.NewForm()
	frmCfg.SetPadded(true)

	frmCfg.Append("Github Token", entryGithubToken, false)
	frmCfg.Append("", checkSaveToken, false)
	frmCfg.Append("Repo owner", entryRepoOwner, false)
	frmCfg.Append("Repo name", entryRepoName, false)
	frmCfg.Append("Monitored branchs (comma separated)", entryMonitoredBranches, false)

	vbox.Append(frmCfg, false)

	hbox := ui.NewHorizontalBox()
	hbox.SetPadded(false)
	hbox.Append(btnConfigSave, false)

	vbox.Append(hbox, false)

	loadConfigValue()
	return vbox

}
