package command

import (
	"strings"
	//"fmt"
	//"gopkg.in/src-d/go-git.v4"
	//"gopkg.in/src-d/go-git.v4/storage/memory"
	//"gopkg.in/src-d/go-git.v4/plumbing"
)

type ListCommand struct {
	Meta
}

func (c *ListCommand) Run(args []string) int {
	//c.appName = args[1]
	// git branch --merged
	//r, err := git.Clone(memory.NewStorage(), nil, &git.CloneOptions{
	//	URL: "https://github.com/zafarella/envs",
	//})
	//CheckIfError(err)
	//
	//err = r.Fetch(&git.FetchOptions{
	//})
	//
	//list, err := r.Branches()
	//CheckIfError(err)
	//
	//list.ForEach(func(reference *plumbing.Reference) error {
	//	fmt.Println(reference.Name())
	//	return nil
	//})
	return 0
}

func (c *ListCommand) Synopsis() string {
	return ""
}

func (c *ListCommand) Help() string {
	helpText := `
					List branches for the given environment.
`
	return strings.TrimSpace(helpText)
}
