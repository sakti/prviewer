package main

import (
	"bytes"

	"github.com/BurntSushi/toml"
	"github.com/shibukawa/configdir"
)

var (
	cfg = Config{}
)

const (
	configName = "config.toml"
)

type Config struct {
	GithubToken       string
	MonitoredBranches []string
	RepoOwner         string
	RepoName          string
}

func (c *Config) String() string {
	buf := new(bytes.Buffer)
	if err := toml.NewEncoder(buf).Encode(c); err != nil {
		panic(err)
	}
	return buf.String()
}

var (
	configDirs = configdir.New("github.com:sakti", "prviewer")
)

func GithubToken() string {
	return ""
}

func MonitoredBranches() []string {
	return []string{""}
}

func SetGithubToken(token string) {
}

func SetMonitoredBranches(branches []string) {
}

func ConfigSave(saveToken bool) {
	savedCfg := Config(cfg)
	if !saveToken {
		savedCfg.GithubToken = ""
	}

	folders := configDirs.QueryFolders(configdir.Global)
	cfgFolder := folders[0]
	if len(folders) == 0 {
		Error("config folder not found")
	}
	folder := configDirs.QueryFolderContainsFile(configName)
	if folder == nil {
		Error("config directory with config file not found")
		err := cfgFolder.WriteFile(configName, []byte(savedCfg.String()))
		if err != nil {
			Errorf("cannot save config file %s", err)
		}
		return
	}
	if err := folder.WriteFile(configName, []byte(savedCfg.String())); err != nil {
		Errorf("cannot save config file %s", err)
	}
}

func ConfigSetup() error {
	cfg = Config{
		GithubToken:       "",
		MonitoredBranches: []string{"master"},
	}

	folders := configDirs.QueryFolders(configdir.Global)
	if len(folders) == 0 {
		panic("config folder not found")
	}
	cfgFolder := folders[0]

	folder := configDirs.QueryFolderContainsFile(configName)
	if folder != nil {
		Info("reading config file")
		data, _ := folder.ReadFile(configName)
		if _, err := toml.Decode(string(data), &cfg); err != nil {
			panic(err)
		}
	} else {
		// create default config
		Info("create default config file")
		err := cfgFolder.WriteFile(configName, []byte(cfg.String()))
		if err != nil {
			panic(err)
		}
	}

	return nil
}
