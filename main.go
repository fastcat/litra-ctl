package main

import (
	"github.com/fastcat/litra-ctl/cmd"
	"github.com/spf13/cobra"
)

func main() {
	cobra.CheckErr(cmd.Execute())
}
