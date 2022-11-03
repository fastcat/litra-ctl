package cmd

import (
	"fmt"
	"strconv"

	"github.com/derickr/go-litra-driver"
	"github.com/spf13/cobra"
)

func newBrightness() *cobra.Command {
	var dev *litra.LitraDevice
	brightness := &cobra.Command{
		Use:   "brightness",
		Short: "Sets the brightness",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			value := args[0]
			return setBrightness(dev, value)
		},
	}
	withDevice(brightness, &dev)
	return brightness
}

func setBrightness(dev *litra.LitraDevice, value string) error {
	brightness, err := strconv.Atoi(value)
	if err != nil {
		return fmt.Errorf("bad brightness value %q: %w", value, err)
	}
	if brightness < 0 {
		return fmt.Errorf("brightness must be >= 0")
	} else if brightness > 100 {
		return fmt.Errorf("brightness must be <= 100")
	}
	dev.SetBrightness(brightness)
	return nil
}
