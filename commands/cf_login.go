package commands

import (
	"fmt"
)

type CFLoginCommand struct {
	Lockfile string `short:"l" long:"lockfile" env:"ENVIRONMENT_LOCK_METADATA" description:"path to a lockfile"`
	File     bool   `short:"f" long:"file" description:"write a script file but do not run it"`

	Env           EnvReader
	CFLoginRunner ToolRunner
}

func (c *CFLoginCommand) Execute(args []string) error {
	data, err := c.Env.Read(c.Lockfile)
	if err != nil {
		return err
	}

	fmt.Printf("Logging in to: %s\n", data.OpsManager.URL.String())

	return c.CFLoginRunner.Run(data, c.File)
}
