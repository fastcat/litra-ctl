package cmd

import (
	"fmt"

	"github.com/derickr/go-litra-driver"
	"github.com/spf13/cobra"
)

func newScript() *cobra.Command {
	var dev *litra.LitraDevice
	script := &cobra.Command{
		Use:   "script",
		Short: "Run a script of commands in sequence",
		Long:  "Run a script of commands in sequence, each command being one of the single-step items such as on, off",
		Args:  cobra.MinimumNArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			for i := 0; i < len(args); i++ {
				c := args[i]
				switch c {
				case "on":
					fmt.Println("Turning on")
					dev.TurnOn()
				case "off":
					fmt.Println("Turning off")
					dev.TurnOff()
				case "brightness":
					if len(args) <= i+1 {
						return fmt.Errorf("brightness requires an arg")
					}
					fmt.Println("Setting brightness to", args[i+1])
					if err := setBrightness(dev, args[i+1]); err != nil {
						return err
					}
					i++
				case "temperature":
					if len(args) <= i+1 {
						return fmt.Errorf("temperature requires an arg")
					}
					fmt.Println("Setting temperature to", args[i+1])
					if err := setTemperature(dev, args[i+1]); err != nil {
						return err
					}
					i++
				default:
					return fmt.Errorf("invalid command %q", c)
				}
			}
			return nil
		},
	}
	withDevice(script, &dev)
	return script
}
