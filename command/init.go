package command

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"strings"
)

type InitCommand struct {
	Meta
}

func (c *InitCommand) Run(args []string) int {
	// Write your code here
	fmt.Println("args = ", args)

	osGitUser, err := exec.Command("git", "config", "--get", "user.name").Output()
	if err != nil {
		log.Fatal(err)
		fmt.Println("unable to find git installed. Not in $PATH ? Not installed?")
		os.Exit(126)
	}

	if len(osGitUser) != 0 {
		fmt.Println("Do you want me to use your git username as ?", osGitUser)
	} else {
		fmt.Println("What should I use as your git username ?")
	}

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

type Configs struct {
	Git struct {
		gitUserName   string
		gitUserEmail  string
		gitUserEditor string
	}
	currVer string
}
