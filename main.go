package main

import (
	"fmt"
	"log"
	"os"
	"runtime/debug"

	"github.com/Meiroudii/rockbuster/v3/cli/dir"
	"github.com/Meiroudii/rockbuster/v3/cli/dns"
	"github.com/Meiroudii/rockbuster/v3/cli/fuzz"
	"github.com/Meiroudii/rockbuster/v3/cli/gcs"
	"github.com/Meiroudii/rockbuster/v3/cli/s3"
	"github.com/Meiroudii/rockbuster/v3/cli/tftp"
	"github.com/Meiroudii/rockbuster/v3/cli/vhost"
	"github.com/Meiroudii/rockbuster/v3/libgobuster"
	"github.com/urfave/cli/v2"

	"go.uber.org/automaxprocs/maxprocs"
)

func main() {
	if _, err := maxprocs.Set(); err != nil {
		fmt.Printf("Error on gomaxprocs: %v\n", err) // nolint forbidigo
	}

	cli.VersionPrinter = func(_ *cli.Context) {
		fmt.Printf("rockbuster version %s\n", libgobuster.VERSION) // nolint:forbidigo
		if info, ok := debug.ReadBuildInfo(); ok {
			fmt.Printf("Build info:\n") // nolint forbidigo
			fmt.Printf("%s", info)      // nolint forbidigo
		}
	}

	app := &cli.App{
		Name:      "rockbuster",
		Usage:     "Use this tool ethically",
		UsageText: "rockbuster command [command options]",
		Authors: []*cli.Author{
			{
				Name: "Meiroudii" (@meiroudii)",
			},
		},
		Version: librockbuster.GetVersion(),
		Commands: []*cli.Command{
			dir.Command(),
			vhost.Command(),
			dns.Command(),
			fuzz.Command(),
			tftp.Command(),
			s3.Command(),
			gcs.Command(),
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
