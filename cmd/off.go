package cmd

import (
	"github.com/derickr/go-litra-driver"
	"github.com/spf13/cobra"
)

func newOff() *cobra.Command {
	var dev *litra.LitraDevice
	off := &cobra.Command{
		Use:   "off",
		Short: "Turns the light off",
		RunE: func(cmd *cobra.Command, args []string) error {
			dev.TurnOff()
			return nil
		},
	}
	withDevice(off, &dev)
	return off
}
