package main

import (
	"github.com/mitchellh/cli"
	"github.com/zafarella/envs/command"
)

func Commands(meta *command.Meta) map[string]cli.CommandFactory {
	return map[string]cli.CommandFactory{
		"deploy": func() (cli.Command, error) {
			return &command.DeployCommand{
				Meta: *meta,
			}, nil
		},
		"remove": func() (cli.Command, error) {
			return &command.RemoveCommand{
				Meta: *meta,
			}, nil
		},
		"reset": func() (cli.Command, error) {
			return &command.ResetCommand{
				Meta: *meta,
			}, nil
		},
		"list": func() (cli.Command, error) {
			return &command.ListCommand{
				Meta: *meta,
			}, nil
		},
		"update": func() (cli.Command, error) {
			return &command.UpdateCommand{
				Meta: *meta,
			}, nil
		},
		"init": func() (cli.Command, error) {
			return &command.InitCommand{
				Meta: *meta,
			}, nil
		},

		"version": func() (cli.Command, error) {
			return &command.VersionCommand{
				Meta:     *meta,
				Version:  Version,
				Revision: GitCommit,
				Name:     Name,
			}, nil
		},
	}
}
