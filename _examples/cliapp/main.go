package main

import (
	"github.com/gookit/color"
	"github.com/gookit/gcli/v3"
	"github.com/gookit/gcli/v3/_examples/cmd"
	// "github.com/gookit/gcli/v3/builtin/filewatcher"
	// "github.com/gookit/gcli/v3/builtin/reverseproxy"
)

// local run:
// 	go run ./_examples/cliapp
// 	go build ./_examples/cliapp && ./cliapp
//
// run on windows(cmd, powerShell):
// 	go build ./_examples/cliapp && ./cliapp
func main() {
	app := gcli.NewApp(func(app *gcli.App) {
		app.Version = "1.0.6"
		app.Desc = "this is my cli application"
		app.On(gcli.EvtAppInit, func(data ...interface{}) {
			// do something...
			// fmt.Println("init app")
		})

		// app.SetVerbose(gcli.VerbDebug)
		// app.DefaultCommand("example")
		app.Logo.Text = `   ________    _______
  / ____/ /   /  _/   |  ____  ____
 / /   / /    / // /| | / __ \/ __ \
/ /___/ /____/ // ___ |/ /_/ / /_/ /
\____/_____/___/_/  |_/ .___/ .___/
                     /_/   /_/`
	})

	var customGOpt string
	app.GOptsBinder = func(gf *gcli.Flags) {
		// gcli.Logf(gcli.VerbInfo, "custom add and global option flag")
		gf.StrVar(&customGOpt, &gcli.FlagMeta{Name: "custom", Desc: "desc message for the option"})
	}

	// app.Strict = true

	app.Add(cmd.GitCmd)

	// app.Add(cmd.ExampleCommand())
	// app.Add(cmd.DaemonRunCommand())
	// app.Add(cmd.EnvInfoCommand())
	// app.Add(cmd.ColorCommand(), cmd.EmojiDemoCmd())
	// app.Add(cmd.ShowDemoCommand(), cmd.ProgressDemoCmd(), cmd.SpinnerDemoCmd(), cmd.InteractDemoCommand())
	// app.Add(builtin.GenEmojiMapCommand())
	// app.Add(builtin.GenAutoComplete())

	// app.Add(filewatcher.FileWatcher(nil))
	// app.Add(reverseproxy.ReverseProxyCommand())

	app.Add(&gcli.Command{
		Name:    "test",
		Aliases: []string{"ts"},
		Desc:    "this is a description <info>message</> for command {$cmd}",
		Func: func(cmd *gcli.Command, args []string) error {
			gcli.Print("hello, in the test command\n")
			return nil
		},
	})

	// create by func
	gcli.NewCommand("test1", "description1", func(c *gcli.Command) {
		// some config for the command
	}).WithFunc(func(c *gcli.Command, args []string) error {
		color.Green.Println("hello, command is: ", c.Name)
		return nil
	}).AttachTo(app)

	// fmt.Printf("%+v\n", gcli.CommandNames())

	// running
	app.Run(nil)
}
