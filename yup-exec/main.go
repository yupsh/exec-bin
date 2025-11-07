package main

import (
	"fmt"
	"os"

	"github.com/urfave/cli/v2"

	yup "github.com/gloo-foo/framework"
	. "github.com/yupsh/exec"
)

const (
	flagWorkingDir   = "directory"
	flagEnvVar       = "env"
	flagShell        = "shell"
	flagUseShell     = "use-shell"
	flagIgnoreErrors = "ignore-errors"
	flagQuiet        = "quiet"
	flagInteractive  = "interactive"
	flagInheritEnv   = "inherit-env"
)

func main() {
	app := &cli.App{
		Name:  "exec",
		Usage: "execute external commands",
		UsageText: `exec [OPTIONS] COMMAND [ARG...]

   Execute COMMAND with given arguments.`,
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:    flagWorkingDir,
				Aliases: []string{"C"},
				Usage:   "run command in DIRECTORY",
			},
			&cli.StringSliceFlag{
				Name:    flagEnvVar,
				Aliases: []string{"e"},
				Usage:   "set environment variable (NAME=VALUE)",
			},
			&cli.StringFlag{
				Name:  flagShell,
				Usage: "shell to use for execution",
			},
			&cli.BoolFlag{
				Name:    flagUseShell,
				Aliases: []string{"s"},
				Usage:   "execute command through shell",
			},
			&cli.BoolFlag{
				Name:  flagIgnoreErrors,
				Usage: "ignore command execution errors",
			},
			&cli.BoolFlag{
				Name:    flagQuiet,
				Aliases: []string{"q"},
				Usage:   "suppress command output",
			},
			&cli.BoolFlag{
				Name:    flagInteractive,
				Aliases: []string{"i"},
				Usage:   "run in interactive mode",
			},
			&cli.BoolFlag{
				Name:  flagInheritEnv,
				Usage: "inherit environment variables from parent",
			},
		},
		Action: action,
	}

	if err := app.Run(os.Args); err != nil {
		fmt.Fprintf(os.Stderr, "exec: %v\n", err)
		os.Exit(1)
	}
}

func action(c *cli.Context) error {
	var params []any

	// Add all arguments as command
	for i := 0; i < c.NArg(); i++ {
		params = append(params, c.Args().Get(i))
	}

	// Add flags based on CLI options
	if c.IsSet(flagWorkingDir) {
		params = append(params, WorkingDir(c.String(flagWorkingDir)))
	}
	if c.IsSet(flagEnvVar) {
		for _, env := range c.StringSlice(flagEnvVar) {
			params = append(params, EnvVar(env))
		}
	}
	if c.IsSet(flagShell) {
		params = append(params, Shell(c.String(flagShell)))
	}
	if c.Bool(flagUseShell) {
		params = append(params, UseShell)
	}
	if c.Bool(flagIgnoreErrors) {
		params = append(params, IgnoreErrors)
	}
	if c.Bool(flagQuiet) {
		params = append(params, Quiet)
	}
	if c.Bool(flagInteractive) {
		params = append(params, Interactive)
	}
	if c.Bool(flagInheritEnv) {
		params = append(params, InheritEnv)
	}

	// Create and execute the exec command
	cmd := Exec(params...)
	return yup.Run(cmd)
}
