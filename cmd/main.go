package main

import (
	"fmt"
	"github.com/TurnsCoffeeIntoScripts/nexus/pkg/engine"
	"github.com/TurnsCoffeeIntoScripts/nexus/pkg/engine/cli"
	"github.com/TurnsCoffeeIntoScripts/nexus/pkg/system"
)

func main() {
	e := engine.CreateEngine()
	e.Start()

	w, h := system.GetTerminalDimension()

	s := cli.Shell{Width: w, Height: h}

	fmt.Printf("w --> %v\nh --> %v\n", s.Width, s.Height)
}
