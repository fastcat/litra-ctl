package cmd

import (
	"fmt"
	"strconv"

	"github.com/derickr/go-litra-driver"
	"github.com/spf13/cobra"
)

func newTemperature() *cobra.Command {
	var dev *litra.LitraDevice
	temperature := &cobra.Command{
		Use:   "temperature",
		Short: "Sets the temperature",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			return setTemperature(dev, args[0])
		},
	}
	withDevice(temperature, &dev)
	return temperature
}

func setTemperature(dev *litra.LitraDevice, value string) error {
	temperature, err := strconv.Atoi(value)
	if err != nil {
		return fmt.Errorf("bad temperature value %q: %w", value, err)
	}
	if temperature < 2700 {
		return fmt.Errorf("temperature must be >= 2700")
	} else if temperature > 6500 {
		return fmt.Errorf("temperature must be <= 6500")
	}
	dev.SetTemperature(int16(temperature))
	return nil
}
