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
	"github.com/allyring/commtk/cfg"
	"os"
	"strconv"

	boxer "github.com/treilik/bubbleboxer"

)

// =====================================================================================================================
// SECTIONS:

// (Tool bubbles get stored in separate modules that get imported.)

// Model containers & relevant functions


// Model container View, Update, and Init functions

// bubbleboxer tree generator from config parsing

// Model edit function

// =====================================================================================================================

// Main model (bubbleboxer model)

type model struct {
	layout boxer.Boxer
}

// =====================================================================================================================

// Function to generate the bubbleboxer layout from the config

// TODO: Add when tool bubbles made

//func generateBoxerLayout(currentNode cfg.LayoutNode) () {
//
//}

// =====================================================================================================================

// Main function

func main() {
	// JSON to layout
	jsonContent, err := os.ReadFile("defaultLayout.json")
	err, schemaErr := cfg.LoadLayoutFile(string(jsonContent))
	if err != nil {
		panic(err.Error())
		os.Exit(1)
	}
	if len(schemaErr) > 0 {
		fmt.Println("JSON Schema error:")
		for errorCount, errorInfo := range schemaErr {
			fmt.Println(strconv.Itoa(errorCount + 1) + " - " + errorInfo.Description())
		}
		os.Exit(1)
	}

	// Load the config into a new bubbleboxer model.






}