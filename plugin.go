package main

import (
	"context"
	"fmt"

	goplugin "github.com/hashicorp/go-plugin"

	"aspect.build/cli/pkg/bazel"
	"aspect.build/cli/pkg/plugin/sdk/v1alpha3/config"
	aspectplugin "aspect.build/cli/pkg/plugin/sdk/v1alpha3/plugin"
)

func main() {
	goplugin.Serve(config.NewConfigFor(NewDefaultPlugin()))
}

type HelloWorldPlugin struct {
	aspectplugin.Base
}

func NewDefaultPlugin() *HelloWorldPlugin {
	return NewPlugin()
}

func NewPlugin() *HelloWorldPlugin {
	return &HelloWorldPlugin{}
}

func (plugin *HelloWorldPlugin) CustomCommands() ([]*aspectplugin.Command, error) {
	return []*aspectplugin.Command{
		aspectplugin.NewCommand(
			"hello-world",
			"Print 'Hello World!' to the command line.",
			"Print 'Hello World!' to the command line. Echo any given argument. Then run a 'bazel help'",
			func(ctx context.Context, args []string, bzl bazel.Bazel) error {
				fmt.Println("Hello World!")
				fmt.Print("Arguments passed to command: ")
				fmt.Println(args)
				fmt.Println("Going to run: 'bazel help'")

				bzl.Spawn([]string{"help"})

				return nil
			},
		),
	}, nil
}
