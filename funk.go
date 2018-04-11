package main

import (
	"flag"
	"fmt"
	"os"
	"plugin"
	"strings"

	"github.com/sh3rp/funk/funk"
)

var pluginName string

func main() {
	flag.StringVar(&pluginName, "plugin", "funk", "name of the plugin without the .so extension")
	flag.Parse()

	if strings.Contains(pluginName, string(os.PathSeparator)) {
		elements := strings.Split(pluginName, string(os.PathSeparator))
		pluginName = elements[len(elements)-1]
	}

	fmt.Printf("Loading plugin: %s\n", pluginName)

	plugin, err := plugin.Open(fmt.Sprintf("%s.so", pluginName))

	if err != nil {
		fmt.Printf("Error loading plugin: %v\n", err)
		return
	}

	function, err := plugin.Lookup("FunkHandler")

	if err != nil {
		fmt.Printf("Error looking up function: %v\n", err)
		return
	}

	processor := funk.NewGeneratorEventProcessor(function.(func(funk.Event)))
	fmt.Printf("ERROR: %v\n", processor.Process())
}
