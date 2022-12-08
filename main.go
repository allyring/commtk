/*
Commtk is a suite of TUIs for common TCP/IP networking tools.
It displays tools on a scalable grid, where the contents and sizing of the grid's cells are defined by a config file.
It contains TUIs for the following tools:
- ping
- traceroute
- ifconfig

Usage:

	commtk [config file path ...]

Licensed under GNU GPLv3
*/
package main

import (
	"fmt"
	"os"

	// Packages for the project
	"github.com/allyring/commtk/cfg"

	// Bubbletea, stickers, and bubbles - for TUI rendering
	"github.com/76creates/stickers"
)
// ---------------------------------------------------------------------------------------------------------------------
// Bubbletea model and model initialisation function
type model struct {
	err		error				// Contains the most recent error encountered
	flexBox	*stickers.FlexBox	// Pointer to a flexbox in memory
}

//func ModelInit(config string) model {
//	m := model{
//		err: nil,
//		flexBox: stickers.NewFlexBox(0,0),
//	}
//
//}

// ---------------------------------------------------------------------------------------------------------------------

func main() {
	// The args are a path to a custom config. If the args exist, pass them into the config layout generator.
	configPath := ""
	if len(os.Args[1:]) > 1 {
		fmt.Println("Cannot accept several configs.")
		os.Exit(1)
	}
	if len(os.Args[1:]) == 1 {
		configPath = os.Args[1]
	}


	// Create the layout config from the given file / the default config.
	err, schemaErr := cfg.LoadLayoutFile(configPath)
	if err != nil {
		panic(err)
	}

	if schemaErr != nil {
		fmt.Println("JSON errors")
		for _, schemaErrItem := range schemaErr {
			fmt.Println("- " + schemaErrItem.Description())
		}
		panic(err)
	}

	fmt.Println(cfg.Config.Layout[2][0].Tool)

}