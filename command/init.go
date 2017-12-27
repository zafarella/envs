package command

import (
	"log"
	"os"
	"os/exec"
	"strings"
	"fmt"
	"github.com/spf13/viper"
)

type InitCommand struct {
	Meta
}

/**
 Used to check the configs and create one if not exists.
 */

func (c *InitCommand) Run(args []string) int {
	// git user
	osGitUser, err := exec.Command("git", "config", "--get", "user.name").Output()
	if err != nil {
		log.Fatal(err)
		Error("unable to find git installed. Not in $PATH ? Not installed?")
		os.Exit(126)
	}

	if len(osGitUser) == 0 {
		Info("Was unable to find gituser. Enter your git username please:")
		fmt.Scanln(&osGitUser)
	} else {
		Info("Using %s as your git username.", osGitUser)
	}
	// get email
	email, err := exec.Command("git", "config", "--get", "user.email").Output()
	if len(email) != 0 {
		Info("Using %s as your email (for notifications)", email)
	}
	// default editor for conflicts resolution
	defaultEditor := os.Getenv("EDITOR")
	if defaultEditor == "" {
		defaultEditor = "vim"
	}
	Info("Using %s as your default editor for conflicts resolutions.", defaultEditor)

	const configFile = "git_configs"
	//confs = config.gitConfigs
	viper.SetConfigName(configFile)
	viper.AddConfigPath(".")
	viper.Set("werwer","asdads")
	viper.WriteConfigAs("git_configs")
	Info("Wrote configs to %s", configFile)
	return 0
}

func (c *InitCommand) Synopsis() string {
	return ""
}

func (c *InitCommand) Help() string {
	helpText := `
	Initializes the tool. You have to provide all necessary information. Could be used to re-configure setting.
`
	return strings.TrimSpace(helpText)
}
