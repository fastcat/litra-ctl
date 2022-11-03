package cmd

import (
	"github.com/derickr/go-litra-driver"
	"github.com/spf13/cobra"
)

func newOn() *cobra.Command {
	var dev *litra.LitraDevice
	on := &cobra.Command{
		Use:   "on",
		Short: "Turns the light on",
		RunE: func(cmd *cobra.Command, args []string) error {
			dev.TurnOn()
			return nil
		},
	}
	withDevice(on, &dev)
	return on
}
